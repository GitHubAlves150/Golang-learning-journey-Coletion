package main


import "strings"


//Retorna um map com a frequẽncia de cada palavra(case insensitive, ignorando pontuação)
func frequenciaPalavras(Texto string) map[string]int{
	
	quant:= strings.Fields(Texto);
	freq := make(map[string]int)

	for _, conteudo := range quant{
		freq[conteudo]++
	}
	return freq;

}




