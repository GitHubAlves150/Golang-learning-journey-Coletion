package main

import (
	"fmt"
	"sync"
	"time"
)

// Simula um banco de dados compartilhado
type BancoDados struct {
	UltimaPosicao map[string]string
	mu            sync.Mutex //Protege o map
}

func (db *BancoDados) SalvarPosicao(veiculoID, posicao string) {
	db.mu.Lock()         //Trava (só uma gouroutine por vez)
	defer db.mu.Unlock() //Destrava

	//Zona crítica (só uma goroutine)
	db.UltimaPosicao[veiculoID] = posicao
	fmt.Printf("\nSalvo: %s -> %s\n", veiculoID, posicao)
}

func (db *BancoDados) LerPosicao(veiculoID string) string {
	db.mu.Lock()
	defer db.mu.Unlock()

	return db.UltimaPosicao[veiculoID]
}

func EnviaDados(db *BancoDados, veiculoID string, wg *sync.WaitGroup) {
	defer wg.Done()

	count := 2
	for i := 0; i < count; i++ {
		posicao := fmt.Sprintf("Leitura %d: lat: -23.44  lon: -32.333", i)
		db.SalvarPosicao(veiculoID, posicao)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {

	fmt.Println("=== MUTEX: PROTEGENDO DADOS COMPARTILHADOS ===\n")
	db:= &BancoDados{
		UltimaPosicao: make(map[string]string),
	}

	var wg sync.WaitGroup

	//10 veiculos tentando salvar no mesmo banco de dados ao mesmo tempo
	for i := 0; i < 10; i++ {
		wg.Add(1)

		veiculoID := fmt.Sprintf("Car-%03d", i)
		go EnviaDados(db, veiculoID, &wg)
	}
	wg.Wait()

	fmt.Println("\nDados finais salvos\n")
	for id, pos:= range db.UltimaPosicao{
		fmt.Printf("\n  %s: %s \n", id, pos)
	}

}


























