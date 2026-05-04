package main

import "strings"

//Retorna as n palavras mais longas do texto (se houver empate, incluir todas).
func palavrasUnicas(texto string) map[string]int{

	palavras:=strings.Fields(texto)

	palavraUnica:=make(map[string]int)
	for _, conteudo := range palavras{
		palavraUnica[conteudo]=1;
	}


	return palavraUnica;

}

