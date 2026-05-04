package main


import "strings"

//Retornar o número total de palavras no texto
func contarPalavras(texto string) int{
	//strign.fields() sepera string por espaços em branco
	palavras:=strings.Fields(texto)
    quantidade:= len(palavras);
	return quantidade;
}
