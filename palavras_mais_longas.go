package main

import (
	"sort"
	"strings"
)




func palavrasMaisLongas(texto string, n int) []string  {
	palavras:=strings.Fields(texto)//separa as palavras do texto por espaço
	
	//usar um map para obter palavras únicas (evita repetição)
	mapUnica:= make(map[string]bool)
	for _, conteudo := range palavras{
		mapUnica[conteudo] = true;
	}

	//Passa do map para um slice para pdermos ordenar
	unicas := make([]string, 0, len(mapUnica) )
	for conteudo:=range mapUnica{
		unicas = append(unicas, conteudo)
	}


	//Ordenar a lista pelo tamanho (da maior para a menor)
	sort.Slice(unicas,  func(i, j int)bool{
		return  len(unicas[i]) > len(unicas[j])
	})

	//Pegar as 5 primeiras (ou menos, se o texto for curto)

	if len(unicas) < n {
		n = len(unicas)
	}

	return  unicas[:n]	
}
