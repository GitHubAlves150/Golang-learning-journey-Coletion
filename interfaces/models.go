package main

// Pessoa é a estrutura de dados que vamos enviar
type EstruturaAPI struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"img"`
}


// Representa um banco de dados - vai crescer conforme adicionamos
var banco []EstruturaAPI