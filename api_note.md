# 📘 Explicação do Código: Servidor HTTP com JSON 

## 🎯 O que este código faz?

Este código cria um servidor HTTP que retorna uma lista de resumos (summaries) em formato JSON quando alguém acessa o endpoint /summary.  

```bash 
É a base de uma API REST - um servidor que responde com dados estruturados (JSON) em vez de páginas HTML.
```

## 🔍 Por que foi abordado assim?

**Este é o PONTO DE PARTIDA para criar APIs em Go!**
- O curso está te ensinando:

```bash
Conceito	            Por que é importante
Servidor HTTP	        Como criar uma API que outros programas podem consumir
JSON como resposta	    APIs modernas usam JSON, não HTML
Endpoint	            Como expor dados em uma URL específica
Headers	                Como informar o tipo de conteúdo (Content-Type: application/json)

```
## 📝 Análise Linha por Linha

## 1. A Struct Summary 

```go
type Summary struct {
    ID          int    `json:"id"`
    Description string `json:"description"`
}
```
```bash
Parte	                O que faz
Summary	                Define o formato dos dados que vamos enviar
json:"id"	            Quando converter para JSON, o campo ID vira "id"
json:"description"	    O campo Description vira "description" no JSON

```

## 2. A Variável summaries 

```go
var summaries []Summary

```
- É uma fatia (slice) vazia de Summary

- ⚠️ PROBLEMA: Está sempre vazia! O servidor sempre retornará [] (lista vazia)

- Em uma API real, você teria dados de um banco ou arquivo

## 3. O Handler summariesHandler 
```go
func summariesHandler(w http.ResponseWriter, _ *http.Request) {
    fmt.Println("list summary")
    
    // Define o cabeçalho: o que estamos enviando é JSON
    w.Header().Set("Content-Type", "application/json")
    
    // Converte a slice summaries para JSON e envia
    if err := json.NewEncoder(w).Encode(summaries); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
```
```bash
Linha	                                O que faz
w http.ResponseWriter	                Onde escrevemos a resposta
_ *http.Request	                        Ignoramos a requisição (não precisamos dela)
w.Header().Set(...)	                    Diz: "Estou enviando JSON"
json.NewEncoder(w).Encode(summaries)	Converte summaries para JSON e envia
http.Error(...)	                        Se der erro, envia status 500 (Erro interno)
```

## 4. A Função main 

```go
func main() {
    http.HandleFunc("/summary", summariesHandler)
    log.Println("Server is Running")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```
```go
Linha	                                            O que faz
http.HandleFunc("/summary", summariesHandler)	    Diz: "quando alguém acessar /summary, chame a função summariesHandler"
log.Println(...)	                                Imprime no console que o servidor está rodando
http.ListenAndServe(":8080", nil)	                Inicia o servidor na porta 8080
log.Fatal(...)	                                    Se o servidor falhar ao iniciar, mostra o erro e para o programa

```
## 📊 Visualização do Fluxo 

```bash
1. Servidor inicia
   └── Porta 8080 aberta
   └── Rota "/summary" registrada

2. Cliente faz requisição GET:
   ┌─────────────────────────────────────┐
   │ GET http://localhost:8080/summary   │
   └─────────────────────────────────────┘

3. Servidor recebe e processa:
   └── Chama summariesHandler()
       ├── Define Content-Type: application/json
       ├── Converte summaries (slice vazia) para JSON: []
       └── Envia resposta

4. Cliente recebe:
   ┌─────────────────────────────────────┐
   │ HTTP/1.1 200 OK                     │
   │ Content-Type: application/json      │
   │                                     │
   │ []                                  │
   └─────────────────────────────────────┘

```
## 🧪 Como Testar
**1. Execute o servidor**
```bash

go run main.go

Saída esperada:
 

Server is Running
```

**1. Faça uma requisição (outro terminal)**

**Usando curl:**
```bash

curl http://localhost:8080/summary

Resposta:
json

[]
```

**1. Veja o log no servidor**  


list summary

## ⚠️ Limitações do Código (O que o curso vai ensinar depois)
```bash
Problema	                                    O que falta aprender
summaries está sempre vazia	                    Como adicionar dados (POST)
Só responde com GET	                            Outros métodos: POST, PUT, DELETE
Dados em memória (somem quando servidor para)	Banco de dados para persistência
Não valida nada	                                Validação de dados de entrada
```

**Este código é o "esqueleto" que o curso vai preencher nas próximas aulas!**
##✅ Resumo
```bash

Conceito	                    O código demonstra
Servidor HTTP               	Como criar um servidor com http.ListenAndServe
Endpoint	                    Como registrar uma rota com HandleFunc
JSON como resposta	            Como retornar JSON com json.NewEncoder(w).Encode()
Header	                        Como definir Content-Type: application/json
Status HTTP	                    Como retornar erro 500 com http.Error
```

```python
// RECEITA DE API MÍNIMA EM GO:
// 1. Defina sua struct com tags JSON
// 2. Crie um handler que retorna JSON
// 3. Registre a rota com HandleFunc
// 4. Inicie o servidor com ListenAndServe
```

## Este é o padrão que você usará em TODAS as APIs que criar em Go! 🚀