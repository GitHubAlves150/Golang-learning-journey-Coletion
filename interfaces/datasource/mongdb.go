package datasource

import (
	"fmt"

)
var namesMongoDB []string;


type MongoDB struct{
	host string
	organization string
}


func NewMongoDB(host, organization string) MongoDB{
	return MongoDB{
		host: host,
		organization: organization,
	}
}

//Save(name string)
//GetAll()[]string

func (m *MongoDB) Save(name string){
	namesMongoDB =append(namesMongoDB, name)
	fmt.Println("Dados persisitido no mongodb")
}

func (m MongoDB) GetAll()[]string{
	return namesMongoDB
}