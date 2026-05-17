package main

import (
	"fmt"
)

func idade(idade int) (int, error) {
	if idade < 0 {
		return idade-10, nil
	} else if idade > 150 {
		return idade, nil

	}
	return idade+10, nil

}

func main() {

	age, err:=idade(11)
	if err != nil{
		fmt.Println("opa")
	}

	fmt.Println(age, err)
	fmt.Println("...FIM...")
}
