# Golang
## Introdução
Este README, é uma breve anotação que fiz ao longo do mês de Abril de 2026 sobe Golang e suas particularidade. Aqui eu anotei os principais conceitos da linguagem para eu guiar um caminho solido e progressivo em Golang.
Sempre que preciso recuperar ou rever conceitos eu reviso por aqui. mas fácil saber onde está do que ficar procurando aleatóriamente pela internet.

**Sintaxe básica**

### Menu
- [Sintaxe Básica - Variáveis em Go](#sintaxe-basica)
- 
  
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
