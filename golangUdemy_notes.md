# 🎯 Operações em arquivos  

```go
 Conceito	               O que é	                        Por que é importante  
 defer file.Close()	       Fecha arquivo automaticamente	EVITA VAZAMENTO DE MEMÓRIA  
os.Open()	               Abre arquivo para leitura	    Retorna *File e error
bufio.Scanner	           Lê linha por linha	            Eficiente para arquivos grandes
os.ReadFile()	           Lê arquivo inteiro de uma vez	Simples para arquivos pequenos 


``` 

**Verificar erros	Sempre if err != nil OBRIGATÓRIO em Go**

# 📂 Estrutura de Arquivo para Testes  
**Crie um arquivo dados.txt para testar:**

```go
linha 1: Hello World
linha 2: Go é incrível
linha 3: Aprendendo arquivos
linha 4: Última linha
```

## 1. 📖 Método 1: os.ReadFile() - Ler arquivo INTEIRO  

- Quando usar: Arquivos PEQUENOS (< 10MB) que cabem em memória 

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    // 1. Lê o arquivo inteiro de uma vez
    dados, err := os.ReadFile("dados.txt")
    if err != nil {
        fmt.Println("Erro ao ler arquivo:", err)
        return
    }
    
    // 2. Converte para string e exibe
    fmt.Println(string(dados))
}
```
Vantagens: ✅ Simples,
✅ Uma linha
✅ Fecha automaticamente  
**Desvantagens:** ❌ Carrega tudo na memória, ❌ Ruim para arquivos gigantes

---
# 2. 📖 Método 2: os.Open() + defer - Abrir e LER COM SEGURANÇA  

**Quando usar: PRÁTICA RECOMENDADA para a maioria dos casos**  

```go 

package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    // 1. Abrir o arquivo
    file, err := os.Open("dados.txt")
    if err != nil {
        fmt.Println("Erro ao abrir arquivo:", err)
        return
    }
    defer file.Close()  // ← ESSENCIAL! Fecha automaticamente no final
    
    // 2. Criar scanner para ler linha por linha
    scanner := bufio.NewScanner(file)
    
    // 3. Ler cada linha
    for scanner.Scan() {
        linha := scanner.Text()
        fmt.Println(linha)
    }
    
    // 4. Verificar se houve erro na leitura
    if err := scanner.Err(); err != nil {
        fmt.Println("Erro ao ler arquivo:", err)
    }
}

```

**Quando usar: PRÁTICA RECOMENDADA para a maioria dos casos**

---
# 🔧 BOAS PRÁTICAS Essenciais  
- 1. SEMPRE use defer file.Close()! 
- 2. SEMPRE verifique erros! 
- 3. Use defer LOGO APÓS abrir o arquivo 
- 4. Propague erros com contexto 
  
## 📊 Comparação dos Métodos   


```go

Método	            Quando usar	                            Vantagem	                     Desvantagem
os.ReadFile()	    Arquivos pequenos (<10MB)	            ✅ Mais simples	                ❌ Carrega tudo na memória
bufio.Scanner	    Arquivos grandes, linha a linha	        ✅ Memória eficiente	            ❌ Linha >64KB quebra
bufio.Reader	    Processamento em chunks	                ✅ Controle granular	            ❌ Mais código
os.OpenFile()	    Precisa de controle de flags	        ✅ Flexível	                    ❌ Mais complexo

``` 
# 🌐 Consumindo API Externa com JSON em Go - Guia Completo

## 🎯 O que este código faz

Este código **consome uma API externa** (JSONPlaceholder), recebe dados em JSON e decodifica para uma struct Go.

> **JSONPlaceholder** é uma API fake gratuita para testes: `https://jsonplaceholder.typicode.com`

---

## 📝 Código Completo (Comentado)

```go
package main

import (
    "encoding/json"  // Pacote para trabalhar com JSON
    "fmt"           // Pacote para entrada/saída (prints)
    "net/http"      // Pacote para requisições HTTP
)

/* 
   Exemplo do JSON que vamos receber da API:
   {
     "userId": 1,
     "id": 1,
     "title": "delectus aut autem",
     "completed": false
   }
*/

// Post representa a estrutura dos dados que vamos receber da API
type Post struct {
    UserId   int    `json:"userId"`   // ← Mapeia "userId" do JSON para UserId
    Id       int    `json:"id"`       // ← Mapeia "id" do JSON para Id
    Title    string `json:"title"`    // ← Mapeia "title" do JSON para Title
    Complete bool   `json:"completed"`// ← Mapeia "completed" do JSON para Complete
    // ⚠️ As tags JSON PRECISAM ser IDÊNTICAS aos campos do JSON da API!
}

func main() {
    fmt.Println("Fazendo requisição para a API json placeholder")
    
    // 1. FAZER A REQUISIÇÃO HTTP
    //    http.Get faz uma requisição GET para a URL e retorna:
    //    - resposta: contém os dados da API (body, headers, status)
    //    - err: erro se a requisição falhar (rede, timeout, etc)
    resposta, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
    if err != nil {
        fmt.Println("Erro ao fazer requisição: ", err)
        return  // Se deu erro, para o programa
    }
    
    // 2. GARANTIR QUE A CONEXÃO SERÁ FECHADA
    //    defer: executa esta linha quando a função main terminar
    //    resposta.Body.Close(): fecha a conexão (IMPORTANTE!)
    //    ⚠️ Sempre feche o Body para evitar vazamento de memória!
    defer resposta.Body.Close()
    
    // 3. CRIAR UMA INSTÂNCIA DA STRUCT (vazia)
    var post Post  // post será preenchida com os dados da API
    
    // 4. DECODIFICAR O JSON PARA A STRUCT
    //    json.NewDecoder: cria um decodificador de JSON
    //    resposta.Body: é um "leitor" que contém o JSON recebido
    //    Decode(&post): converte o JSON para a struct post
    err = json.NewDecoder(resposta.Body).Decode(&post)
    if err != nil {
        fmt.Println("Erro ao decodificar a resposta: ", err)
        return
    }
    
    // 5. EXIBIR OS DADOS
    fmt.Println(post)  // Exibe a struct completa
    
    fmt.Println("...FIM...")
}
```
## 🔍 Explicação Detalhada
**1. A Struct Post**
```bash
type Post struct {
    UserId   int    `json:"userId"`   // ← Tag JSON
    Id       int    `json:"id"`       // ← Tag JSON
    Title    string `json:"title"`    // ← Tag JSON
    Complete bool   `json:"completed"`// ← Tag JSON
}
```
```bash
Campo Go	   Tag JSON	    Tipo	     O que representa
UserId	       "userId"	    int	         ID do usuário que fez o post
Id	           "id"         int          ID único do post
Title	       "title"	    string	     Título do post
Complete	   "completed"	bool	     Status do post (feito/não feito)
``` 

##⚠️ IMPORTANTE: As tags JSON devem ser EXATAMENTE IGUAIS aos campos retornados pela API!  
**2. Fazendo a Requisição**

resposta, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

Retorno	O que é
resposta	Contém os dados da API (body, status, headers)
err	nil se sucesso, erro caso contrário
**3. Fechando a Conexão (defer)**


defer resposta.Body.Close()  

    defer = "execute esta linha quando a função terminar"  

    SEMPRE feche o Body para evitar vazamento de memória!  

    resposta.Body é um io.ReadCloser (precisa ser fechado)  

**4. Decodificando o JSON**

err = json.NewDecoder(resposta.Body).Decode(&post)

```bash
Parte	                O  que faz
json.NewDecoder()	    Cria um decodificador JSON
resposta.Body	        Fon te dos dados (o JSON recebido)
.Decode(&post)	        Converte JSON para struct
&post	                Ponteiro para a struct (necessário para modificar)
```

## 📊 Fluxo dos Dados


1. REQUISIÇÃO
- Go ── http.Get() ──→ https://jsonplaceholder.typicode.com/todos/1

1. RESPOSTA (JSON)
```bash 
   API ──→ {
       "userId": 1,
       "id": 1,
       "title": "delectus aut autem",
       "completed": false
   }
```
2. DECODIFICAÇÃO
```bash
   json.Decoder ──→ Post{
       UserId: 1,
       Id: 1,
       Title: "delectus aut autem",
       Complete: false
   }
```
3. USO
```bash
   fmt.Println(post) → {1 1 delectus aut autem false}
```
## 🧪 Como Testar
1. Execute o programa  

- go run main.go

2. Saída esperada
```bash
Fazendo requisição para a API json placeholder
{1 1 delectus aut autem false}
...FIM...
``` 
3. Teste outras URLs da API
go

// Altere a URL para testar outros endpoints:

// Primeiro post (id=1)
"https://jsonplaceholder.typicode.com/todos/1"

// Décimo post (id=10)
"https://jsonplaceholder.typicode.com/todos/10"

// Todos os posts (array)
"https://jsonplaceholder.typicode.com/todos"

## 🔄 Versão com Slice (múltiplos posts)
```go

func main() {
    // URL que retorna VÁRIOS posts
    url := "https://jsonplaceholder.typicode.com/todos"
    
    resp, err := http.Get(url)
    if err != nil {
        fmt.Println("Erro:", err)
        return
    }
    defer resp.Body.Close()
    
    // ← MUDA AQUI: slice de posts!
    var posts []Post
    
    err = json.NewDecoder(resp.Body).Decode(&posts)
    if err != nil {
        fmt.Println("Erro:", err)
        return
    }
    
    fmt.Printf("Recebidos %d posts\n", len(posts))
    fmt.Println("Primeiro post:", posts[0])
}
``` 

## ⚠️ Erros Comuns
- 1. Tag JSON diferente do campo da API
```go 
// ❌ ERRADO - "completed" vs "complete"
Complete bool `json:"complete"`  // ← faltou a letra 'd'!

// ✅ CORRETO
Complete bool `json:"completed"`
``` 
2. Esquecer de fechar o Body 

```go
// ❌ ERRADO - vazamento de memória!
resposta, _ := http.Get(url)
// Esqueceu: defer resposta.Body.Close()

// ✅ CORRETO
resposta, err := http.Get(url)
if err != nil { return }
defer resposta.Body.Close()

```
1. Esquecer o & no Decode


// ❌ ERRADO - não modifica a variável original!
json.NewDecoder(body).Decode(post)  // ← sem &, passa cópia!

// ✅ CORRETO
json.NewDecoder(body).Decode(&post)  // ← com &, passa referência

📋 Resumo do Fluxo
Etapa	Código	O que acontece
- 1. Requisição	http.Get(url)	Chama a API
- 2. Verificar erro	if err != nil	Se falhou, para
- 3. Fechar Body	defer resp.Body.Close()	Garante limpeza
- 4. Criar struct	var post Post	Prepara lugar para os dados
- 5. Decodificar	json.NewDecoder().Decode(&post)	JSON → Struct
- 6. Verificar erro	if err != nil	Se JSON inválido, para
- 7. Usar dados	fmt.Println(post)	Processa os dados 
  
## 🎯 Aprendizados deste código
```bash
Conceito	        O que você aprendeu
Requisições HTTP	Como chamar APIs externas com http.Get
Consumir JSON	    Como receber e processar JSON de APIs
Decodificação	    json.NewDecoder().Decode()
Tags JSON      	    Mapear JSON API → Struct Go
Gerenciamento       de recursos	defer resposta.Body.Close()
Tratamento          de erros	Verificar erros em requisições e decodificação
```
## ✅ Regras de Ouro

```bash
// 1. As TAGS JSON devem ser IGUAIS aos campos da API
type Post struct {
    UserId int `json:"userId"`  // ← exatamente "userId"
}

// 2. SEMPRE feche o Body
defer resposta.Body.Close()

// 3. SEMPRE use & no Decode
json.NewDecoder(body).Decode(&post)

// 4. SEMPRE verifique erros
if err != nil {
    fmt.Println("Erro:", err)
    return
}
```

## 🚀 Próximos passos

```bash
// 1. Adicione timeout na requisição
client := &http.Client{Timeout: 10 * time.Second}
resp, err := client.Get(url)

// 2. Adicione headers (User-Agent, Auth)
req, _ := http.NewRequest("GET", url, nil)
req.Header.Set("User-Agent", "MeuApp/1.0")
resp, err := client.Do(req)

// 3. Salve os dados em arquivo
jsonData, _ := json.Marshal(post)
os.WriteFile("post.json", jsonData, 0644)

Parte da jornada de aprendizado Go - Consumindo APIs com JSON
``` 

