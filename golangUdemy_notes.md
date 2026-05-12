# 📡 Exemplo de Channels em Go

## 🎯 O que este código demonstra

Comunicação básica entre **goroutines** usando **channels** (canais).

## 🧠 Conceito

> **Channel é um "tubo" que conecta goroutines. O que entra por um lado, sai do outro.**

## 📝 Código

# 🧠 Entendendo Race Condition em Go

## 🎯 O que este código ensina

**O problema de múltiplas goroutines acessando uma variável compartilhada SEM sincronização.**

```go
var counter int = 0  // recurso compartilhado

func IncrementCount(wg *sync.WaitGroup) {
    defer wg.Done()
    counter++  // ← APENAS 1 LINHA, mas NÃO é segura!
}

func main() {
    var wg sync.WaitGroup
    
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go IncrementCount(&wg)
    }
    
    wg.Wait()
    fmt.Println(counter)  // 🤔 Resultado? 1000? Nem sempre!
}
```

# 🔬 O problema: counter++ não é uma operação atômica 
**Parece uma linha, mas o computador faz 3 operações separadas:**

counter++  // O compilador transforma em:
```go

// 1. LEIA o valor atual de counter
// 2. SOME 1 ao valor lido
// 3. ESCREVA o novo valor de volta

```

# 💥 O que acontece com múltiplas goroutines 

**Cenário com 2 goroutines tentando incrementar ao mesmo tempo:** 
```go
Estado inicial: counter = 0

TEMPO   Goroutine A          Goroutine B          counter
─────────────────────────────────────────────────────────
t1      LEU: 0                                       0
t2                           LEU: 0                  0
t3      SOMOU: 0+1 = 1                              0
t4                           SOMOU: 0+1 = 1          0
t5      ESCREVEU: 1                                  1
t6                           ESCREVEU: 1             1  ← PERDEMOS UM INCREMENTO!

```
**Resultado: 2 incrementos, mas counter = 1! ❌** 


# 📊 Resultados típicos ao rodar 1000 goroutines  

´´'go
# Execute várias vezes:
```go 
go run main.go

# Saída 1: 985  (perdeu 15 incrementos)
# Saída 2: 1000 (sorte! nenhuma colisão)
# Saída 3: 923  (perdeu 77 incrementos)
# Saída 4: 997  (perdeu 3 incrementos)

```   
**Cada execução dá um resultado diferente! Isso é uma Race Condition.**  
---
# 🧠 Analogia: O banheiro público  
```go
Sem Mutex (problema):
═══════════════════════════════════════
Banheiro sem chave. 10 pessoas tentam entrar:
- Pessoa A entra
- Pessoa B entra também (não esperou!)
- Pessoa C entra também!
- 🔥 CAOS! Ninguém consegue usar direito.


Com Mutex (solução):
═══════════════════════════════════════
Banheiro com chave:
- Pessoa A: pega a chave → entra → tranca → sai → devolve
- Pessoa B: espera a chave → pega → entra → tranca → sai
- Pessoa C: espera...
- ✅ Uma pessoa por vez! Organizado!

```

# 📋 Próxima branch: A solução com Mutex
