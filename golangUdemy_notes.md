# 📡 Exemplo de Channels em Go

## 🎯 O que este código demonstra

Comunicação básica entre **goroutines** usando **channels** (canais).

## 🧠 Conceito

> **Channel é um "tubo" que conecta goroutines. O que entra por um lado, sai do outro.**

## 📝 Código

```go
func main() {
    ch := make(chan string)        // 1. Cria o canal

    go func() {                     // 2. Goroutine que ESCREVE
        ch <- "Olá!"                //    Envia mensagem
    }()

    msg := <-ch                     // 3. LÊ a mensagem
    fmt.Println(msg)                //    Saída: "Olá!"
}

```

# 🔍 Como funciona  
```go
Linha                   	Ação
make(chan string)	        Cria um canal que transporta string
ch <- "Olá!"	            ESCREVE no canal (seta ←)
msg := <-ch	                LÊ do canal (seta →)

```

# ⚡ Sincronização Automática  


**Sem time.Sleep, sem WaitGroup!**

- A escrita bloqueia até alguém ler  
- A leitura bloqueia até alguém escrever  

```go

ch <- "mensagem"  // ⏸️ Trava até alguém ler
msg := <-ch       // ⏸️ Trava até alguém escrever

```  
# 🖼️ Visualização  

```go

Goroutine 1                    Goroutine 2 (main)
     │                              │
     ├── ch <- "Olá!"               │
     │      │                        │
     │      └──────[ 📦 ]──────────→ │
     │                              ├── msg := <-ch
     │                              │
     └── (desbloqueia)              └── (desbloqueia)

```
```go
✅ Resumo
Operação	            Símbolo	O que faz
Criar	                make(chan tipo)	Cria um novo canal
Escrever	            ch <- valor	Envia valor para o canal
Ler	                    var := <-ch	Recebe valor do canal
``` 

