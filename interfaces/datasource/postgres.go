package datasource

import (
	"fmt"
)

type PostgreSQL struct{
	host string
	port int
	database string

}
var namepgDB []string;


func NewPostgreSQL(host, database string, port int) PostgreSQL{
	return PostgreSQL{
		host: host,
		port: port,
		database: database,
	}
}

func (p *PostgreSQL) Save(name string){
	namepgDB= append(namepgDB, name)
	fmt.Println("dado persistido no postgres")
}

func (p PostgreSQL) GetAll() []string{
	return namepgDB;
}