package main

import "strings"

// LimparPalavra - remove pontuação e converte para minusculo
func limpapalavra(palavra string) string {
	//Remove pontuação
	palavra = strings.Trim(palavra, ".,!;:()\"'")
	return strings.ToLower(palavra)

}

func Eh_palavra_Valida(palavra string) bool {
	if len(palavra) > 0 && !strings.ContainsAny(palavra, "0123456789") {
		return true
	} else {
		return false;

	}
}
