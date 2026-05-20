package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

//==============================================================
//Servidor que ENVIA JSON(GET)
//==============================================================

// Pessoa é a estrutura de dados que vamos enviar


func main() {

	//Cria um Handler(manipulador) que retorna essa estrutura pessoa em formato JSON
	http.HandleFunc("/receber", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintln(w, "Horário é ", time.Now() )

	})

	fmt.Println("Servidor rodando")
	fmt.Println("Acesse: http:/localhost:8080/people")

	//Abre servidor
	http.ListenAndServe(":8080", nil)

}
