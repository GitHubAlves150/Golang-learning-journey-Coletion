
🎯 Exercício 1 — Nível 1: Goroutine básica

Contexto: você tem um sistema de rastreamento GPS. Quando um veículo é detectado, o sistema precisa registrar a detecção em segundo plano sem travar o fluxo principal.

Sua tarefa:

Crie uma função chamada registrarVeiculo que recebe uma string com o ID do veículo e imprime: "Veículo CAR-001 registrado"
No main(), chame essa função como goroutine para três veículos: CAR-001, CAR-002, CAR-003
Coloque um time.Sleep(time.Second) no final do main() apenas para dar tempo das goroutines executarem — vamos ver por que isso é problemático no próximo exercício

O que observar na saída: a ordem dos veículos impressos vai ser sempre a mesma?

```bash
// Estudos de Goroutine
// Lucas Lorenço Alves
// Ano 09/2026

/*
	Recaptular os estudos de Goroutine.
	[*]Nível 1 : Gouroutine básica
	[*]Nível 2 : Waitgroup (esperar gorourine)
    [*]Nível 3 : Channel unbuffered
	[*]Nível 4 : Channel buffered
	[ ]Nível 5 : Armadilha do loop
	[ ]Nível 6 : Mutex (proteger estado)
	[ ]Nível 7 : Workerpool
*/

package  main


import (
	"fmt"
	"time"
)

func RegistrarVeiculo(idVeiculo string) {
	fmt.Printf("Veiculo %s registrado\n", idVeiculo)
}


func main() {

	go RegistrarVeiculo("CAR-009")
	go RegistrarVeiculo("CAR-004")
	go RegistrarVeiculo("CAR-003")

	time.Sleep(time.Second)


}


```




🎯 Exercício 1.2 — Goroutine com retorno de dados

No exercício anterior a goroutine só imprimia. Mas e se ela precisasse processar algo?

Contexto: seu sistema GPS recebe a velocidade de um veículo e precisa verificar em segundo plano se ele está acima do limite.

Sua tarefa:

Crie uma função verificarVelocidade que recebe idVeiculo string e velocidade float64
Dentro dela, verifique:
Se velocidade > 80 → imprime "⚠️ CAR-001 acima do limite: 95.5 km/h"
Se velocidade <= 80 → imprime "✅ CAR-001 dentro do limite: 60.0 km/h"
No main(), dispare como goroutine para 4 veículos com velocidades diferentes:
CAR-001 → 95.5
CAR-002 → 60.0
CAR-003 → 110.2
CAR-004 → 75.0
Mantenha o time.Sleep por enquanto — vamos eliminar ele no WaitGroup

O que observar: a ordem dos resultados vai ser sempre a mesma?


```bash

// Estudos de Goroutine
// Lucas Lorenço Alves
// Ano 09/2026

/*
	Recaptular os estudos de Goroutine.
	[*]Nível 1 : Gouroutine básica
	[ ]Nível 2 : Waitgroup (esperar gorourine)
    [ ]Nível 3 : Channel unbuffered
	[ ]Nível 4 : Channel buffered
	[ ]Nível 5 : Armadilha do loop
	[ ]Nível 6 : Mutex (proteger estado)
	[ ]Nível 7 : Workerpool
*/

package  main


import (
	"fmt"
	"time"
)

func verificarVelocidade(idVeiculo string, velocidade float64){
	if velocidade > 80.0{
		fmt.Printf("%s | %f acima do limite : 95 Km/h\n", idVeiculo, velocidade)
	}
	if velocidade <=80.0 {
		fmt.Printf("%s | %f dentro do limite: 60.0 Km/h\n", idVeiculo, velocidade)
	}

}


func main() {

	go verificarVelocidade("CAR-001", 90)
	go verificarVelocidade("CAR-009", 50)
	go verificarVelocidade("CAR-005", 77)

	time.Sleep(time.Second)


}

``` 

🎯 Exercício 1.3 — Goroutine com slice de veículos

Agora vamos deixar mais real. Em vez de chamar goroutine uma por uma, vamos iterar sobre uma lista.

Sua tarefa:

Crie uma struct Veiculo com campos ID string e Velocidade float64
No main() crie um slice com 5 veículos e suas velocidades
Use um for range para disparar uma goroutine verificarVelocidade para cada veículo
Mantenha o time.Sleep no final
go
// Dica da struct
type Veiculo struct {
    ID        string
    Velocidade float64
}

// Dica do slice
veiculos := []Veiculo{
    {"CAR-001", 95.5},
    {"CAR-002", 60.0},
    ...
}

O que observar: com 5 goroutines disparadas num loop, a ordem vai ser ainda mais imprevisível!

```bash
Operação          Sintaxe
──────────────    ──────────────────────────────
criar             []Tipo{} ou make([]Tipo, 0)
adicionar         s = append(s, elemento)
acessar           s[i]
modificar         s[i] = novoValor
tamanho           len(s)
fatiar            s[inicio:fim]
remover índice i  append(s[:i], s[i+1:]...)
copiar            copy(destino, origem)
iterar            for _, v := range s
ordenar           sort.Slice(s, func...)

```

🎯 Exercício 2 — Nível 2: WaitGroup

O problema do exercício 1 é que time.sleep é frágil e o **WaitGroup** resolve isso de forma correta - A main() só termina quando todas as goroutine terminarem e isso é independente do tempo.

- Remova o **timesleep**
- Importe o **sync**
- Crie um **sync.WaitGroup**
- Use **wg.Add(n)** antes das goroutine
- Passe **wg** para dentro da função e use **defer wg,Done()**  dentro dela
- Use **wg.Wait()** dentro da função main() para esperar que todas as goroutine sejam finalizadas, para só depois a main() terminar


```bash

import (
	"fmt"
	"sync"
)

type Veiculo struct {
	ID         string
	Velocidade float64
}

func verificarVelocidade(v Veiculo, wg *sync.WaitGroup) {
	defer wg.Done()

	if v.Velocidade > 80 {
		fmt.Printf("%s | %.1f ; Acima do permitido\n", v.ID, v.Velocidade)
	} else {
		fmt.Printf("%s | %.1f ; Dentro do permitido\n", v.ID, v.Velocidade)
	}
}
func main() {

	var wg sync.WaitGroup

	vehicle := []Veiculo{
		{"CAR-001", 22.3},
		{"CAR-002", 34.5},
		{"CAR-009", 210.1},
	}
	wg.Add(len(vehicle))
	// Criar um slice com 3 veiculos
	for _, valores:= range(vehicle)  {		

		go verificarVelocidade(valores, &wg)
	}

	// time.Sleep(time.Second) removido por causa do wait.group
	wg.Wait()

}

```

🎯 Exercício 3 — Channel unbuffered

Agora o problema muda, Até agora as gouroutines só **imprimiam**. Mas e quando você precisa
**receber o resultado** de volta.
Não dá para fazer isso:
```bash
// goroutine não aceita isso
go func(){
	return 43 // Isso não exisite em Go
}()

```
O waitgroup só executa e não retorna nada, para retornar algo deve se usar um canal. O lance do canal (channel) é que já tem um waitgroup embutido nele mas não quer dizer que não há casos que não precise usar os dois juntos.

Vamos ver o primeiro caso específico de channel que é quando ele retorna algo, ou seja, alguém escreve algo no canal (barramento ou protocolo)
```bash

import (
	"fmt"
)

type Veiculo struct {
	ID         int
	Velocidade float64
}

func calcularDistancia(idVeiculo int, km float64, ch chan string) {
	pedagio := km * 0.10
	ch <- fmt.Sprintf("CAR-00 %d | %.1f km | pedágio: R$ %.2f", idVeiculo, km, pedagio)
}

func main() {

	// var wg sync.WaitGroup Não precisa
	ch := make(chan string)

	veiculo := []Veiculo{
		{1, 330.5},
		{24, 450.9},
		{3, 90.0},
	}

	// wg.Add(len(ch))  Não precisa
	for _, valores := range veiculo {
		go calcularDistancia(valores.ID, valores.Velocidade, ch)
	}

	for i := 0; i < len(veiculo); i++ {
		resultado := <-ch
		fmt.Println(resultado)
	}

	// wg.Wait() Não precisa

	// WaitGroup resolve:    "esperar goroutines terminarem"
	// Channel resolve:      "receber DADOS das goroutines"

	// Quando você usa channel para coletar resultados
	// o próprio recebimento já sincroniza — WaitGroup seria redundante
}

```
Vamos fazer um resumo breve sobre o que estudamos no exercício 3 (channel) que resolve o problema que o waitGroup não reolve que é o cenário quando é preciso enviar dados de uma goroutine. veja um resumo prático abaixo.
- WaitGroup -> "me avisa quando terminar"
- Channel   -> "me manda o resultado quando a goroutine terminar"

Se você recebe resultado pelo **channel** o aviso de término já está embutido no recebimento e o WaitGroup seria trabalho redundante.  
Uma pergunta que me fiz enquanto estava fazendo este exercício foi **Quando escolho usar um ou outro?** E a resposta é **depensdo do designer**
- A gouroutine precisa me devolver algum dado? Se sim, então uso channel, se não, uso WaitGroup

Normalmente usamos WaitGroup quando precisamos calcular algo, persisitir algo no banco, enviar e-mail, imprimir logs ou processar arquivos sem retornos.
Um detalhe importante a saber é, quando, se tem muitos produtores e precisa calcualr e retornar valores ai sim faz sentido ter um sistema de redundância para garantir a sincronia das goroutines e, para isso, usa-se channel + WaitGroup

Outra situação é quando precisa de um **timeout e cancelamento** ppara isso precisa de **channel + context**

```bash
// Exemplo simples para relembrar

// Buscar posição GPS com timeout de 3 segundos
ch := make(chan Posicao)
go buscarPosicao(ch)

select {
case posicao := <-ch:
    fmt.Println("posição recebida:", posicao)
case <-time.After(3 * time.Second):
    fmt.Println("timeout — dispositivo não respondeu")
}

```

Consulte sempre esta tabela quando tiver dúvida.
 **Situação												 Use**
- goroutine que só executa(salvar, enviar etc..)         WaitGroup
- goroutine que retorna um valor                         channel
- goroutine que retorna um valor + muitos workers        WaitGroup + Channel
- Precisa de timeout                                     Context + Channel
- Precisa cancelar goroutines                            Context + Channel
- Proteger variável compartilhada						 Mutex

Você vai usar o resultado da goroutine?
│
├── NÃO → WaitGroup
│         "espera terminar"
│
└── SIM → Channel
          "recebe o resultado"
              │
              └── são muitos workers enviando?
                  │
                  ├── NÃO → Channel simples
                  └── SIM → WaitGroup + Channel + close()

🎯 Exercício 3.2 — Channel unbuffered com nova temática

Um sistema de e-commerce precisa calcular o valor final de cada pedido em paralelo e coletar os resultados

```bash

type Pedido struct{
	Id string
	Produto string
	Quantidade int
	Precounitario float64
}

func CalcularTotal(p Pedido, ch chan string){
	Total := float64(p.Quantidade) * p.Precounitario
	ch <- fmt.Sprintf("PEDIDO-00%s | %s    | Total: %.1f R$ ",p.Id, p.Produto, Total)
}

func main() {

	ch := make(chan string)
	pedidos :=[]Pedido{
		{"1", "Teclado", 2, 22.4},
		{"2", "Mouse", 1, 12.4},
		{"3", "Roda", 4, 2.12},
		{"4", "Carregador", 3, 104.56},
	}

	for _, valores := range(pedidos){
		go CalcularTotal(valores, ch)
	}

	for i := 0; i < len(pedidos); i++ {
		resultado := <-ch
		fmt.Println(resultado)
	}
}

```


Exercicio  3.3: 

Este programa em Go simula um sistema de monitoramento e auditoria de frota de veículos em tempo real. O objetivo principal é receber métricas individuais de cada veículo (como velocidade atual, quilometragem percorrida e consumo de combustível) e calcular o relatório de status e custos operacionais de forma concorrente, acelerando o processamento de grandes frotas.
O Papel do channel + sync.WaitGroup nesta Solução

Neste cenário, a combinação de Channels e WaitGroup implementa o padrão clássico de Sincronização e Coleta de Resultados Concorrentes:

    sync.WaitGroup (Gerenciador de Ciclo de Vida):

        Problema que resolve: A função main precisa saber quando todas as goroutines terminaram de processar os veículos para encerrar o canal com segurança.

        Na prática: wg.Add(len(veiculos)) sinaliza a quantidade de tarefas em execução. Cada goroutine chama defer wg.Done() ao finalizar. Uma goroutine separada executa wg.Wait() e faz o close(ch) para avisar que nenhum novo relatório será enviado.

    chan Relatorio (Tubo de Comunicação Seguro):

        Problema que resolve: Evita a necessidade de travar a memória (como usar sync.Mutex para alterar um slice compartilhado), seguindo o lema do Go: "Não se comunique compartilhando memória; compartilhe memória se comunicando".

        Na prática: O canal transmite os relatórios calculados pelas goroutines trabalhadoras direto para a main, que consome os dados de forma thread-safe através do loop for r := range ch.

Visão Geral da Arquitetura de Concorrência

[Veículo 1] ──(go)──> ProcessarVeiculo() ──┐
[Veículo 2] ──(go)──> ProcessarVeiculo() ──┼──> [ Canal ch ] ──> [ Loop na main ] ──> ImprimeRelatorio()
[Veículo 3] ──(go)──> ProcessarVeiculo() ──┤
[Veículo 4] ──(go)──> ProcessarVeiculo() ──┘
                                             ▲
                                    wg.Wait() fecha 'ch'

Essa abordagem garante alta performance, prevenindo data races (corrida de dados) e garantindo que o relatório final só seja impresso após o processamento completo de todos os veículos da frota.


```bash

// Estudos de Goroutine
// Lucas Lorenço Alves
// Ano 09/2026

/*
	Recaptular os estudos de Goroutine.
	[*]Nível 1 : Gouroutine básica
	[*]Nível 2 : Waitgroup (esperar gorourine)
    [*]Nível 3 : Channel unbuffered
	[ ]Nível 4 : Channel buffered
	[ ]Nível 5 : Armadilha do loop
	[ ]Nível 6 : Mutex (proteger estado)
	[ ]Nível 7 : Workerpool
*/

// Criar uma função que calcula o consumo de combustível
// Criar uma função que calcula a velocidade do veiculo

// Motrar na tela o valor do combustível
// Mostrar o toal de veículo
// Mostrar o número de veículo que esta dentro do limite
// Mostrar o número de veículo que está acima do limite

package main

import (
	"fmt"
	"sync"
)

type Veiculo struct {
	ID          string
	Velocidade  float64
	Kmrodados   float64
	LitrosporKm float64
}

type Relatorio struct {
	ID               string
	StatusVeiculo    string
	CombustivelUsado float64
	CustoTotal       float64
}

func VerificaVelocidade(v Veiculo) string {
	if v.Velocidade > 60 {
		return "ACIMA DA VELOCIDADE"
	} else {
		return "DENTRO DO LIMITE"
	}

}

func CalcularCombustivel(v Veiculo) (float64, float64) {
	combustivel := (v.Kmrodados * v.LitrosporKm)
	custo := combustivel * 6.50
	return combustivel, custo
}

func processarVeiculo(v Veiculo, ch chan<- Relatorio, wg *sync.WaitGroup) {
	defer wg.Done()

	status := VerificaVelocidade(v)
	combustivel, custo := CalcularCombustivel(v)

	ch <- Relatorio{
		ID:               v.ID,
		StatusVeiculo:    status,
		CombustivelUsado: combustivel,
		CustoTotal:       custo,
	}
}

func ImprimeRelatorio(relatorio []Relatorio) {
	acima := 0
	dentro := 0

	fmt.Println("═══════════════════════════════════════")
	fmt.Println("   RELATÓRIO DE MONITORAMENTO DE FROTA")
	fmt.Println("═══════════════════════════════════════")

	for _, valores := range relatorio {
		fmt.Printf("ID: %s | Status: %s | Combustivel: %.1f | CustoTotal: %.1f\n", valores.ID, valores.StatusVeiculo, valores.CombustivelUsado, valores.CustoTotal)

		if valores.StatusVeiculo == "ACIMA DA VELOCIDADE" {
			acima++
		} else {
			dentro++
		}
	}
	fmt.Println("═══════════════════════════════════════")
	fmt.Printf("Total de veículos : %d\n", len(relatorio))
	fmt.Printf("Acima do limite   : %d\n", acima)
	fmt.Printf("Dentro do limite  : %d\n", dentro)
	fmt.Println("═══════════════════════════════════════")

}

func main() {
	veiculos := []Veiculo{
		{"CAR-00", 100.0, 34.00, 15},
		{"CAR-01", 30.10, 240, 20},
		{"CAR-02", 90.20, 140, 22},
		{"CAR-03", 40.0, 320, 19},
		{"CAR-04", 60.0, 440, 12},
		{"CAR-05", 60.0, 140, 10},
	}

	ch := make(chan Relatorio)

	var wg sync.WaitGroup

	wg.Add(len(veiculos))

	for _, valores := range veiculos {
		go processarVeiculo(valores, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	var relatorios []Relatorio

	for v := range ch {
		relatorios = append(relatorios, v)
	}

	ImprimeRelatorio(relatorios)

}

```

🎯 Exercício 4.1 — Channel buffered

Contexto: sua central GPS recebe alertas de velocidade de vários dispositivos. Os alertas chegam rápido e precisam ser enfileirados para processamento.

Sua tarefa:

- Crie um channel buffered com capacidade para 4 mensagens
- Crie uma função gerarAlerta que recebe idVeiculo string e velocidade float64 e envia uma string pelo channel
- Lance 4 goroutines — uma por veículo — sem WaitGroup
- Receba e imprima os 4 alertas no main com um loop simples.

Dica! 
```bash

// Dica — channel buffered
ch := make(chan string, 4) // capacidade 4

// Dica — como receber N vezes
for i := 0; i < 4; i++ {
    fmt.Println(<-ch)
}

```
**O que observar:**
 com buffer, as goroutines não bloqueiam esperando o main receber — elas jogam na fila e terminam. Por isso não precisa de WaitGroup aqui.

 ```bash
package main

import "fmt"

type GerenciaVeiculo struct {
	ID         string
	Velocidade float64
}

func GeraAlerta(manager GerenciaVeiculo, ch chan<- GerenciaVeiculo) {

	if manager.Velocidade > 60 {

		ch <- GerenciaVeiculo{
			ID:         manager.ID + " Acima da velocidade",
			Velocidade: manager.Velocidade,
		}

	} else {

		ch <- GerenciaVeiculo{
			ID:         manager.ID + " No limite",
			Velocidade: manager.Velocidade,
		}

	}

}

func main() {

	veiculo := []GerenciaVeiculo{
		{"CAR-001", 23.9},
		{"CAR-002", 63.9},
		{"CAR-003", 53.9},
		{"CAR-004", 73.1},
		{"CAR-005", 32.1},
		{"CAR-006", 173.1},
	}

	ch := make(chan GerenciaVeiculo, len(veiculo))

	for _, valores := range veiculo {

		go GeraAlerta(valores, ch)
	}

	for i := 0; i < len(veiculo); i++ {
		fmt.Println(<-ch)
	}

}
 ```

// Estudos de Goroutine — Nível 6: Mutex
// Lucas Lorenço Alves
// Ano 09/2026

/*
═══════════════════════════════════════════════════
   O QUE É MUTEX E QUANDO USAR
═══════════════════════════════════════════════════

PROBLEMA:
    Quando duas ou mais goroutines acessam e modificam
    a mesma variável ao mesmo tempo, ocorre uma
    "race condition" — o resultado fica imprevisível.

    Exemplo do problema:
        goroutine 1 lê contador = 0
        goroutine 2 lê contador = 0  ← leu antes da 1 escrever!
        goroutine 1 escreve 0+1 = 1
        goroutine 2 escreve 0+1 = 1  ← sobrescreveu!
        resultado: 1 em vez de 2 ❌

SOLUÇÃO — sync.Mutex:
    Mutex (Mutual Exclusion) garante que apenas UMA
    goroutine acessa o dado por vez.

    mu.Lock()   → trava o acesso — só eu uso agora
    *contador++ → modifica com segurança
    mu.Unlock() → libera — próxima goroutine pode entrar

REGRA DE OURO:
    Sempre que múltiplas goroutines modificam
    a mesma variável → use Mutex.

QUANDO USAR CADA UM:
    Mutex   → proteger variável compartilhada (contador, mapa)
    Channel → comunicar dados entre goroutines
    WaitGroup → esperar goroutines terminarem

COMO PASSAR O MUTEX PARA A FUNÇÃO:
    Sempre por ponteiro — nunca por cópia!
    ✅ func minhaFunc(mu *sync.Mutex)
    ❌ func minhaFunc(mu sync.Mutex)  → copia o mutex, não funciona!

ESTRUTURA DESTE EXERCÍCIO:
    1. var mu sync.Mutex          → declara o mutex
    2. var wg sync.WaitGroup      → declara o waitgroup
    3. wg.Add(N)                  → avisa quantas goroutines
    4. go func(..., &mu, &wg)     → lança passando ponteiros
    5. defer wg.Done()            → dentro da goroutine
    6. mu.Lock() / mu.Unlock()    → protege o dado compartilhado
    7. wg.Wait()                  → espera todas terminarem
    8. fmt.Printf(resultado)      → imprime APÓS o Wait

ARMADILHA COMUM:
    ❌ imprimir o resultado ANTES do wg.Wait()
       o valor ainda é zero pois as goroutines
       não terminaram

    ✅ sempre: wg.Wait() → depois → fmt.Printf()

DETECÇÃO DE RACE CONDITION:
    Go tem uma ferramenta nativa para detectar:
    $ go run -race main.go
    Se houver race condition, Go avisa no terminal.
═══════════════════════════════════════════════════
*/

```bash


func registrarInfracao(id string, mu *sync.Mutex, contador *int, wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	*contador++ // * para acessar o valor do ponteiro
	mu.Unlock()
	fmt.Printf("%s registrou infração\n", id)
}

func main() {

	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(8)
	count := 0
	go registrarInfracao("Car-00", &mu, &count, &wg)
	go registrarInfracao("Car-01", &mu, &count, &wg)
	go registrarInfracao("Car-02", &mu, &count, &wg)
	go registrarInfracao("Car-03", &mu, &count, &wg)
	go registrarInfracao("Car-04", &mu, &count, &wg)
	go registrarInfracao("Car-05", &mu, &count, &wg)
	go registrarInfracao("Car-06", &mu, &count, &wg)
	go registrarInfracao("Car-07", &mu, &count, &wg)
	wg.Wait()
	fmt.Printf("Total de infrações: %d\n", count)

}

```

