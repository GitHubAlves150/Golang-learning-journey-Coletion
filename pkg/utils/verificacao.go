package utils

import (
	"fmt"
	"unicode"
)

func EhLetra(n string) bool {
	if len(n)==0{
		fmt.Println("campo vazio")
	}

	for _, char := range n {
		if unicode.IsDigit(char) {
			return false
		}
	}
	return true
}
