## Context.Context em Golang
Para destravar soobre contexto em golang eu pesquisei na I.A Cloud sobre esse assunto. Entçao ela gerou um exemplo muito bom para entender de forma bem leigo sobre essa abstração de context.

🎨 A Analogia do Comandante e do Soldado

Imagine que a função ```main``` é um **comandante** e sua Goroutine é um **soldado** enviado para uma missão na floresta(Buscar dados em uma APU).

- O Comandante entrega um **Rádio** Ao soldado antes de ele ir. **Esse rádio ṕe o contect.Context**.
- O soldado entra na floresta e começa a trabalhar.
- Enquanto trabalhada , o soldado fica com o ouvido colado no rádio.No gfou, isso ṕe o ```case <- ctx.Done():```.
- Se o comantende olhar oara o relógio e ver que o tempo acabou(Timeout), ou se ele mudar de idéia e apertar um botão de emergência(Cancel(), o rádio do soldado apita vai **apitar**.
- Ao ouvir o apito(<-ctx.Done()), o soldado para o que está fazendo imediatamente, arruma as coisas eu volta pra casa.

Se vocẽ passar o **context** (o rádio), o saldado vai ficar perdido na floresta trabalhando para sempre, mesmo que o Comandante já tenha desistido da missão. Isso no Go se chama **vazamento de goroutine**(goroutine leak) e consome a memória do servidor até ele cair.

lembrando que **ctx.Done()** é um canal do **context.Context**. A grande sacada para o destravamento é o **select** que é um **switch** para channel, ele que detecta comportamento vindo de canais, é como se ele fosse o chupa cabra de barramento.

## Vamos fazer mais um exemplo partindo da histŕia do soldado.

Imagine o seguinte cenário: O comandante (main) envia o **Soldado** (uma Gouroutine) para uma missão que demora **3 segundos** para ser concluída (procurar suplementos na floresta). No entanto, o comandante avisa pelo rádio (context) que a báse corre perigo e eles tem **2 segundos** antes de abortar tudo.

```bash
package main

import (
	"context"
	"fmt"
	"time"
)

// O SOLDADO (nossa goroutine)
// Repare que o Context é sempre o primeiro parâmetro da função
func missãoDoSoldado(ctx context.Context) {
	fmt.Println("Soldado: Entrei na floresta e comecei a procurar por suprimentos")

	// Críamos um canal falso para simular o tempo que leva para achar os suprimentos
	suprimentosAchados := time.After(3 * time.Second)

	// O soldado fica ouvindo o rádio através do SELECT, enquanto o trabalho.
	select {
	case <-suprimentosAchados:
		// Se os 3 segundos passarem antes do rádio apitar, a missão foi um sucesso
		fmt.Printf("\nMissão com sucesso\n")
	case <-ctx.Done():
		// O rádio apitou ! o canal do contexto fechou porque o tempo acabou lá na base
		fmt.Println("Soldado! Recebi ordens pelo rádio, para abortar a missão e voltar pra base")
		fmt.Printf("Soldado (Motivo gravado no relatório): %v\n", ctx.Err())
	}
}

// O COMANDANTE (A função Main)
func main() {

	// O comandante pega um contexto base "vazio"
	contextBase := context.Background()

	// O Comandante define o limite : "Só temos 2 segundos"
	// Ele ganha o "contextoComComPrazo" (o rádio) e a função abortaMissão (o botão de emergência)
	contextoComPrazo, abortaMissao:=context.WithTimeout(contextBase, 5 * time.Second)

	// O defer garante que o botão de cancelar seja limpo da memória assim que a main acabar
	defer abortaMissao()

	// O Comandante envia o Soldado para a missão e entrega o rádio para ele
	go missãoDoSoldado(contextoComPrazo)


	// A Main (comandante precisa esperar um pouco na base para ver o soldado responder)
	time.Sleep(4 * time.Second)
	fmt.Println("Comandante: Operação encerrada")
}
``` 
## 🔬 O Teste da Vitória (Modificando o tempo)

Para você fixar e ver o outro lado da moeda funcionando, faça uma alteração no código:

- Vá na linha 38 (context.WithTimeout) e mude o tempo de 2*time.Second para 4*time.Second.

- Salve o arquivo e rode novamente.

O que vai acontecer agora? Como o Comandante deu 4 segundos de prazo, o Soldado vai conseguir terminar a missão dele (que leva 3 segundos) com sucesso antes do rádio apitar! O terminal vai mostrar: "Sucesso! Achei os suprimentos".

