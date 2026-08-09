# 📘 Apostila Completa de Go (Golang)

Esta apostila reúne, de forma organizada e progressiva, os principais conceitos da linguagem Go (Golang): desde a sintaxe básica até tópicos avançados como concorrência, generics e arquitetura de projetos. O objetivo é servir como material de consulta rápida e também como guia de estudo sequencial para quem está construindo uma base sólida na linguagem.

## Sumário

- [Introdução](#introdução)
- [1. Sintaxe Básica — Variáveis, Tipos e Constantes](#1-sintaxe-básica--variáveis-tipos-e-constantes)
- [2. Operadores em Go](#2-operadores-em-go)
- [3. Estruturas Condicionais](#3-estruturas-condicionais)
- [4. Arrays, Slices, Maps, Structs e Ponteiros](#4-arrays-slices-maps-structs-e-ponteiros)
- [5. Tratamento de Erros, Defer, Panic e Interfaces](#5-tratamento-de-erros-defer-panic-e-interfaces)
- [6. Concorrência: Goroutines e Channels](#6-concorrência-goroutines-e-channels)
- [7. JSON, Servidor HTTP e Banco de Dados](#7-json-servidor-http-e-banco-de-dados)
- [8. Testes, Generics e Ferramentas de Qualidade](#8-testes-generics-e-ferramentas-de-qualidade)
- [9. Arquitetura de Projetos em Go](#9-arquitetura-de-projetos-em-go)
- [Recursos Adicionais](#recursos-adicionais)

---

## Introdução

Este material nasceu de anotações de estudo sobre Go e foi reorganizado em formato de apostila para consulta contínua. A ideia central é simples: **é mais fácil saber onde procurar do que procurar aleatoriamente pela internet**. Cada seção traz:

- Uma explicação conceitual do "porquê" a linguagem funciona daquele jeito;
- Exemplos de código comentados, prontos para copiar e rodar;
- A saída esperada de cada exemplo, para você validar seu entendimento sem precisar compilar;
- Avisos (`⚠️`) sobre armadilhas comuns e diferenças em relação a outras linguagens (C, Java, Python, JavaScript).

> 💡 Dica de uso: se você já programa em outra linguagem, preste atenção especial aos blocos marcados como **🔴 CRÍTICO** — eles indicam pontos onde Go se comporta de forma diferente do que você provavelmente está acostumado.

[Voltar ao topo](#sumário)

---

## 1. Sintaxe Básica — Variáveis, Tipos e Constantes

### 🎯 Objetivo

Entender como declarar, inicializar e usar variáveis em Go, incluindo particularidades como **zero values**, inferência de tipo e escopo. Go é uma linguagem **estaticamente tipada** e **compilada**, o que significa que o tipo de cada variável é conhecido em tempo de compilação — isso traz mais segurança, mas exige um pouco mais de disciplina na hora de declarar valores.

### 1.1 Declaração com Inicialização

A forma mais explícita de declarar uma variável em Go usa a palavra-chave `var`, seguida do nome, do tipo e (opcionalmente) do valor inicial.

```bash
var nome string = "João"
var idade int = 30
var ativo bool = true

// Múltiplas variáveis do mesmo tipo
var x, y int = 10, 20

// Bloco de declaração (fica mais legível quando há várias variáveis)
var (
    nome   string = "Maria"
    cidade string = "São Paulo"
)
```

Essa forma é útil principalmente em três situações: quando você quer deixar o tipo explícito por clareza, quando está declarando uma variável **no escopo do pacote** (fora de função, onde `:=` não é permitido) e quando você precisa da variável com o **zero value**, sem atribuir nada ainda.

### 1.2 Inferência de Tipo (`:=` — operador curto)

Dentro de funções, é muito mais comum usar o operador `:=`, que declara **e** atribui ao mesmo tempo, deixando o compilador inferir o tipo a partir do valor:

```bash
nome := "Carlos"        // string
idade := 25              // int
preco := 99.99           // float64
ativo := true            // bool

// Declaração múltipla
nome, idade := "Ana", 28
x, y := 10, 20
```

⚠️ **Regra importante:** `:=` só pode ser usado **dentro de funções**. No escopo de pacote (fora de qualquer função), é obrigatório usar `var`.

```bash
package main

// nome := "Erro" // ❌ não compila: := não é permitido fora de função
var nome = "Correto" // ✅ funciona (var permite inferência também)

func main() {}
```

### 1.3 Zero Values (Valores Padrão)

Diferente de linguagens como JavaScript ou Python, Go **não tem `null`/`undefined`**. Toda variável declarada e não inicializada recebe automaticamente um **zero value** de acordo com seu tipo — isso evita boa parte dos erros de "referência nula" comuns em outras linguagens.

| Tipo | Zero Value |
|---|---|
| `int`, `int8`, `int64`, etc. | `0` |
| `float32`, `float64` | `0.0` |
| `bool` | `false` |
| `string` | `""` (string vazia) |
| `pointer`, `slice`, `map`, `chan`, `func`, `interface` | `nil` |

```bash
package main

import "fmt"

func main() {
    var a int       // 0
    var b string    // ""
    var c bool      // false
    var d *int      // nil

    fmt.Printf("a=%d, b='%s', c=%t, d=%v\n", a, b, c, d)
}
```

**Saída:**

```bash
a=0, b='', c=false, d=<nil>
```

### 1.4 Tipos Básicos em Go

Go oferece um conjunto rico de tipos numéricos, o que permite escolher o tamanho exato de acordo com a necessidade de memória e desempenho — algo especialmente relevante em sistemas embarcados, protocolos binários e APIs de alta performance.

```bash
// Inteiros com sinal (escolha o menor necessário)
var i int     // depende da arquitetura (32 ou 64 bits)
var i8 int8   // -128 a 127
var i16 int16 // -32.768 a 32.767
var i32 int32 // -2.147.483.648 a 2.147.483.647
var i64 int64 // -9.223.372.036.854.775.808 a 9.223.372.036.854.775.807

// Inteiros sem sinal (unsigned)
var u uint    // 0 a 4.294.967.295 (32 bits) ou mais
var u8 uint8  // 0 a 255 (também chamado de "byte")
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

💡 **Caso de uso prático:** use `rune` sempre que precisar iterar caractere a caractere sobre uma string que pode conter acentos ou emojis, já que `string` em Go é internamente uma sequência de bytes UTF-8 e não de caracteres.

```bash
package main

import "fmt"

func main() {
    palavra := "café"
    for i, r := range palavra {
        fmt.Printf("índice=%d rune=%c\n", i, r)
    }
}
```

**Saída:**

```bash
índice=0 rune=c
índice=1 rune=a
índice=2 rune=f
índice=3 rune=é
```

### 1.5 Conversão de Tipos (Type Conversion)

Go é rigoroso quanto a tipos: **não existe conversão implícita**. Mesmo entre `int` e `float64`, você precisa converter explicitamente.

```bash
package main

import "fmt"

func main() {
    var x int = 10
    var y float64 = float64(x) // ✅ conversão explícita

    var a int = 10
    var b int64 = int64(a) // ✅ ok

    fmt.Println(y, b)

    // ❌ Isso NÃO compila:
    // var c float64 = a
    // erro: cannot use a (variable of type int) as float64 value in variable declaration
}
```

**Saída:**

```bash
10 10
```

### 1.6 Constantes e `iota`

Constantes são valores conhecidos em tempo de compilação e imutáveis durante a execução. Go tem um recurso exclusivo chamado `iota`, um contador automático usado dentro de blocos `const` — extremamente útil para criar enumerações.

```bash
package main

import "fmt"

// Constantes tipadas
const Pi float64 = 3.14159
const Nome string = "GoLang"

// Constantes não tipadas (mais flexíveis, adaptam-se ao contexto de uso)
const (
    SegundosPorMinuto = 60
    MinutosPorHora    = 60
    HorasPorDia       = 24
)

// iota - enumerador automático, incrementado a cada linha do bloco const
type DiaSemana int

const (
    Domingo DiaSemana = iota // 0
    Segunda                  // 1
    Terca                    // 2
    Quarta                   // 3
    Quinta                   // 4
    Sexta                    // 5
    Sabado                   // 6
)

func main() {
    fmt.Println(Pi, Nome)
    fmt.Println(SegundosPorMinuto * MinutosPorHora * HorasPorDia, "segundos em um dia")
    fmt.Println("Hoje é", Quarta)
}
```

**Saída:**

```bash
3.14159 GoLang
86400 segundos em um dia
Hoje é 3
```

> 🔵 Observação: como `DiaSemana` não implementa o método `String()`, o `fmt.Println` imprime o valor numérico subjacente (`3`). Para imprimir "Quarta" em vez de `3`, seria necessário implementar a interface `Stringer` — assunto abordado na seção de [Interfaces](#54-interfaces).

### 1.7 Escopo de Variáveis

Go segue um modelo de escopo por blocos, semelhante ao de C, mas com uma particularidade importante: o **shadowing** (sombreamento), quando uma variável local "esconde" temporariamente uma variável de escopo mais amplo com o mesmo nome.

```bash
package main

import "fmt"

var global = "acessível em todo o pacote" // escopo de pacote

func main() {
    local := "acessível só dentro da função" // escopo de função

    if true {
        bloco := "acessível só dentro deste bloco" // escopo de bloco
        fmt.Println(bloco) // ✅ ok
    }

    // fmt.Println(bloco) // ❌ erro: bloco não definido aqui fora

    // Shadowing (sombreamento de variável)
    global := "variável local com mesmo nome" // ⚠️ cria uma NOVA variável, não altera a global
    fmt.Println(local, global)
}
```

**Saída:**

```bash
acessível só dentro deste bloco
acessível só dentro da função variável local com mesmo nome
```

⚠️ O `global :=` dentro de `main` não altera a variável de pacote `global` — ele cria uma variável completamente nova, local à função. Esse é um erro sutil e comum: usar `:=` sem perceber que já existe uma variável com esse nome em outro escopo.

### 1.8 Ponteiros — Introdução

Ponteiros armazenam o **endereço de memória** de uma variável, em vez do valor em si. Em Go eles são mais seguros que em C: não existe aritmética de ponteiros, o que elimina uma classe inteira de bugs de acesso indevido à memória.

```bash
package main

import "fmt"

func main() {
    x := 10
    ptr := &x // ptr armazena o endereço de memória de x

    fmt.Println(x)    // 10
    fmt.Println(ptr)  // endereço, ex: 0xc000012088
    fmt.Println(*ptr) // 10 (dereferencing — acessa o valor apontado)

    *ptr = 20 // altera o valor de x através do ponteiro
    fmt.Println(x) // 20
}
```

**Saída:**

```bash
10
0xc000012088
10
20
```

[Voltar ao Sumário](#sumário)

---

## 2. Operadores em Go

### Visão Geral

| Operador | Particularidade em Go | Importância |
|---|---|---|
| `=` vs `:=` | `:=` declara **e** atribui (exclusivo do Go) | 🔴 CRÍTICO |
| `++` / `--` | São instruções, não expressões | 🔴 CRÍTICO |
| `==` / `!=` | Compara structs campo a campo | 🟡 IMPORTANTE |
| `&` / `*` | Ponteiros (sem aritmética) | 🟡 IMPORTANTE |
| `&&` / `\|\|` | Short-circuit (igual a outras linguagens) | 🟢 BÁSICO |
| `<<` / `>>` | Shift, com atenção a tipos unsigned | 🟢 BÁSICO |
| `&^` (AND NOT) | Bit clear — exclusivo do Go | 🔵 DIFERENTE |

### 2.1 Operadores Aritméticos (iguais a C/Java)

```bash
package main

import "fmt"

func main() {
    a := 10 + 5   // Adição
    b := 10 - 5   // Subtração
    c := 10 * 5   // Multiplicação
    d := 10 / 3   // Divisão de inteiros (trunca!)
    e := 10.0 / 3 // Divisão com float
    f := 10 % 3   // Módulo (resto)

    fmt.Println(a, b, c, d, e, f)
}
```

**Saída:**

```bash
15 5 50 3 3.3333333333333335 1
```

⚠️ **Atenção:** divisão entre inteiros trunca o resultado, não arredonda:

```bash
package main

import "fmt"

func main() {
    fmt.Println(10 / 3)   // 3   (inteiro, trunca)
    fmt.Println(10 / 3.0) // 3.3333333333333335 (float)
}
```

**Saída:**

```bash
3
3.3333333333333335
```

### 2.2 Operadores de Atribuição

O que **não** funciona em Go, mas é comum em outras linguagens:

```bash
// Isso NÃO compila como "incremento composto invertido":
// x =+ 1    // Go interpreta como x = (+1), não como x += 1
// x =- 1    // Go interpreta como x = (-1), não como x -= 1

// Operador ternário não existe em Go:
// resultado = (idade > 18) ? "maior" : "menor" // ❌ não compila
```

O que funciona (igual a C/Java):

```bash
package main

import "fmt"

func main() {
    x := 10
    x += 5 // x = 15
    x *= 2 // x = 30
    x -= 6 // x = 24
    x /= 4 // x = 6
    x %= 4 // x = 2

    fmt.Println(x)
}
```

**Saída:**

```bash
2
```

🔴 **CRÍTICO: o operador curto `:=`** (exclusivo do Go)

```bash
package main

import "fmt"

func main() {
    // := declara E atribui (só dentro de funções)
    nome := "João" // declara variável nome do tipo string
    idade := 30    // declara idade como int

    // Equivalente a:
    // var nome string = "João"
    // var idade int = 30

    // Reatribuição com := é permitida se PELO MENOS UMA variável for nova
    x := 10        // declara x
    x, y := 20, 30 // x é reatribuído, y é uma nova variável (✅ ok)
    // x := 40     // ❌ erro: x já foi declarada neste escopo

    fmt.Println(nome, idade, x, y)
}
```

**Saída:**

```bash
João 30 20 30
```

### 2.3 `++` e `--` são INSTRUÇÕES, não expressões

Esta é uma das diferenças mais surpreendentes para quem vem de C, Java, Python ou JavaScript: em Go, `++` e `--` **não retornam valor** e não podem ser usados dentro de uma expressão.

```bash
package main

import "fmt"

func main() {
    // ✅ Isso funciona (instrução independente)
    x := 5
    x++ // x = 6
    x-- // x = 5

    // ❌ Isso NÃO funciona em Go:
    // y := x++         // erro: x++ usado como expressão
    // fmt.Println(x++) // erro
    // return x++       // erro

    // ✅ Forma correta:
    x++
    y := x // agora sim, y recebe o valor já incrementado

    fmt.Println(x, y)
}
```

**Saída:**

```bash
6 6
```

### 2.4 Operadores Relacionais (Comparação)

```bash
==   Igual
!=   Diferente
<    Menor
<=   Menor ou igual
>    Maior
>=   Maior ou igual
```

🟡 **Comparação de structs:** duas structs são iguais se **todos os campos** forem iguais.

```bash
package main

import "fmt"

type Pessoa struct {
    Nome  string
    Idade int
}

func main() {
    p1 := Pessoa{"João", 30}
    p2 := Pessoa{"João", 30}
    p3 := Pessoa{"Maria", 25}

    fmt.Println(p1 == p2) // true  (compara campo a campo)
    fmt.Println(p1 == p3) // false

    // ⚠️ Slices, maps e functions NÃO são comparáveis com == (exceto com nil)
    // slice1 := []int{1, 2, 3}
    // slice2 := []int{1, 2, 3}
    // fmt.Println(slice1 == slice2) // ❌ erro de compilação!
}
```

**Saída:**

```bash
true
false
```

Comparação com `nil`:

```bash
package main

import "fmt"

func main() {
    var s []int
    var m map[string]int
    var p *int

    fmt.Println(s == nil) // true (slice nil)
    fmt.Println(m == nil) // true (map nil)
    fmt.Println(p == nil) // true (ponteiro nil)
}
```

**Saída:**

```bash
true
true
true
```

### 2.5 Operadores Lógicos (Short-circuit)

```bash
package main

import "fmt"

func valida(x int) bool {
    fmt.Println("Validando...")
    return x > 0
}

func main() {
    if false && valida(10) { // valida() NÃO é chamada (curto-circuito)
        fmt.Println("nunca executa")
    }

    if true || valida(10) { // valida() NÃO é chamada (curto-circuito)
        fmt.Println("sempre executa")
    }
}
```

**Saída:**

```bash
sempre executa
```

Note que `"Validando..."` **nunca** aparece na saída: em ambos os casos, o Go conseguiu determinar o resultado da expressão sem precisar avaliar `valida(10)`.

### 2.6 Operadores Bit a Bit (para performance e flags)

```bash
package main

import "fmt"

func main() {
    x := 0b1100 // 12
    y := 0b1010 // 10

    fmt.Printf("%04b\n", x&y)  // AND  -> 1000 (8)
    fmt.Printf("%04b\n", x|y)  // OR   -> 1110 (14)
    fmt.Printf("%04b\n", x^y)  // XOR  -> 0110 (6)
    fmt.Printf("%04b\n", x&^y) // AND NOT (bit clear) -> 0100 (4)

    // Shift
    fmt.Println(1 << 10)  // 1024 (desloca à esquerda = multiplica por 2^n)
    fmt.Println(1024 >> 3) // 128 (desloca à direita = divide por 2^n)
}
```

**Saída:**

```bash
1000
1110
0110
0100
1024
128
```

**Uso prático do `&^` (bit clear) — sistema de permissões com flags:**

```bash
package main

import "fmt"

const (
    Read  = 1 << iota // 1 (0b001)
    Write             // 2 (0b010)
    Execute           // 4 (0b100)
)

func main() {
    permissoes := Read | Write // 0b011 (3)
    fmt.Printf("Permissões iniciais: %03b\n", permissoes)

    // Remover a permissão de escrita usando &^
    permissoes = permissoes &^ Write // 0b001 (1) - apenas Read
    fmt.Printf("Após remover Write:  %03b\n", permissoes)
}
```

**Saída:**

```bash
Permissões iniciais: 011
Após remover Write:  001
```

### 2.7 Operadores de Ponteiros

```bash
package main

import "fmt"

func main() {
    x := 42
    ptr := &x // ptr recebe o endereço de x

    fmt.Println(ptr)  // ex: 0xc000012088
    fmt.Println(*ptr) // 42 (valor apontado)

    *ptr = 100 // altera o valor de x através do ponteiro
    fmt.Println(x) // 100
}
```

**Saída:**

```bash
0xc000012088
42
100
```

⚠️ Não há aritmética de ponteiros em Go (diferente de C):

```bash
// ptr++ // ❌ erro: invalid operation (aritmética de ponteiros não permitida)
```

### 2.8 Precedência de Operadores

| Precedência | Operadores |
|---|---|
| 1 (mais alta) | `*` `/` `%` `<<` `>>` `&` `&^` |
| 2 | `+` `-` `\|` `^` |
| 3 | `==` `!=` `<` `<=` `>` `>=` |
| 4 | `&&` |
| 5 (mais baixa) | `\|\|` |

Use parênteses sempre que houver dúvida — a legibilidade vale mais que economizar caracteres:

```bash
package main

import "fmt"

func main() {
    x := 1 + 2*3   // 7 (multiplicação tem precedência), não 9
    y := (1 + 2) * 3 // 9 (parênteses deixam a intenção explícita)

    fmt.Println(x, y)
}
```

**Saída:**

```bash
7 9
```

### 2.9 Operador Especial: Vírgula em Declarações

```bash
package main

import "fmt"

func obterNome() (string, error) {
    return "Carlos", nil
}

func main() {
    // Declaração múltipla
    x, y := 10, 20

    // Swap (troca de valores, sem variável temporária)
    a, b := 5, 10
    a, b = b, a // a=10, b=5

    // Ignorar valores com blank identifier "_"
    nome, _ := obterNome() // ignora o segundo valor de retorno (o erro)

    fmt.Println(x, y, a, b, nome)
}
```

**Saída:**

```bash
10 20 10 5 Carlos
```

### 2.10 Erros Comuns com Operadores

```bash
package main

func main() {
    // 1. Confundir = com :=
    // x = 10     // ❌ erro se x não foi declarado antes
    x := 10       // ✅ correto (declara e atribui)

    // 2. Usar ++ como expressão
    // y = x++    // ❌ erro

    // 3. Comparar tipos diferentes sem conversão
    // if 10 == 10.0 {} // ❌ erro: tipos incompatíveis (int vs float64)

    // 4. Divisão de inteiro com float sem conversão
    var a int = 10
    var b float64 = 3.0
    // c := a / b       // ❌ erro: tipos diferentes
    c := float64(a) / b // ✅ conversão explícita

    // 5. Operador ternário não existe em Go
    // resultado := idade > 18 ? "maior" : "menor" // ❌ erro

    // Forma correta, usando if/else:
    idade := 20
    var resultado string
    if idade > 18 {
        resultado = "maior"
    } else {
        resultado = "menor"
    }

    _ = x
    _ = c
    _ = resultado
}
```

[Voltar ao Sumário](#sumário)

---

## 3. Estruturas Condicionais

### Resumo Rápido

| Recurso | Sintaxe | Particularidade |
|---|---|---|
| `if` | `if condicao { }` | Chaves `{}` obrigatórias, mesmo com 1 linha |
| `if-else` | `if cond { } else { }` | `else` deve ficar na MESMA linha da `}` |
| `if` com init | `if x := 10; x > 0 { }` | A variável só existe dentro do `if`/`else` |
| `switch` | `switch valor { case 1: }` | `fallthrough` é opcional (não é automático) |
| `switch true` | `switch { case x > 0: }` | Switch sem expressão = `switch true` |
| type switch | `switch v.(type) { }` | Usado para inspecionar o tipo de uma interface |

### 3.1 If Statement (Chaves Obrigatórias)

```bash
package main

import "fmt"

func main() {
    idade := 20

    // ✅ Correto
    if idade >= 18 {
        fmt.Println("Maior de idade")
    }

    // ❌ Erro - chaves são obrigatórias MESMO com apenas 1 linha
    // if idade >= 18 fmt.Println("Maior") // NÃO COMPILA

    // ✅ else deve ficar na MESMA linha da chave de fechamento
    if idade >= 18 {
        fmt.Println("Maior")
    } else {
        fmt.Println("Menor")
    }

    // ✅ Encadeamento com else if
    nota := 6
    if nota >= 7 {
        fmt.Println("Aprovado")
    } else if nota >= 5 {
        fmt.Println("Recuperação")
    } else {
        fmt.Println("Reprovado")
    }
}
```

**Saída:**

```bash
Maior de idade
Maior
Recuperação
```

### 3.2 If com Inicialização (Init Statement)

Este é um recurso característico do Go que você vai usar constantemente, especialmente ao tratar erros:

```bash
package main

import (
    "errors"
    "fmt"
)

func processar() error {
    return errors.New("falha simulada")
}

func main() {
    // A variável err existe SOMENTE dentro deste bloco if/else
    if err := processar(); err != nil {
        fmt.Println("Erro:", err)
    }
    // err NÃO existe aqui fora!

    // Múltiplas variáveis no init statement
    nome, idade := "Ana", 20
    if nome, idade := nome, idade; idade >= 18 {
        fmt.Printf("%s é maior de idade\n", nome)
    }
}
```

**Saída:**

```bash
Erro: falha simulada
Ana é maior de idade
```

**Por que isso é útil?** Limita o escopo da variável ao contexto onde ela é relevante, evitando "vazamento" de variáveis temporárias (como `err`) para o resto da função e deixando explícito que ela só serve para aquela validação.

### 3.3 Switch Statement (Diferente de C/Java)

🔴 **Ponto crítico: em Go não existe fallthrough automático.** Cada `case` já se comporta como se tivesse um `break` implícito.

```bash
package main

import "fmt"

func main() {
    dia := 2

    switch dia {
    case 1:
        fmt.Println("Domingo")
    case 2:
        fmt.Println("Segunda") // Executa SÓ este case
        // break é implícito, não é preciso escrever
    case 3:
        fmt.Println("Terça")
    default:
        fmt.Println("Outro dia")
    }
}
```

**Saída:**

```bash
Segunda
```

**Fallthrough explícito** (quando você realmente quer que um `case` "caia" para o próximo):

```bash
package main

import "fmt"

func main() {
    dia := 1

    switch dia {
    case 1:
        fmt.Println("Domingo")
        fallthrough // ✅ força a execução do próximo case também
    case 2:
        fmt.Println("Segunda também executa!")
    }
}
```

**Saída:**

```bash
Domingo
Segunda também executa!
```

**Switch com múltiplos valores em um mesmo `case`:**

```bash
package main

import "fmt"

func main() {
    dia := 3

    switch dia {
    case 1, 7:
        fmt.Println("Fim de semana")
    case 2, 3, 4, 5, 6:
        fmt.Println("Dia útil")
    default:
        fmt.Println("Dia inválido")
    }
}
```

**Saída:**

```bash
Dia útil
```

### 3.4 Switch sem Expressão (`switch true`)

Muito útil para substituir uma cadeia longa de `if-else if`, tornando o código mais legível:

```bash
package main

import "fmt"

func main() {
    nota := 85

    // Forma 1: switch sem expressão (equivalente a "switch true")
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
}
```

**Saída:**

```bash
B
```

### 3.5 Switch com Inicialização (igual ao `if`)

```bash
package main

import "fmt"

func calcularNota() int {
    return 95
}

func main() {
    switch nota := calcularNota(); {
    case nota >= 90:
        fmt.Println("Excelente")
    case nota >= 70:
        fmt.Println("Bom")
    default:
        fmt.Println("Precisa melhorar")
    }
    // nota não existe aqui fora
}
```

**Saída:**

```bash
Excelente
```

### 3.6 Type Switch (para interfaces)

Um `type switch` permite inspecionar, em tempo de execução, qual é o tipo concreto armazenado em uma variável de interface — muito comum ao lidar com `interface{}` (ou `any`, a partir do Go 1.18).

```bash
package main

import "fmt"

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

// Com acesso ao valor já convertido para o tipo identificado
func processar(v interface{}) {
    switch valor := v.(type) {
    case int:
        fmt.Printf("Inteiro: %d\n", valor)
    case string:
        fmt.Printf("String: %s\n", valor)
    }
}

func main() {
    identificarTipo(42)
    identificarTipo("texto")
    processar(10)
    processar("olá")
}
```

**Saída:**

```bash
É um inteiro
É uma string
Inteiro: 10
String: olá
```

### 3.7 Comparações Importantes

Em Go, o `if` aceita **apenas** expressões booleanas — não existe "truthy/falsy" como em Python ou JavaScript.

```bash
package main

import "fmt"

func main() {
    x := 5

    // ✅ Correto
    if x > 0 {
        fmt.Println("x é positivo")
    }

    // ❌ Erro (diferente de Python/JS/C)
    // if x { } // erro: x é int, não bool

    // ✅ Correto
    if x != 0 {
        fmt.Println("x é diferente de zero")
    }
}
```

**Saída:**

```bash
x é positivo
x é diferente de zero
```

⚠️ **Slice/map vazio não é a mesma coisa que `nil`:**

```bash
package main

import "fmt"

func main() {
    var s []int // nil
    if s == nil {
        fmt.Println("s é nil")
    }

    s = []int{} // agora é vazio, mas NÃO é nil (foi alocado)
    if s != nil {
        fmt.Println("s não é mais nil, apenas vazio")
    }
}
```

**Saída:**

```bash
s é nil
s não é mais nil, apenas vazio
```

### 3.8 Tabela Comparativa: Go vs Outras Linguagens

| Característica | Go | C/Java | Python | JavaScript |
|---|---|---|---|---|
| `{}` obrigatório | ✅ | ✅ | ❌ | ❌ |
| Parênteses na condição | ❌ | ✅ | ❌ | ✅ |
| `else` na mesma linha | ✅ | ❌ | N/A | ❌ |
| Fallthrough automático | ❌ | ✅ | N/A | ✅ (sem `break`) |
| `if` com init statement | ✅ | ❌ | ❌ | ❌ |
| Condição precisa ser booleana | ✅ | ✅ | ❌ (truthy) | ❌ (truthy) |

[Voltar ao Sumário](#sumário)

---

## 4. Arrays, Slices, Maps, Structs e Ponteiros

### 4.1 Arrays

Arrays em Go possuem **tamanho fixo**, determinado em tempo de compilação. O tamanho faz parte da assinatura do tipo — isso significa que `[3]int` e `[5]int` são tipos **incompatíveis**, mesmo armazenando o mesmo tipo de elemento.

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
```

**Saída:**

```bash
Array de Números: [10 20 30] (Tipo: [3]int)
Array de Cores: [Vermelho Verde Azul] (Tamanho: 3)
```

💡 **Caso de uso prático:** arrays são úteis para representar blocos de tamanho fixo conhecido, como um checksum MD5 (`[16]byte`), um hash SHA-256 (`[32]byte`) ou buffers de I/O de baixo nível.

### 4.2 Slices

Slices são a estrutura de lista dinâmica mais usada em Go — uma abstração sobre um array subjacente. Internamente, um slice guarda três informações: um **ponteiro** para o array, o **comprimento** (`len`) e a **capacidade** (`cap`).

```bash
package main

import "fmt"

func main() {
    // make([]Tipo, len, cap)
    items := make([]string, 0, 3)

    items = append(items, "Servidor 01")
    items = append(items, "Servidor 02")
    items = append(items, "Servidor 03")

    fmt.Printf("Slice: %v | Len: %d | Cap: %d\n", items, len(items), cap(items))

    // Ao exceder a capacidade, Go realoca automaticamente um novo array maior
    items = append(items, "Servidor 04")
    fmt.Printf("Após Expansão: %v | Len: %d | Cap: %d\n", items, len(items), cap(items))

    // Fatiamento (slicing) — items[inicio:fim], fim é exclusivo
    subGrupo := items[1:3]
    fmt.Println("Subgrupo (1:3):", subGrupo)
}
```

**Saída:**

```bash
Slice: [Servidor 01 Servidor 02 Servidor 03] | Len: 3 | Cap: 3
Após Expansão: [Servidor 01 Servidor 02 Servidor 03 Servidor 04] | Len: 4 | Cap: 6
Subgrupo (1:3): [Servidor 02 Servidor 03]
```

⚠️ **Cuidado com realocação silenciosa:** quando a capacidade é excedida, o `append` cria um novo array e copia os dados. Isso significa que um slice obtido antes da expansão pode continuar apontando para o array antigo — uma fonte clássica de bugs sutis em código concorrente.

💡 **Caso de uso prático:** listas dinâmicas de itens em APIs, como resultados paginados vindos do banco de dados.

### 4.3 Maps

`map` é uma estrutura de dados de chave-valor, não ordenada (implementada como tabela hash).

```bash
package main

import "fmt"

func main() {
    rates := map[string]float64{
        "USD": 5.05,
        "EUR": 5.45,
    }

    rates["GBP"] = 6.30

    // Leitura segura com o idioma "comma ok"
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
```

**Saída:**

```bash
Moeda JPY não encontrada no cache!
Tabela de Cotações: map[EUR:5.45 USD:5.05]
```

⚠️ **Aviso de concorrência:** o `map` nativo do Go **não é seguro** para leitura/escrita concorrente por múltiplas goroutines. Use `sync.RWMutex` ou `sync.Map` quando houver acesso simultâneo — caso contrário o programa pode entrar em pânico em tempo de execução (`fatal error: concurrent map writes`).

### 4.4 Structs, Métodos e Ponteiros

Structs agrupam dados relacionados e podem ter métodos associados — é a forma que Go usa para simular orientação a objetos, sem herança clássica, apenas composição.

```bash
package main

import "fmt"

type Account struct {
    Owner   string  `json:"owner"`
    Balance float64 `json:"balance"`
}

// Receiver por VALOR: recebe uma cópia, não altera a instância original
func (a Account) Display() string {
    return fmt.Sprintf("Titular: %s | Saldo: R$%.2f", a.Owner, a.Balance)
}

// Receiver por PONTEIRO: modifica o estado interno da instância original
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
```

**Saída:**

```bash
Antes: Titular: Ana Silva | Saldo: R$1000.00
Após Depósito: Titular: Ana Silva | Saldo: R$1500.50
```

> 🟡 **Regra prática:** use receiver por ponteiro (`*Account`) sempre que o método precisar alterar o estado da struct, ou quando a struct for grande (evita copiar dados desnecessariamente). Use receiver por valor apenas para structs pequenas e imutáveis.

[Voltar ao Sumário](#sumário)

---

## 5. Tratamento de Erros, Defer, Panic e Interfaces

### 5.1 Tratamento Idiomático de Erros e Error Wrapping

Em Go, erros são **valores** retornados explicitamente por funções — não exceções lançadas e capturadas silenciosamente como em Java ou Python. Essa filosofia obriga o desenvolvedor a tratar cada erro no ponto exato em que ele ocorre.

```bash
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

// Unwrap permite que errors.Is/errors.As "enxerguem" o erro original encapsulado
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

        // errors.Is verifica a cadeia de erros encapsulados (wrapping)
        if errors.Is(err, ErrNotFound) {
            fmt.Println("Causa raiz: o recurso realmente não existe.")
        }
    }
}
```

**Saída:**

```bash
Erro capturado: erro na operação 'SELECT': recurso não encontrado
Causa raiz: o recurso realmente não existe.
```

💡 **Por que usar `Unwrap`?** Ele permite que `errors.Is` e `errors.As` "enxerguem através" de erros customizados encapsulados, comparando com o erro sentinela original mesmo que ele esteja embrulhado em várias camadas — um padrão muito usado em aplicações reais com múltiplas camadas (handler → service → repository).

### 5.2 Use `defer` com Sabedoria

A instrução `defer` adia a execução de uma chamada até o momento em que a função que a contém retornar — seja normalmente, seja por `panic`. É a ferramenta idiomática para liberar recursos (fechar arquivos, desbloquear mutexes, fechar conexões).

```bash
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
    defer c.mu.Unlock() // garante o destravamento ao sair da função, mesmo em caso de pânico
    c.value++
}

func main() {
    counter := SafeCounter{}
    counter.Inc()
    counter.Inc()
    counter.Inc()
    fmt.Println("Valor do Contador:", counter.value)
}
```

**Saída:**

```bash
Valor do Contador: 3
```

⚠️ **Ordem de execução:** múltiplos `defer` em uma mesma função executam em ordem **LIFO** (último a entrar, primeiro a sair) — como uma pilha:

```bash
package main

import "fmt"

func main() {
    fmt.Println("Início")
    defer fmt.Println("Primeiro defer")
    defer fmt.Println("Segundo defer")
    defer fmt.Println("Terceiro defer")
    fmt.Println("Fim da função")
}
```

**Saída:**

```bash
Início
Fim da função
Terceiro defer
Segundo defer
Primeiro defer
```

### 5.3 Panic e Recover

Go não tem exceções no sentido tradicional, mas possui `panic` e `recover` para lidar com erros verdadeiramente excepcionais (situações das quais o programa não consegue se recuperar de forma normal, como um índice de slice fora dos limites). O uso de `panic`/`recover` deve ser reservado para casos raros — o caminho idiomático continua sendo retornar `error`.

```bash
package main

import "fmt"

func divisaoSegura(a, b int) (resultado int, err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("recuperado de um panic: %v", r)
        }
    }()

    resultado = a / b // se b for 0, isso causa panic: division by zero
    return resultado, nil
}

func main() {
    res, err := divisaoSegura(10, 2)
    fmt.Println(res, err)

    res, err = divisaoSegura(10, 0)
    fmt.Println(res, err)
}
```

**Saída:**

```bash
5 <nil>
0 recuperado de um panic: runtime error: integer divide by zero
```

⚠️ `recover()` só tem efeito quando chamado **diretamente dentro de uma função `defer`**. Fora desse contexto, ele retorna `nil` e não faz nada.

### 5.4 Interfaces Desacopladas (Satisfação Implícita)

Interfaces em Go especificam um **contrato de comportamento** (quais métodos um tipo deve ter). Diferente de Java, **não é preciso declarar explicitamente** que um tipo implementa uma interface — basta que ele possua os métodos exigidos (satisfação implícita/estrutural).

```bash
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
```

**Saída:**

```bash
[EMAIL SENT]: Bem-vindo ao sistema!
[SMS SENT]: Seu código de verificação é 8492.
```

💡 **Por que isso importa?** Esse desacoplamento é o que torna código Go fácil de testar: em testes, basta criar um tipo "fake" que implemente a mesma interface (por exemplo, um `MockSender`), sem precisar de frameworks de mocking complexos.

[Voltar ao Sumário](#sumário)

---

## 6. Concorrência: Goroutines e Channels

### 6.1 O que é uma Goroutine? 📚

**🔹 Nível leigo (analogia do restaurante)**

Imagine um único funcionário (a CPU) em um restaurante. Ele precisa atender telefone ☎️, levar pratos aos clientes 🍽️ e anotar pedidos ✍️.

- **Sem goroutine:** ele faz uma coisa de cada vez. Se está atendendo o telefone, todos os clientes esperam. Lento e ineficiente.
- **Com goroutine:** é como se esse funcionário se multiplicasse, ganhando "cópias leves" de si mesmo. Uma cuida do telefone, outra leva os pratos, outra anota os pedidos — e o "gerente" (o *scheduler* do Go) decide quem faz o quê e por quanto tempo, alternando rapidamente entre elas. Parece que tudo acontece ao mesmo tempo, mas na prática é intercalado (concorrente).

> 🧠 **Resumo leigo:** goroutine é uma tarefa muito leve que roda "ao mesmo tempo" que outras, sem exigir muitos recursos.

**🔹 Nível engenharia / sistemas**

- Goroutine é uma **thread leve** gerenciada pelo *runtime* do Go, não pelo sistema operacional.
- Cada goroutine ocupa cerca de **~2KB** de stack inicial (contra 1-8MB de uma thread do SO).
- O *scheduler* do Go usa o modelo **M:N** — mapeia N goroutines para M threads reais do sistema operacional.
- A partir do Go 1.14, o scheduler usa **preempção assíncrona**, interrompendo goroutines em pontos seguros mesmo em loops de longa duração, sem depender apenas de chamadas de função.
- A troca de contexto entre goroutines é muito mais barata que entre threads do SO, pois não envolve *syscalls* e preserva menos estado.

> Fundamento técnico: o scheduler usa temporizadores e sinais do sistema operacional (como `SIGURG` no Linux) para forçar a preempção. O hardware continua executando instruções sequencialmente; a sensação de paralelismo vem da alternância extremamente rápida entre tarefas.

### 6.2 Goroutines no Back-end

No desenvolvimento de APIs, microsserviços e workers, goroutines são usadas para:

- **Requisições HTTP concorrentes:** cada `http.Request` roda em sua própria goroutine — o pacote `net/http` já faz isso por padrão, sem configuração extra.
- **I/O não bloqueante:** enquanto uma goroutine aguarda dados do banco, da rede ou do disco, o scheduler pausa ela automaticamente e executa outra. Em outras linguagens isso exigiria `async`/`await` manual.
- **Processamento assíncrono:** salvar arquivos, enviar e-mails, processar filas — tudo rodando em goroutines separadas, sem travar a resposta imediata ao cliente.
- **Padrões concorrentes:** pipelines (*fan-in/fan-out*), *worker pools*, timeouts combinados com `select` e canais.

```bash
go processarPedido(pedido) // roda em background, sem travar a resposta
```

### 6.3 Goroutines Existem Apenas no Go?

Não. O conceito de threads leves gerenciadas pelo runtime existe em várias linguagens:

| Linguagem | Nome | Notas |
|---|---|---|
| Erlang/Elixir | Processos | Isolados, comunicação via troca de mensagens |
| Kotlin | Coroutines | Leves, suportadas nativamente pela linguagem |
| Java | Virtual Threads (Project Loom) | Desde o Java 21, muito similar às goroutines |
| C++ | Coroutines (C++20) | Leves, mas sem scheduler automático embutido |
| Rust | Tokio / async-std (tasks) | Green threads via bibliotecas assíncronas |

O diferencial do Go é que goroutines vêm **embutidas na linguagem desde o início**, com canais (`chan`) nativos e um scheduler automático — sem necessidade de bibliotecas externas ou palavras-chave como `async`/`await`.

### 6.4 Concorrência vs. Paralelismo

Concorrência é a capacidade de **lidar** com várias tarefas ao mesmo tempo, mas não necessariamente executá-las simultaneamente no exato mesmo instante.

- **Concorrência:** estruturação do programa em tarefas que podem ser executadas em ordem intercalada.
- **Paralelismo:** execução simultânea real, em múltiplos núcleos/CPUs.

> Analogia: **concorrente** é você cozinhar e responder mensagens no celular, alternando entre as duas tarefas. **Paralelo** é você cozinhar enquanto outra pessoa responde as mensagens, cada um usando uma CPU diferente.

Esse conceito não é exclusivo do Go — vem da ciência da computação desde os anos 1960 (C. A. R. Hoare, com o modelo **CSP — Communicating Sequential Processes**, que inspirou diretamente os canais do Go). A célebre frase de Rob Pike resume bem a filosofia da linguagem:

> "Concorrência é sobre lidar com muitas coisas ao mesmo tempo. Paralelismo é sobre fazer muitas coisas ao mesmo tempo."

### 6.5 Worker Pools e Channels na Prática

O lema da concorrência em Go é: **"Não comunique compartilhando memória; compartilhe memória se comunicando."** Channels são o mecanismo usado para que goroutines troquem dados de forma segura, sem precisar de locks manuais na maioria dos casos.

```bash
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
        time.Sleep(50 * time.Millisecond) // processamento simulado
        results <- Result{Job: job, Square: job.Value * job.Value}
    }
}

func main() {
    const numJobs = 5
    const numWorkers = 2

    jobs := make(chan Job, numJobs)
    results := make(chan Result, numJobs)
    var wg sync.WaitGroup

    // Inicializa os workers
    for w := 1; w <= numWorkers; w++ {
        wg.Add(1)
        go worker(w, jobs, results, &wg)
    }

    // Envia os jobs
    for j := 1; j <= numJobs; j++ {
        jobs <- Job{ID: j, Value: j * 10}
    }
    close(jobs)

    // Aguarda a conclusão dos workers e fecha o canal de resultados
    go func() {
        wg.Wait()
        close(results)
    }()

    // Processa os resultados conforme chegam
    for res := range results {
        fmt.Printf("Job #%d processado! Entrada: %d | Quadrado: %d\n", res.Job.ID, res.Job.Value, res.Square)
    }
}
```

**Saída (a ordem pode variar entre execuções, já que os workers rodam concorrentemente):**

```bash
Job #1 processado! Entrada: 10 | Quadrado: 100
Job #2 processado! Entrada: 20 | Quadrado: 400
Job #3 processado! Entrada: 30 | Quadrado: 900
Job #4 processado! Entrada: 40 | Quadrado: 1600
Job #5 processado! Entrada: 50 | Quadrado: 2500
```

### 6.6 `select` com Timeout

O `select` permite que uma goroutine aguarde múltiplos canais simultaneamente, agindo assim que qualquer um deles estiver pronto — combinado com `time.After`, é a forma idiomática de implementar timeouts.

```bash
package main

import (
    "fmt"
    "time"
)

func consultarServico() <-chan string {
    resultado := make(chan string)
    go func() {
        time.Sleep(200 * time.Millisecond) // simula uma chamada lenta
        resultado <- "resposta do serviço"
    }()
    return resultado
}

func main() {
    select {
    case res := <-consultarServico():
        fmt.Println("Recebido:", res)
    case <-time.After(100 * time.Millisecond):
        fmt.Println("Timeout: o serviço demorou demais para responder")
    }
}
```

**Saída:**

```bash
Timeout: o serviço demorou demais para responder
```

> 💡 Como o serviço simulado leva 200ms para responder e o timeout é de 100ms, o `select` escolhe o caso do `time.After`, evitando que a aplicação fique bloqueada indefinidamente esperando uma resposta lenta.

**Resumo Rápido**

| Pergunta | Resposta |
|---|---|
| Goroutine é só no Go? | Não, mas é nativa e possui scheduler próprio |
| Concorrência é só no Go? | Não, é um conceito universal da computação |
| Goroutine no back-end? | Roda requisições, I/O e tarefas assíncronas |
| Leigo → goroutine? | Vários atendentes leves em um restaurante |
| Engenharia → goroutine? | Contexto leve (~2KB stack), mapeamento M:N, preempção |

[Voltar ao Sumário](#sumário)

---

## 7. JSON, Servidor HTTP e Banco de Dados

### 7.1 APIs RESTful Idiomáticas com `net/http`

O pacote padrão `net/http` já é suficiente para construir APIs REST robustas, sem depender de frameworks externos — algo bastante comum no ecossistema Go, que valoriza a biblioteca padrão.

```bash
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

func main() {
    http.HandleFunc("/users", CreateUserHandler)
    http.ListenAndServe(":8080", nil)
}
```

**Exemplo de chamada HTTP com cURL:**

```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name": "Eduardo", "email": "eduardo@example.com"}'
```

**Saída esperada (JSON):**

```bash
{
  "id": 101,
  "name": "Eduardo",
  "email": "eduardo@example.com",
  "created_at": "2026-08-09T12:57:00Z"
}
```

💡 Sobre as *struct tags* (` `json:"name"` `): elas controlam como os campos Go são serializados/desserializados em JSON. Sem elas, o encoder usaria o nome exato do campo Go (`Name`, com maiúscula), o que geralmente não é o formato desejado em uma API pública.

### 7.2 Persistência Segura com `database/sql`

Consultas devem sempre usar **prepared statements** (placeholders como `$1`, `?`) em vez de concatenar strings — isso evita **SQL Injection** de forma automática.

```bash
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
    // Timeout de contexto garante o cancelamento de queries presas
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
```

⚠️ **Nunca faça isso:**

```bash
// ❌ Vulnerável a SQL Injection — NUNCA concatene entrada do usuário diretamente
// query := "SELECT id, name FROM users WHERE id = " + userInput
```

✅ **Sempre use placeholders**, como no exemplo de `GetUserByID` acima — o driver do banco cuida do *escaping* correto automaticamente.

[Voltar ao Sumário](#sumário)

---

## 8. Testes, Generics e Ferramentas de Qualidade

### 8.1 Testes Orientados a Tabela (Table-Driven Tests)

Esse é o padrão mais idiomático para escrever testes em Go: define-se uma tabela (slice de structs) com casos de entrada e saída esperada, e um único loop executa todos eles.

```bash
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
```

**Execução:**

```bash
go test -v ./...
```

**Saída:**

```bash
=== RUN   TestAbs
=== RUN   TestAbs/Número_Positivo
=== RUN   TestAbs/Número_Negativo
=== RUN   TestAbs/Zero
--- PASS: TestAbs (0.00s)
    --- PASS: TestAbs/Número_Positivo (0.00s)
    --- PASS: TestAbs/Número_Negativo (0.00s)
    --- PASS: TestAbs/Zero (0.00s)
PASS
ok      math_test   0.002s
```

### 8.2 Generics (Go 1.18+)

Generics permitem reutilizar algoritmos para diferentes tipos, sem perder a segurança de tipos em tempo de compilação — antes do Go 1.18, isso só era possível usando `interface{}` e *type assertions* manuais, mais verboso e menos seguro.

```bash
package main

import "fmt"

// Constraint "comparable" permite que K seja usado como chave de map
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
```

**Saída:**

```bash
Chaves extraídas: [a b]
```

**Outro exemplo — função genérica de soma para múltiplos tipos numéricos:**

```bash
package main

import "fmt"

type Numero interface {
    int | int64 | float64
}

func Soma[T Numero](valores []T) T {
    var total T
    for _, v := range valores {
        total += v
    }
    return total
}

func main() {
    inteiros := []int{1, 2, 3, 4}
    floats := []float64{1.5, 2.5, 3.0}

    fmt.Println("Soma de inteiros:", Soma(inteiros))
    fmt.Println("Soma de floats:", Soma(floats))
}
```

**Saída:**

```bash
Soma de inteiros: 10
Soma de floats: 7
```

### 8.3 Fuzzing (Go 1.18+)

O *fuzzing* nativo do Go gera automaticamente entradas aleatórias e "estranhas" para tentar quebrar sua função, encontrando casos-limite que testes manuais dificilmente cobririam.

```bash
package main

import "testing"

func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func FuzzReverse(f *testing.F) {
    f.Add("Go") // seed inicial
    f.Fuzz(func(t *testing.T, s string) {
        rev := Reverse(s)
        doubleRev := Reverse(rev)
        if s != doubleRev {
            t.Errorf("Reverso do reverso deveria ser igual ao original: %q != %q", s, doubleRev)
        }
    })
}
```

**Execução:**

```bash
go test -fuzz=FuzzReverse -fuzztime=10s
```

**Saída (exemplo):**

```bash
fuzz: elapsed: 0s, gathering baseline coverage: 0/1 completed
fuzz: elapsed: 0s, gathering baseline coverage: 1/1 completed, now fuzzing with 8 workers
fuzz: elapsed: 3s, execs: 245012 (81670/sec), new interesting: 4 (total: 5)
fuzz: elapsed: 10s, execs: 812004 (78400/sec), new interesting: 6 (total: 7)
PASS
ok      example.com/reverse   10.223s
```

### 8.4 Roteiro de Qualidade e Comandos CLI

Para manter os padrões de produção em um projeto Go, é comum rodar esta sequência de validação — ideal para incluir em um pipeline de CI:

```bash
# 1. Formatação do código de acordo com o padrão oficial do Go
gofmt -w .

# 2. Análise estática contra bugs comuns (ex: uso incorreto de printf, comparações erradas)
go vet ./...

# 3. Execução da suíte de testes com o detector de race conditions ativado
go test -race -v ./...

# 4. Verificação de cobertura de código
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

[Voltar ao Sumário](#sumário)

---

## 9. Arquitetura de Projetos em Go

### 9.1 Estrutura de Diretórios Padronizada (Clean Architecture)

Uma organização de pastas comum em projetos Go de médio/grande porte, inspirada em Clean Architecture e no [golang-standards/project-layout](https://github.com/golang-standards/project-layout):

```bash
meu-projeto/
├── cmd/
│   └── api/
│       └── main.go          # Ponto de entrada do executável
├── internal/
│   ├── domain/               # Entidades puras e regras de negócio
│   ├── usecase/               # Casos de uso e orquestração
│   ├── handler/                # Controllers HTTP/gRPC
│   └── repository/              # Acesso a banco de dados e APIs externas
├── migrations/               # Arquivos SQL de migração de banco
├── go.mod
└── go.sum
```

📌 **Por que usar `internal/`?** O Go trata pastas chamadas `internal` de forma especial: pacotes dentro dela **não podem ser importados** por módulos externos ao projeto. Isso protege a lógica interna da aplicação de ser acoplada por outros times/projetos que dependam do seu módulo.

### 9.2 Boas Práticas de Camadas

| Camada | Responsabilidade | Depende de |
|---|---|---|
| `domain` | Entidades e regras de negócio puras | Nada (é o núcleo) |
| `usecase` | Orquestra regras de negócio e chama repositórios | `domain` |
| `handler` | Recebe requisições HTTP/gRPC e traduz para chamadas de `usecase` | `usecase` |
| `repository` | Implementa acesso a banco de dados/APIs externas | `domain` (via interfaces) |

Esse desenho segue o princípio de **inversão de dependência**: as camadas externas (`handler`, `repository`) dependem de interfaces definidas em `domain`/`usecase`, e não o contrário — o que facilita testes com mocks e troca de implementações (por exemplo, trocar PostgreSQL por MongoDB sem alterar as regras de negócio).

[Voltar ao Sumário](#sumário)

---

## Recursos Adicionais

- [Documentação Oficial do Go](https://go.dev/doc/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go by Example](https://gobyexample.com/)
- [Awesome Go](https://github.com/avelino/awesome-go)
- [golang-standards/project-layout](https://github.com/golang-standards/project-layout)

---

<p align="center">📘 Apostila de Go — material de estudo e consulta contínua.</p>