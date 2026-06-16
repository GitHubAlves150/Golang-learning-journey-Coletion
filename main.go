// exercicio1_frota.go
package main

import (
    "fmt"
    "sync"
    "time"
)

type Veiculo struct {
    ID    string
    Tempo int // tempo de processamento em ms
}

func processarVeiculo(v Veiculo, wg *sync.WaitGroup) {
    defer wg.Done()
    
    fmt.Printf("🚗 Processando %s (%dms)\n", v.ID, v.Tempo)
    time.Sleep(time.Duration(v.Tempo) * time.Second)
    fmt.Printf("✅ %s concluído\n", v.ID)
}

func main() {
    frota := []Veiculo{
        {"CAR-001", 10},
        {"CAR-002", 5},
        {"CAR-003", 2},
        {"CAR-004", 3},
        {"CAR-005", 1},
    }
    
    var wg sync.WaitGroup
    inicio := time.Now()
    
    for _, v := range frota {
        wg.Add(1)
        go processarVeiculo(v, &wg)
    }
    
    wg.Wait()
    
    fmt.Printf("\n🎉 Frota processada em %v\n", time.Since(inicio))
    fmt.Println("💡 Se fosse sequencial, demoraria a soma de todos os tempos!")
}