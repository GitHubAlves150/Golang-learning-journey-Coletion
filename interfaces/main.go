package main

import (
	"log"
	"net/http"
)

// ==============================================================
// Servidor que ENVIA JSON(GET)
// ==============================================================




func main() {

	http.HandleFunc("/rota", created)

	log.Println("Acesse o servidor em http.localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}