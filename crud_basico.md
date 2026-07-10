O domínio das rotas GET, POST, PUT e DELETE (CRUD) é a base de uma aplicação web. Esta sequência de branches foi estruturada para ensinar o básico de cada método, seguindo o princípio de dividir para conquistar.


##1. Instalar as dependências.

```bash
go mod init meu-projeto-chi
go get -u github.com/go-chi/chi/v5
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
``` 