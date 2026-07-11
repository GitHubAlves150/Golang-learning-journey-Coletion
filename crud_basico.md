![Imagem Gerada pelo Gemini](Gemini_Generated_Image_xcwu5dxcwu5dxcwu.png)

# 🚀 Roteamento Essencial: Entendendo o Método POST

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

## 🎯 Branch Atual: `Topic/API_CRUD_POST_I`

Nesta *branch*, avançamos com um `POST` simples para aprender de forma macro como funciona a gravação de requisições no banco de dados. Como ainda não possuímos uma interface de inserção no *front-end*, utilizamos o script `curl` abaixo para simular o envio de dados via terminal:

```bash
curl -X POST http://localhost:8080/usuario \
  -H "Content-Type: application/json" \
  -d '{"nome": "Carlos Silva", "email": "carlos@gmail.com", "sexo": "Masc"}'
``` 

---

## 📚 O que Aprendemos com este Código?

### 1. O Ciclo de Vida Completo dos Dados na API
Este código funciona como uma aula prática sobre a transformação de dados no ecossistema de uma API:
* **Entrada (JSON Puro):** O cliente envia os dados através de textos simples no corpo da requisição (`nome`, `email`, `sexo`).
* **Processamento (Struct Go):** O método `json.NewDecoder` captura esse fluxo de texto e o traduz para um objeto que o Go consegue manipular (`NovoUsuario`).
* **Persistência (SQL/GORM):** O comando `db.Create` converte a estrutura Go em uma instrução SQL equivalente e salva o registro fisicamente no Postgres.
* **Saída (JSON Enriquecido):** O `json.NewEncoder` converte o objeto atualizado de volta em JSON e devolve a resposta para quem chamou a API.

### 2. O Trabalho em Equipe entre Go, GORM e Postgres
Compreendemos o papel fundamental que as tags da *struct* exercem como um contrato de dados:
* As tags `gorm:"default:..."` sinalizam para o ORM que aqueles campos específicos estão intencionalmente vazios no ambiente do Go.
* Isso faz com que o GORM monte a instrução `INSERT` omitindo essas colunas da query, abrindo caminho para que as funções nativas do banco de dados (`gen_random_uuid()` e `CURRENT_TIMESTAMP`) criem e preencham os valores automaticamente.

### 3. Anatomia de um Handler HTTP Nativo
O código reforça a padronização e as regras de comunicação exigidas pela arquitetura web:
* **`w.Header().Set("Content-Type", "application/json")`**: Configura o cabeçalho para avisar o cliente (aplicativo ou navegador) que a resposta enviada é um JSON válido.
* **`w.WriteHeader(http.StatusCreated)`**: Responde explicitamente com o status HTTP **201**, que é a convenção global de sucesso para registros criados no banco de dados.

### 4. O Uso do Roteador Chi com Assinatura Padrão
O código demonstra como o **Chi** atua de forma invisível, modular e extremamente leve. Ele preserva o uso das assinaturas nativas do ecossistema Go (`http.ResponseWriter` e `*http.Request`), entregando um código limpo, de fácil manutenção e totalmente livre de amarras com frameworks pesados.
