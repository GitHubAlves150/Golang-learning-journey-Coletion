# Golang
## Introdução
Este README, é uma breve anotação que fiz ao longo do mês de Abril de 2026 sobe Golang e suas particularidade. Aqui eu anotei os principais conceitos da linguagem para eu guiar um caminho solido e progressivo em Golang.
Sempre que preciso recuperar ou rever conceitos eu reviso por aqui. mas fácil saber onde está do que ficar procurando aleatóriamente pela internet.

**Sintaxe básica**

### Menu
- [Sintaxe Básica - Variáveis em Go](#sintaxe-basica)
- [Operadores em Golnag](#operadores)
- [Condicionais](#condicionais)
<br><br>:

---
# Sintaxe Basica

## 🎯 Objetivo
Entender como declarar, inicializar e usar variáveis em Go, incluindo particularidades como **zero values**, inferência de tipo e escopo.

---

**2. Declaração com Inicialização**

```go
var nome string = "João"
var idade int = 30
var ativo bool = true

// Múltiplas
var x, y int = 10, 20
var (
    nome   string = "Maria"
    cidade string = "São Paulo"
)
```

**3. Inferência de Tipo (:= - operador curto)**

```go
nome := "Carlos"        // string
idade := 25             // int
preco := 99.99          // float64
ativo := true           // bool

// Múltiplas
nome, idade := "Ana", 28
x, y := 10, 20

```
⚠️ Regra: := só pode ser usado dentro de funções

**4. Zero Values (Valores Padrão)**
- Em Go, toda variável declarada tem um valor padrão (não é null ou undefined):
  
```go
Tipo                    	                    Zero Value

int, int8, int64	                            0
float32, float64	                            0.0
bool	                                        false
string	                                        "" (vazia)
pointer, slice, map, chan, func, interface	    nil
----------------Exemplo----------------------------------------
var a int       // 0
var b string    // ""
var c bool      // false
var d *int      // nil

fmt.Printf("a=%d, b='%s', c=%t, d=%v\n", a, b, c, d)
// Output: a=0, b='', c=false, d=<nil>

```
**5. Tipos Básicos em Go**
```go
// Inteiros (escolha o menor necessário)
var i int     // depende da arquitetura (32 ou 64 bits)
var i8 int8   // -128 a 127
var i16 int16 // -32768 a 32767
var i32 int32 // -2.147.483.648 a 2.147.483.647
var i64 int64 // -9.223.372.036.854.775.808 a 9.223.372.036.854.775.807

// Inteiros sem sinal
var u uint    // 0 a 4294967295 (32 bits) ou mais
var u8 uint8  // 0 a 255 (byte)
var u16 uint16
var u32 uint32
var u64 uint64

// Ponto flutuante
var f32 float32 // ~6 casas decimais de precisão
var f64 float64 // ~15 casas decimais (padrão, mais preciso)

// Texto
var s string   // sequência imutável de bytes (UTF-8)
var b byte     // alias para uint8
var r rune     // alias para int32 (representa um caractere Unicode)

// Booleanos
var verdadeiro bool = true
var falso bool = false


```
**6. Conversão de Tipos (Type Conversion)**
- Em Go não há conversão implícita - você precisa fazer explícita:
  
```go
    var x int = 10
var y float64 = float64(x) // ✅ conversão explícita

var a int = 10
var b int64 = int64(a) // ✅ ok

// ❌ Isso NÃO funciona:
// var c float64 = a  // erro: cannot use a (type int) as type float64

```

**7. Constantes**

```go
// Constantes tipadas
const Pi float64 = 3.14159
const Nome string = "GoLang"

// Constantes não tipadas (mais flexíveis)
const (
    SegundosPorMinuto = 60
    MinutosPorHora    = 60
    HorasPorDia       = 24
)

// Iota - enumerador automático
type DiaSemana int
const (
    Domingo DiaSemana = iota // 0
    Segunda                   // 1
    Terca                     // 2
    Quarta                    // 3
    Quinta                    // 4
    Sexta                     // 5
    Sabado                    // 6
)
```
**8. Escopo de Variáveis**
```go
package main

var global = "acessível em todo o pacote" // Pacote scope

func main() {
    local := "acessível só dentro da função" // Função scope
    
    if true {
        bloco := "acessível só dentro deste bloco" // Bloco scope
        fmt.Println(bloco)  // ✅ ok
    }
    
    // fmt.Println(bloco)  // ❌ erro: bloco não definido
    
    // Shadowing (sobrescrita de variável)
    global := "variável local com mesmo nome"  // ⚠️ cria uma nova, não altera a global
}
```
**9. Ponteiros (introdução)**
```go
x := 10
ptr := &x       // ptr armazena o endereço de memória de x

fmt.Println(x)   // 10
fmt.Println(ptr) // 0xc000012088 (endereço)
fmt.Println(*ptr) // 10 (dereferencing - valor apontado)

*ptr = 20       // muda o valor de x via ponteiro
fmt.Println(x)   // 20

```
## Fim do conteúdo sobre sintaxe e declarações ##

[Voltar](#menu)

# Operadores
## Operadores em GO

```go
Operador	            Particularidade em Go	                Importância
= vs :=	:=              declara + atribui (único do Go)	        🔴 CRÍTICO
++ / --	                São instruções, não expressões	        🔴 CRÍTICO
== / !=	                Compara structs campo a campo	        🟡 IMPORTANTE
& / *	                Ponteiros (sem aritmética)	            🟡 IMPORTANTE
&& / ||	                Short-circuit (igual outras langs)	    🟢 BÁSICO
<< / >>	                Shift com tipos unsigned	            🟢 BÁSICO
&^ (AND NOT)	        Bit clear - exclusivo do Go	            🔵 DIFERENTE
```

**1. Operadores Aritméticos (iguais a C/Java)**
```go
+   Adição
-   Subtração  
*   Multiplicação
/   Divisão
%   Módulo (resto)

// Exemplos
a := 10 + 5    // 15
b := 10 - 5    // 5
c := 10 * 5    // 50
d := 10 / 3    // 3 (divisão de inteiros)
e := 10.0 / 3  // 3.333... (float)
f := 10 % 3    // 1
```
---

⚠️ Atenção: Divisão entre inteiros trunca (não arredonda):  
```go
fmt.Println(10 / 3)   // 3 (não 3.33)
fmt.Println(10 / 3.0) // 3.333...
```
----

**2. Operadores de Atribuição (o MAIS importante)**

**❌ O que NÃO funciona em Go:**  

```go
// Isso não compila:
x =+ 1    // Go interpreta como x = (+1)
x =- 1    // Go interpreta como x = (-1)

// Operador ternário não existe:
// resultado = (idade > 18) ? "maior" : "menor"  // ❌

```

✅ O que funciona (igual C/Java):  

```go
=    Atribuição simples
+=   Adição e atribuição
-=   Subtração e atribuição
*=   Multiplicação e atribuição
/=   Divisão e atribuição
%=   Módulo e atribuição
&=   AND bit a bit e atribuição
|=   OR bit a bit e atribuição
^=   XOR bit a bit e atribuição
<<=  Left shift e atribuição
>>=  Right shift e atribuição

// Exemplos
x := 10
x += 5    // x = 15
x *= 2    // x = 30
```

🔴 CRÍTICO: Operador curto := (único do Go)  

```go
// := declara E atribui (só dentro de funções)
nome := "João"        // declara variável nome do tipo string
idade := 30           // declara idade como int

// Equivalente a:
var nome string = "João"
var idade int = 30

// Múltiplas variáveis
x, y := 10, 20
nome, idade := "Maria", 25

// Reatribuição com := (se pelo menos uma variável for nova)
x := 10         // declara x
x, y := 20, 30  // x é reatribuído, y é declarado (✅ ok)
// x := 40      // ❌ erro: x já declarada
```

## 3. 🔴 CRÍTICO: ++ e -- são INSTRUÇÕES, não expressões  
**Diferença fundamental de C/Java/Python/JavaScript:**  

```go 
// ✅ Isso funciona (instrução)
x := 5
x++        // x = 6 (instrução válida)
x--        // x = 5

// ❌ Isso NÃO funciona em Go
// y := x++      // Erro: x++ usado como expressão
// fmt.Println(x++) // Erro
// return x++    // Erro

// ✅ Correto:
x := 5
x++
y := x      // Agora sim, y = 6
```
**4. Operadores Relacionais (Comparação)**
```go
==   Igual
!=   Diferente
<    Menor
<=   Menor ou igual
>    Maior
>=   Maior ou igual
```

**🟡 IMPORTANTE: Comparação de structs**
```go
type Pessoa struct {
    Nome string
    Idade int
}

p1 := Pessoa{"João", 30}
p2 := Pessoa{"João", 30}
p3 := Pessoa{"Maria", 25}

fmt.Println(p1 == p2)  // true (compara campo a campo)
fmt.Println(p1 == p3)  // false

// ⚠️ Slices, maps e functions NÃO são comparáveis com == (exceto com nil)
// slice1 := []int{1,2,3}
// slice2 := []int{1,2,3}
// fmt.Println(slice1 == slice2)  // ❌ Erro de compilação!
```

**Comparação com nil:**  

```go
var s []int
var m map[string]int
var p *int

fmt.Println(s == nil)  // true (slice nil)
fmt.Println(m == nil)  // true (map nil)
fmt.Println(p == nil)  // true (ponteiro nil)
```
**5. Operadores Lógicos (Short-circuit)**
```go
&&   AND lógico (short-circuit)
||   OR lógico (short-circuit)
!    NOT lógico

// Short-circuit: segunda condição só avalia se necessário
func valida(x int) bool {
    fmt.Println("Validando...")
    return x > 0
}

if false && valida(10) {  // valida() NÃO é chamada
    fmt.Println("nunca executa")
}

if true || valida(10) {   // valida() NÃO é chamada
    fmt.Println("sempre executa")
}

```

**6. Operadores Bit a Bit (para performance)**
```go
&    AND bit a bit
|    OR bit a bit
^    XOR bit a bit
&^   AND NOT (bit clear) - Exclusivo do Go!
<<   Left shift
>>   Right shift

// Exemplos práticos
x := 0b1100 (12)
y := 0b1010 (10)

fmt.Printf("%04b\n", x & y)   // 1000 (8)  - AND
fmt.Printf("%04b\n", x | y)   // 1110 (14) - OR
fmt.Printf("%04b\n", x ^ y)   // 0110 (6)  - XOR

// 🔵 &^ (AND NOT): zera bits onde y tem 1
fmt.Printf("%04b\n", x &^ y)  // 0100 (4)  - limpa bits 1 e 3

// Shift
fmt.Println(1 << 10)  // 1024 (deslocamento para esquerda = multiplica por 2^n)
fmt.Println(1024 >> 3) // 128 (deslocamento para direita = divide por 2^n)
```

**Uso prático de bit clear (&^):**  

```go
// Flags de permissão
const (
    Read  = 1 << iota  // 1 (0b001)
    Write              // 2 (0b010)
    Execute            // 4 (0b100)
)

permissoes := Read | Write  // 0b011 (3)

// Remover permissão de escrita usando &^
permissoes = permissoes &^ Write  // 0b001 (1) - apenas Read
```
**7. Operadores de Ponteiros**

```go
&    Endereço de memória (referência)
*    Ponteiro (dereferência)

// Exemplo
x := 42
ptr := &x        // ptr recebe endereço de x
fmt.Println(ptr) // 0xc000012088
fmt.Println(*ptr) // 42 (valor apontado)

*ptr = 100       // muda o valor de x através do ponteiro
fmt.Println(x)   // 100
```

⚠️ Não há aritmética de ponteiros em Go seguro (diferente de C):

```go
// ptr++  // ❌ Erro: aritmética de ponteiros não permitida
```
**8. Precedência de Operadores**

```go
Precedência	                    Operadores
   1 (mais alta)	            * / % << >> & &^
   2	                         + - | ^
   3	                         == != < <= > >=
   4	                           &&
   5                           (mais baixa)	||
```

**Use parênteses para clareza:** 
```go
// Ambíguo
x := 1 + 2*3  // 7 (não 9)

// Claro
y := (1 + 2) * 3  // 9
```

**9. Operador Especial: Vírgula (em declarações)**

```go
// Declaração múltipla
x, y := 10, 20

// Swap (simples!)
a, b := 5, 10
a, b = b, a  // a=10, b=5

// Ignorar valores com blank identifier
nome, _ := obterNome()  // ignora segundo retorno

```

**🐛 Erros Comuns com Operadores**
```go
// 1. Confundir = com :=
x = 10        // ❌ erro se x não foi declarado antes
x := 10       // ✅ correto (declara e atribui)

// 2. Usar ++ como expressão
// y = x++     // ❌ erro

// 3. Comparar tipos diferentes
// if 10 == 10.0  // ❌ erro: tipos incompatíveis (int vs float64)

// 4. Divisão inteiro com float
var a int = 10
var b float64 = 3.0
// c := a / b    // ❌ erro: tipos diferentes
c := float64(a) / b  // ✅ conversão explícita

// 5. Operador ternário
// resultado := idade > 18 ? "maior" : "menor"  // ❌ erro

// Correção:
var resultado string
if idade > 18 {
    resultado = "maior"
} else {
    resultado = "menor"
}

```

[Voltar](#menu)

# Condicionais

**📘 Condicionais em Go - Sintaxe Essencial**

```go
🎯 Resumo Rápido (para consulta diária)  
Recurso	                    Sintaxe	                    Particularidade
if                          if  condicao { }	        Obrigatório {} mesmo com 1 linha
if-else	                    if cond { } else { }	    else na MESMA linha da }
if com init	                if x := 10; x > 0 { }	    Variável existe SOMENTE no if/else
switch	                    switch valor { case 1: }	Fallthrough é opcional (não automático)
switch true	                switch { case x > 0: }	    Switch sem expressão = switch true
type switch	                switch v.(type) { }	        Para interfaces


```

**1. If Statement (Obrigatório {})**  
```go
// ✅ Correto
if idade >= 18 {
    fmt.Println("Maior de idade")
}

// ❌ Erro - chaves obrigatórias MESMO com 1 linha
// if idade >= 18 fmt.Println("Maior")  // NÃO COMPILA

// ✅ Else na MESMA linha da chave de fechamento
if idade >= 18 {
    fmt.Println("Maior")
} else {
    fmt.Println("Menor")
}

// ✅ Else if
if nota >= 7 {
    fmt.Println("Aprovado")
} else if nota >= 5 {
    fmt.Println("Recuperação")
} else {
    fmt.Println("Reprovado")
}

```
**2. 🟡 If com Inicialização (Init Statement)**

**Característica única do Go que você vai usar MUITO:**

```go
// Variável declarada SÓ existe dentro do if/else
if err := processar(); err != nil {
    fmt.Println("Erro:", err)
    return
}
// err NÃO existe aqui fora!

// Múltiplas variáveis
if nome, idade := obterDados(); idade >= 18 {
    fmt.Printf("%s é maior de idade\n", nome)
}

// Uso real (muito comum)
if file, err := os.Open("config.txt"); err != nil {
    log.Fatal(err)
} else {
    defer file.Close()
    // usar file aqui
}

```
**Por que isso é útil? Limita o escopo da variável, evitando vazamentos e deixando explícito que ela só serve para aquela validação.**

**3. Switch Statement (Diferente de C/Java)**  

**🔴 PONTO CRÍTICO: Sem fallthrough automático!**  

```go
// Em C/Java: case "cai" para o próximo
// Em Go: case NÃO cai automaticamente

dia := 2

switch dia {
case 1:
    fmt.Println("Domingo")
case 2:
    fmt.Println("Segunda")  // Executa SÓ este
    // break é implícito, não precisa colocar
case 3:
    fmt.Println("Terça")
default:
    fmt.Println("Outro dia")
}
// Saída: "Segunda"
```

**Fallthrough explícito (quando você QUER que caia):**

```go
switch dia {
case 1:
    fmt.Println("Domingo")
    fallthrough  // ✅ Força cair para o próximo case
case 2:
    fmt.Println("Segunda também executa!")
}

```
**Switch com múltiplos valores:**  
```go
switch dia {
case 1, 7:
    fmt.Println("Fim de semana")
case 2, 3, 4, 5, 6:
    fmt.Println("Dia útil")
default:
    fmt.Println("Dia inválido")
}

```
**4. 🟡 Switch sem Expressão (switch true)**  
**Muito útil para substituir if-else longo:**  
```go
nota := 85

// Estes dois são IDÊNTICOS:

// Forma 1: switch tradicional com true
switch {
case nota >= 90:
    fmt.Println("A")
case nota >= 80:
    fmt.Println("B")
case nota >= 70:
    fmt.Println("C")
default:
    fmt.Println("Reprovado")
}

// Forma 2: switch true explícito (mesma coisa)
switch true {
case nota >= 90:
    fmt.Println("A")
// ... igual
}
```

**Vantagem: Mais limpo que vários if-else if.**

**5. Switch com Inicialização (Igual if)**  

```go
switch nota := calcularNota(); {
case nota >= 90:
    fmt.Println("Excelente")
case nota >= 70:
    fmt.Println("Bom")
default:
    fmt.Println("Precisa melhorar")
}
// nota não existe aqui fora
```

**6. Type Switch (Para interfaces)**

```go
func identificarTipo(v interface{}) {
    switch v.(type) {
    case int:
        fmt.Println("É um inteiro")
    case string:
        fmt.Println("É uma string")
    case bool:
        fmt.Println("É um booleano")
    case nil:
        fmt.Println("É nil")
    default:
        fmt.Println("Tipo desconhecido")
    }
}

// Com acesso ao valor
func processar(v interface{}) {
    switch valor := v.(type) {
    case int:
        fmt.Printf("Inteiro: %d\n", valor)
    case string:
        fmt.Printf("String: %s\n", valor)
    }
}
```
**7. Comparações Importantes**
**Em Go, if aceita APENAS booleanos:** 

```go
// ✅ Correto
if x > 0 { }

// ❌ Erro (diferente de Python/JS/C)
// if x { }  // Erro: x é int, não bool

// ✅ Correto
if x != 0 { }
```

**Comparação com nil:**
```go
var err error  // nil

if err != nil {
    fmt.Println("Tem erro")
} else {
    fmt.Println("Tudo ok")
}

// ⚠️ Atenção: slice/map vazio NÃO é nil
var s []int  // nil
s = []int{}  // NÃO é nil (é vazio mas alocado)

if s == nil {
    fmt.Println("É nil")
}
```
**📋 Tabela de Comparação: Go vs Outras Linguagens**  


```go
Característica	        Go	        C/Java	        Python	        JavaScript
{} obrigatório	        ✅	        ✅	            ❌	            ❌
Parênteses na condição	❌	        ✅	            ❌	            ✅
else na mesma linha	    ✅           ❌	            N/A	             ❌
Fallthrough automático	❌	        ✅	            N/A	             ✅(sem break)
If com init	            ✅	        ❌	            ❌	            ❌
Condição só booleana	✅	        ✅	            ✅(Truthy)	    ✅(Truthy)
```

[voltar](#menu)









---

<br><br>
<br><br>
<br><br>
<br><br><br><br><br><br><br><br><br><br>


## Goroutine

Imagine que você tem um único funcionário (a CPU) num restaurante. Ele precisa:

    Atender telefone ☎️

    Levar pratos aos clientes 🍽️

    Anotar pedidos ✍️

Sem goroutine: ele faz uma coisa de cada vez. Se atender telefone, todos os clientes esperam. Isso é lento e ineficiente.

Com goroutine: é como se esse funcionário multiplicasse (ganhasse "cópias leves" de si mesmo). Cada goroutine cuida de uma tarefa:

    Uma atende telefone

    Outra leva pratos

    Outra anota pedidos

E o "gerente" (scheduler do Go) decide quem faz o quê e por quanto tempo, trocando muito rápido entre elas. Parece que tudo acontece ao mesmo tempo, mas na verdade é intercalado (concorrente).

    🧠 Resumo leigo: goroutine é uma tarefa super leve que roda "ao mesmo tempo" que outras, sem precisar de muitos recursos.

<br><br>

---



1. O que é goroutine? 📚
🔹 Nível leigo (analogia do restaurante)

Imagine que você tem um único funcionário (a CPU) num restaurante. Ele precisa:

    Atender telefone ☎️

    Levar pratos aos clientes 🍽️

    Anotar pedidos ✍️

Sem goroutine: ele faz uma coisa de cada vez. Se atender telefone, todos os clientes esperam. Isso é lento e ineficiente.

Com goroutine: é como se esse funcionário multiplicasse (ganhasse "cópias leves" de si mesmo). Cada goroutine cuida de uma tarefa:

    Uma atende telefone

    Outra leva pratos

    Outra anota pedidos

E o "gerente" (scheduler do Go) decide quem faz o quê e por quanto tempo, trocando muito rápido entre elas. Parece que tudo acontece ao mesmo tempo, mas na verdade é intercalado (concorrente).

    🧠 Resumo leigo: goroutine é uma tarefa super leve que roda "ao mesmo tempo" que outras, sem precisar de muitos recursos.

🔹 Nível Engenharia Elétrica/Computação

Do ponto de vista de sistemas e arquitetura:

    Goroutine é uma thread leve gerenciada pelo runtime do Go, não pelo sistema operacional.

    Cada goroutine ocupa ~2KB de stack inicial (vs. 1-8MB de uma thread OS).

    O scheduler do Go (M:N scheduling) mapeia N goroutines para M threads do SO.

    Ele usa cooperação preemptiva: em Go 1.14+, goroutines são preemptíveis em pontos seguros (ex: chamadas a função ou quando o scheduler decide).

    Context switching entre goroutines é muito mais barato que entre threads do SO, pois não envolve chamadas de sistema (syscalls) e preserva menos contexto.

Fundamento elétrico subjacente: o scheduler usa timers e interrupções virtuais baseadas em sinais do sistema operacional (ex: SIGURG no Linux) para forçar preempção. O hardware (CPU) continua executando instruções sequenciais, mas a ilusão de paralelismo vem da alternância extremamente rápida.
2. No back-end, o que a goroutine faz?

No desenvolvimento back-end (APIs, microsserviços, workers):

    Manipulação de requisições HTTP concorrentes
    Cada requisição http.Request pode rodar em sua própria goroutine. O servidor web do Go (net/http) já faz isso por padrão.

    I/O não-bloqueante
    Enquanto uma goroutine aguarda dados do banco de dados, rede ou disco, o scheduler automaticamente pausa ela e executa outra. Em outras linguagens, isso exigiria async/await manual.

    Processamento assíncrono
    Ex: salvar um arquivo, enviar e-mail, processar fila — tudo rodando em goroutines separadas, sem travar a resposta imediata ao cliente.

    Padrões concorrentes
    Como pipelines (fan-in/fan-out), workers pools, timeouts com select e canais.

Exemplo simples:
go

go processarPedido(pedido) // roda em background, sem travar a resposta

3. Existe apenas no Go?

Não. O conceito de threads leves gerenciadas pelo runtime existe em outras linguagens:
Linguagem	Nome	Notas
Erlang/Elixir	Processos	Isolados, comunicação via mensagens
Kotlin	Coroutines	Leves, suportadas a nível de linguagem
Java	Virtual Threads (Project Loom)	Desde Java 21, similar a goroutines
C++	Coroutines (C++20)	Leves, mas sem scheduler automático
Rust	Tokio/async-std (tasks)	Green threads via bibliotecas assíncronas

O diferencial do Go: goroutines vêm embutidas na linguagem desde o início, com canais (chan) nativos e scheduler automático, sem necessidade de bibliotecas externas ou palavras-chave async/await.
4. O que é concorrência?

Concorrência é a capacidade de lidar com várias tarefas ao mesmo tempo, mas não necessariamente executá-las simultaneamente no exato mesmo instante.

    Concorrência ≠ Paralelismo

        Concorrência: estruturação do programa em tarefas que podem ser executadas em ordem arbitrária (intercaladas).

        Paralelismo: execução simultânea real (múltiplos núcleos/CPUs).

Exemplo:

    Concorrente: você cozinha e responde mensagens no celular, alternando entre as duas.

    Paralelo: você cozinha e alguém responde mensagens ao mesmo tempo, cada um numa CPU.

5. É chamado de concorrência apenas no Go?

Não. Concorrência é um conceito geral da ciência da computação, muito anterior ao Go:

    Anos 1960: C. A. R. Hoare (CSP – Communicating Sequential Processes, que inspirou os canais do Go).

    Sistemas operacionais: processos concorrentes, threads.

    Banco de dados: controle de concorrência (transações).

O que o Go fez foi popularizar o termo no contexto de linguagens modernas, principalmente com a famosa frase de Rob Pike:

    "Concorrência é sobre lidar com muitas coisas ao mesmo tempo. Paralelismo é sobre fazer muitas coisas ao mesmo tempo."

Portanto: concorrência não é exclusividade Go, mas o Go foi desenhado desde o início para torná-la simples e eficiente com goroutines e canais.
Resumo rápido 🧠
Pergunta	Resposta
Goroutine é só no Go?	Não, mas é nativa e com scheduler próprio
Concorrência é só no Go?	Não, conceito universal
Goroutine no back-end?	Roda requisições, I/O, tarefas assíncronas
Leigo → goroutine?	Vários atendentes leves no restaurante
Engenharia → goroutine?	Contexto leve (2KB stack), mapeamento M:N, preempção

Se quiser, posso mostrar exemplos práticos de concorrência vs paralelismo com código Go.
