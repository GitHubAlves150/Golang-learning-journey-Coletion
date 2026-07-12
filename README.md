![Imagem Gerada pelo Gemini](Gemini_Generated_Image_xcwu5dxcwu5dxcwu.png)

# 🚀 Roteamento Essencial: Entendendo o Método PUT

O domínio das rotas `GET`, `POST`, `PUT` e `DELETE` (CRUD) é a base de qualquer aplicação web. Esta sequência de *branches* foi estruturada para ensinar o básico de cada método, seguindo o princípio de **dividir para conquistar**.

## 1. Instalar as Dependências

Execute os comandos abaixo no terminal para inicializar o módulo e instalar os pacotes necessários:

```bash
go mod init meu-projeto-chi
go get -u github.com/go-chi/chi/v5
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
``` 

---

## 🎯 Branch Atual: `Topic/API_CRUD_PUT_I`


&nbsp;&nbsp;&nbsp;&nbsp;O objetivo do PUT é atualizar um registro existente. Como desenvolvedor pleno, você sabe que o padrão HTTP dita que o PUT deve substituir o recurso inteiro ou atualizá-lo com base em um identificador único (que no nosso caso é o UUID enviado na URL).

Para fazer isso no main.go do jeito mais básico, a lógica será:

    Capturar o id (UUID) da URL usando o Chi.

    Receber o JSON com os novos dados no corpo (Body) da requisição.

    Buscar se esse usuário realmente existe (boa prática!).

    Atualizar os campos no Postgres usando o GORM e retornar o usuário atualizado.

## para testar
    ```bash
curl -X PUT http://localhost:8080/usuarios/COLE-O-UUID-AQUI \
  -H "Content-Type: application/json" \
  -d '{"nome": "Lucas Alves Refatorado", "email": "lucas.novo@gmail.com", "sexo": "Masc"}'
    ```

    ## No banco de dados, pegue o id do usuario e cole no teste acima. Repare que quando for atualizado perceberá que a linha atualizada foi para o final da lista da tabela.

    ![alt text](image.png)