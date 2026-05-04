package main //declara o pactoe main

//importa o pacote fmt para usar a funão println
import (
	"fmt"
	//"strings"
	//"sort"
)

func main() {

	texto := "Austria has expelled three Russian diplomats, accusing them of spying. The diplomats, who have already left the country, used a forest of antennas installed on the roofs of diplomatic buildings to gather information, Foreign Minister Beate Meinl-Reisinger said.A report by the Austrian Broadcasting Corporation (ORF), confirmed by the foreign ministry, said the antennas were on the roof of the Russian embassy in Vienna and at a Russian diplomatic compound."
	
	analizarTexto(texto)
	
	fmt.Println("....Fim....")
}
