
# Funções anônimas  

**Anotações da aula de funões anônimas, closures e funções variádicas da UDEMY**
- Author: Lucas Lorenço Alves
  
**Nota:**

**Anotações rápidas sobre funçoes em golang**


```go
func soma(a, b int) int{
    
    return um_int
} // Tem dois argumentos a e b do tipo int e retorna um int

func divide(a , b int) (int, int){


    return quociente_int, dividendo_int
}

```
---

- O que é função anônima? È uma função que roda dentro de outra função   
  
```go
anon := func(a, b int)int{
	return a+b;
}


fmt.Println("...", anon(2, 3) );
  ```
  ---
  O que é clousures?    
  È um função que captura e lembra das variasveis definidas fora dela mesmo depois que a função externa terminou a esecução.  
```go
func acumulador(inicial int) func(int) int {
    total := inicial  // ← variável capturada
    return func(valor int) int {
        total += valor
        return total
    }
}

func main() {
    acc := acumulador(100)
    
    fmt.Println(acc(10))  // 110
    fmt.Println(acc(20))  // 130
    fmt.Println(acc(5))   // 135
}

```

```go
package main //declara o pactoe main

//importa o pacote fmt para usar a funão println
import (
	"fmt"
)

func oddGenerator() func() int {
	i := -1

	return func() int {
		i += 2
		return i
	}
}
func main() {

	nextOddgenerator := oddGenerator()

	fmt.Println(nextOddgenerator())
	fmt.Println(nextOddgenerator())
	

	fmt.Println("....Fim....")
}



  ```  
  
---
  Oque são funçoes variadicas??   
  Funções variádicas são funções que podem receber um número variável de argumentos do mesmo tipo. Em Go, isso é feito usando ... antes do tipo do último parâmetro.

```go

// A função soma pode receber N números inteiros
func soma(numeros ...int) int {
    total := 0
    for _, n := range numeros {
        total += n
    }
    return total
}

func main() {
    fmt.Println(soma(1, 2))           // 3
    fmt.Println(soma(1, 2, 3))        // 6
    fmt.Println(soma(1, 2, 3, 4, 5))  // 15
    fmt.Println(soma())               // 0 (pode receber zero argumentos)
}


```