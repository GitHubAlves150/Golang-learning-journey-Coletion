package main

import(
    "fmt"
    "math/rand"
    "sync"
    "time"
)
 
var votos = make(map[string]int)  // ← map compartilhado!
var mu sync.Mutex


func votar(wg *sync.WaitGroup) {
    defer wg.Done()
    mu.Lock()//Comente esta linha e clica no F5 para debugar linha por linha. Verás que o map vai ser cnsultado simultaneo
    // Escolhe um candidato aleatório
    candidatos := []string{"A", "B", "C"}
    escolhido := candidatos[rand.Intn(3)]
    
    // ← PROBLEMA! Acessando map sem proteção
    votos[escolhido]++
       mu.Unlock()//Comente esta linha e clica no F5 para debugar linha por linha. Verás que o map vai ser cnsultado simultaneo
}

func sum(s []int, c chan int) {
      sum := 0
       for _, v := range s {
               sum += v
       }
       c <- sum // send sum to c
 }
 

 func main() {
    
    var wg sync.WaitGroup
    numGoroutines := 1000
    
    inicio := time.Now()
    
    for i := 0; i < numGoroutines; i++ {
        wg.Add(1)
        go votar(&wg)
    }
    
    wg.Wait()
    
    total := 0
    for _, v := range votos {
        total += v
    }
    
    fmt.Printf("Esperado: %d votos\n", numGoroutines)
    fmt.Printf("Total apurado: %d votos\n", total)
    fmt.Printf("Distribuição: %+v\n", votos)
    fmt.Printf("Tempo: %v\n", time.Since(inicio))
    
 
}