package main

import (
	"fmt"

	"github.com/GitHubAlves150/Golang-learning-journey-Coletion.git/interfaces/datasource"
)

//import "fmt"

func main() {

	datasource := datasource.NewPostgreSQL("localhost", "teste de polimorfismo", 9)

	saveNewname(&datasource, "LucasDb")
	saveNewname(&datasource, "AlvesDb")
	GetNames(&datasource)
}

func saveNewname(d datasource.Datasource, name string) {
	d.Save(name)
}

func GetNames(d datasource.Datasource) {
	fmt.Println(d.GetAll())
}
