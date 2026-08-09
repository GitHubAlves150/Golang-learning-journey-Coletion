
# Golang
## Introdução
Este README, é uma breve anotação que fiz ao longo do mês de Abril de 2026 sobe Golang e suas particularidade. Aqui eu anotei os principais conceitos da linguagem para eu guiar um caminho solido e progressivo em Golang.
Sempre que preciso recuperar ou rever conceitos eu reviso por aqui. mas fácil saber onde está do que ficar procurando aleatóriamente pela internet.

**Sintaxe básica**

### Menu
- [Sintaxe Básica - Variáveis em Go](#sintaxe-basica)
- [Operadores em Golnag](#operadores)
- [Condicionais](#condicionais)
- [Apostila](#Apostila)

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


=======================================================================================================
## Apostila
=======================================================================================================

## Guia Completo de Go (Golang)

### Sumário

- Arrays, Slices, Maps, Structs e Ponteiros
- Tratamento de Erros, Defer, Panic e Interfaces
- JSON, Servidor HTTP, Módulos e Banco de Dados
- Testes, Fuzzing, Generics e Concorrência
- Arquitetura de Software e Projeto Prático

---

## 1. Arrays, Slices, Maps, Structs e Ponteiros

### 1.1 Arrays

Arrays em Go possuem tamanho fixo determinado no momento da compilação. O tamanho faz parte da assinatura do tipo, o que significa que [3]int e [5]int são tipos incompatíveis.

```bash
package main

import "fmt"

func main() {
    // Declaração tradicional
    var numbers [3]int
    numbers[0] = 10
    numbers[1] = 20
    numbers[2] = 30

    // Declaração literal
    colors := [3]string{"Vermelho", "Verde", "Azul"}

    fmt.Printf("Array de Números: %v (Tipo: %T)\n", numbers, numbers)
    fmt.Printf("Array de Cores: %v (Tamanho: %d)\n", colors, len(colors))
}

Saída:
bash

Array de Números: [10 20 30] (Tipo: [3]int)
Array de Cores: [Vermelho Verde Azul] (Tamanho: 3)

💡 Caso de Uso Prático: Leitura de headers com tamanho fixo (ex: checksum MD5/SHA256, blocos de criptografia ou buffers de IO de baixo nível).
1.2 Slices

Slices são abstrações dinâmicas sobre arrays internos. Eles contêm três componentes: um ponteiro para o array subjacente, o comprimento (len) e a capacidade (cap).
bash

package main

import "fmt"

func main() {
    // Inicialização de Slice dinâmico com make
    // make([]Tipo, len, cap)
    items := make([]string, 0, 3)

    items = append(items, "Servidor 01")
    items = append(items, "Servidor 02")
    items = append(items, "Servidor 03")

    fmt.Printf("Slices: %v | Len: %d | Cap: %d\n", items, len(items), cap(items))

    // Disparo de realocação de memória ao exceder a capacidade
    items = append(items, "Servidor 04")
    fmt.Printf("Após Expansão: %v | Len: %d | Cap: %d\n", items, len(items), cap(items))

    // Fatia de slice (Slicing)
    subGroup := items[1:3]
    fmt.Println("Subgrupo (1:3):", subGroup)
}

Saída:
bash

Slices: [Servidor 01 Servidor 02 Servidor 03] | Len: 3 | Cap: 3
Após Expansão: [Servidor 01 Servidor 02 Servidor 03 Servidor 04] | Len: 4 | Cap: 6
Subgrupo (1:3): [Servidor 02 Servidor 03]

💡 Caso de Uso Prático: Listas dinâmicas de itens em APIs, como paginação de registros recuperados do banco de dados.
1.3 Maps

map é uma estrutura de dados de chave-valor não ordenada (tabela hash).
bash

package main

import "fmt"

func main() {
    rates := map[string]float64{
        "USD": 5.05,
        "EUR": 5.45,
    }

    rates["GBP"] = 6.30

    // Leitura segura com idiom "comma ok"
    val, ok := rates["JPY"]
    if !ok {
        fmt.Println("Moeda JPY não encontrada no cache!")
    } else {
        fmt.Println("Cotação JPY:", val)
    }

    // Remoção de elemento
    delete(rates, "GBP")

    fmt.Println("Tabela de Cotações:", rates)
}

Saída:
bash

Moeda JPY não encontrada no cache!
Tabela de Cotações: map[EUR:5.45 USD:5.05]

⚠️ Aviso de Concorrência: map nativo do Go não é seguro para escrita/leitura concorrente. Utilize sync.RWMutex ou sync.Map para acesso por múltiplas Goroutines.
1.4 Structs, Métodos e Ponteiros

Structs agrupam dados e comportamentos em tipos customizados.
bash

package main

import "fmt"

type Account struct {
    Owner   string  `json:"owner"`
    Balance float64 `json:"balance"`
}

// Receiver por valor (não altera a instância original)
func (a Account) Display() string {
    return fmt.Sprintf("Titular: %s | Saldo: R$%.2f", a.Owner, a.Balance)
}

// Receiver por ponteiro (modifica o estado interno da instância)
func (a *Account) Deposit(amount float64) {
    if amount > 0 {
        a.Balance += amount
    }
}

func main() {
    acc := Account{Owner: "Ana Silva", Balance: 1000.00}
    fmt.Println("Antes:", acc.Display())

    acc.Deposit(500.50)
    fmt.Println("Após Depósito:", acc.Display())
}

Saída:
bash

Antes: Titular: Ana Silva | Saldo: R$1000.00
Após Depósito: Titular: Ana Silva | Saldo: R$1500.50

2. Tratamento de Erros, Defer, Panic e Interfaces
2.1 Tratamento Idiomático de Erros e Error Wrapping

Em Go, erros são valores explicitamente retornados por funções, e não exceções capturadas silenciosamente.
bash

package main

import (
    "errors"
    "fmt"
)

var ErrNotFound = errors.New("recurso não encontrado")

type DatabaseError struct {
    Op  string
    Err error
}

func (e *DatabaseError) Error() string {
    return fmt.Sprintf("erro na operação '%s': %v", e.Op, e.Err)
}

func (e *DatabaseError) Unwrap() error {
    return e.Err
}

func FetchUser(id int) (string, error) {
    if id <= 0 {
        return "", &DatabaseError{
            Op:  "SELECT",
            Err: ErrNotFound,
        }
    }
    return "Carlos", nil
}

func main() {
    _, err := FetchUser(-1)
    if err != nil {
        fmt.Println("Erro capturado:", err)

        // Verificação semântica com errors.Is
        if errors.Is(err, ErrNotFound) {
            fmt.Println("Causa raiz: O recurso realmente não existe.")
        }
    }
}

Saída:
bash

Erro capturado: erro na operação 'SELECT': recurso não encontrado
Causa raiz: O recurso realmente não existe.

2.2 Use Defer com Sabedoria

A instrução defer adia a execução de uma função até que a função externa retorne. É ideal para limpeza de recursos.
bash

package main

import (
    "fmt"
    "sync"
)

type SafeCounter struct {
    mu    sync.Mutex
    value int
}

func (c *SafeCounter) Inc() {
    c.mu.Lock()
    defer c.mu.Unlock() // Garante o destravamento ao sair do escopo
    c.value++
}

func main() {
    counter := SafeCounter{}
    counter.Inc()
    fmt.Println("Valor do Contador:", counter.value)
}

Saída:
bash

Valor do Contador: 1

2.3 Interfaces Desacopladas (Satisfação Implícita)

Interfaces especificam contratos de comportamento. Qualquer tipo que possua os métodos exigidos satisfaz a interface implicitamente.
bash

package main

import (
    "context"
    "fmt"
)

type NotificationSender interface {
    Send(ctx context.Context, message string) error
}

type EmailService struct{}

func (e EmailService) Send(ctx context.Context, message string) error {
    fmt.Println("[EMAIL SENT]:", message)
    return nil
}

type SMSService struct{}

func (s SMSService) Send(ctx context.Context, message string) error {
    fmt.Println("[SMS SENT]:", message)
    return nil
}

func NotifyUser(ctx context.Context, sender NotificationSender, msg string) {
    _ = sender.Send(ctx, msg)
}

func main() {
    ctx := context.Background()

    email := EmailService{}
    sms := SMSService{}

    NotifyUser(ctx, email, "Bem-vindo ao sistema!")
    NotifyUser(ctx, sms, "Seu código de verificação é 8492.")
}

Saída:
bash

[EMAIL SENT]: Bem-vindo ao sistema!
[SMS SENT]: Seu código de verificação é 8492.

3. JSON, Servidor HTTP, Módulos e Banco de Dados
3.1 APIs RESTful Idiomáticas com net/http
bash

package main

import (
    "encoding/json"
    "net/http"
    "time"
)

type UserRequest struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}

type UserResponse struct {
    ID        int       `json:"id"`
    Name      string    `json:"name"`
    Email     string    `json:"email"`
    CreatedAt time.Time `json:"created_at"`
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
        return
    }

    var req UserRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "JSON inválido", http.StatusBadRequest)
        return
    }

    resp := UserResponse{
        ID:        101,
        Name:      req.Name,
        Email:     req.Email,
        CreatedAt: time.Now(),
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    _ = json.NewEncoder(w).Encode(resp)
}

Exemplo de Chamada HTTP cURL:
bash

curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name": "Eduardo", "email": "eduardo@example.com"}'

Saída Esperada (JSON):
bash

{
  "id": 101,
  "name": "Eduardo",
  "email": "eduardo@example.com",
  "created_at": "2026-08-09T12:57:00Z"
}

3.2 Persistência Segura com database/sql

Consultas devem utilizar Prepared Statements ou Placeholders para evitar SQL Injection.
bash

package main

import (
    "context"
    "database/sql"
    "time"
)

type User struct {
    ID    int
    Name  string
    Email string
}

func GetUserByID(ctx context.Context, db *sql.DB, id int) (*User, error) {
    // Timeout de context garante cancelamento de queries presas
    ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
    defer cancel()

    query := `SELECT id, name, email FROM users WHERE id = $1`

    var user User
    err := db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Name, &user.Email)
    if err != nil {
        return nil, err
    }

    return &user, nil
}

4. Testes, Fuzzing, Generics e Concorrência
4.1 Testes Orientados a Tabela (Table-Driven Tests)
bash

package math_test

import "testing"

func Abs(n int) int {
    if n < 0 {
        return -n
    }
    return n
}

func TestAbs(t *testing.T) {
    tests := []struct {
        name     string
        input    int
        expected int
    }{
        {name: "Número Positivo", input: 10, expected: 10},
        {name: "Número Negativo", input: -5, expected: 5},
        {name: "Zero", input: 0, expected: 0},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := Abs(tt.input)
            if got != tt.expected {
                t.Errorf("Abs(%d) = %d; esperado %d", tt.input, got, tt.expected)
            }
        })
    }
}

Execução:
bash

go test -v ./...

4.2 Generics (Go 1.18+)

Permite a reusabilidade de algoritmos sem perda de segurança de tipos.
bash

package main

import "fmt"

// Constraints utilizando comparable
func MapKeys[K comparable, V any](m map[K]V) []K {
    keys := make([]K, 0, len(m))
    for k := range m {
        keys = append(keys, k)
    }
    return keys
}

func main() {
    intMap := map[string]int{"a": 1, "b": 2}
    keys := MapKeys(intMap)

    fmt.Println("Chaves extraídas:", keys)
}

Saída:
bash

Chaves extraídas: [a b]

4.3 Concorrência Avançada: Worker Pools & Channels

O ecossistema Go gerencia concorrência com Goroutines e Channels sob o lema: "Não comunique compartilhando memória; compartilhe memória se comunicando."
bash

package main

import (
    "fmt"
    "sync"
    "time"
)

type Job struct {
    ID    int
    Value int
}

type Result struct {
    Job    Job
    Square int
}

func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
    defer wg.Done()
    for job := range jobs {
        // Processamento simulado
        time.Sleep(50 * time.Millisecond)
        results <- Result{Job: job, Square: job.Value * job.Value}
    }
}

func main() {
    const numJobs = 5
    const numWorkers = 2

    jobs := make(chan Job, numJobs)
    results := make(chan Result, numJobs)
    var wg sync.WaitGroup

    // Inicializa Workers
    for w := 1; w <= numWorkers; w++ {
        wg.Add(1)
        go worker(w, jobs, results, &wg)
    }

    // Envia Jobs
    for j := 1; j <= numJobs; j++ {
        jobs <- Job{ID: j, Value: j * 10}
    }
    close(jobs)

    // Aguarda conclusão dos workers e fecha canal de resultados
    go func() {
        wg.Wait()
        close(results)
    }()

    // Processa Resultados
    for res := range results {
        fmt.Printf("Job #%d processado! Entrada: %d | Quadrado: %d\n", res.Job.ID, res.Job.Value, res.Square)
    }
}

Saída:
bash

Job #1 processado! Entrada: 10 | Quadrado: 100
Job #2 processado! Entrada: 20 | Quadrado: 400
Job #3 processado! Entrada: 30 | Quadrado: 900
Job #4 processado! Entrada: 40 | Quadrado: 1600
Job #5 processado! Entrada: 50 | Quadrado: 2500

5. Arquitetura de Software e Projeto Prático
5.1 Estrutura de Diretórios Padronizada (Clean Standard)
bash

meu-projeto/
├── cmd/
│   └── api/
│       └── main.go          # Ponto de entrada do executável
├── internal/
│   ├── domain/              # Entidades puras e regras de negócio
│   ├── usecase/             # Casos de uso e orquestração
│   ├── handler/             # Controllers HTTP/gRPC
│   └── repository/          # Acesso a banco de dados e APIs externas
├── migrations/              # Arquivos SQL de migração de banco
├── go.mod
└── go.sum

5.2 Roteiro de Qualidade e Comandos CLI

Para garantir os padrões de produção em Go, utilize a sequência de validação:
bash

# 1. Formatação do código de acordo com o padrão Go
gofmt -w .

# 2. Análise estática contra bugs comuns
go vet ./...

# 3. Execução da suíte de testes com Race Detector ativado
go test -race -v ./...

# 4. Verificação de cobertura de código
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

📚 Recursos Adicionais

    Documentação Oficial do Go

    Effective Go

    Go by Example

    Awesome Go

text

