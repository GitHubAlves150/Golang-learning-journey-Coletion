# 🌐 Guia da Documentação `net/http` do Go

## Para que serve o pacote `net/http`?

O pacote `net/http` é a biblioteca padrão do Go para tudo que envolve comunicação HTTP. Ele fornece as ferramentas para construir **clientes** (quem faz requisições para APIs) e **servidores** (quem cria APIs).

Tudo o que você usou nos seus exercícios, como `http.Get`, `http.HandleFunc` e `http.ListenAndServe`, faz parte deste pacote. A documentação oficial, disponível em `https://pkg.go.dev/net/http`, é a fonte definitiva para aprender seus detalhes.

---

## 📚 Organização da Documentação

A página da documentação é dividida em grandes seções que você pode navegar usando o menu **Index** no topo da página.

### 1. Lado do Cliente (Client-Side) - *Consumindo APIs*

É a parte que permite ao seu programa Go agir como um navegador ou aplicativo, fazendo requisições para servidores externos.

*   **Funções de atalho:** `Get`, `Head`, `Post`, `PostForm`. São ideais para requisições simples e rápidas.
*   **`type Client`:** A estrutura principal para clientes mais complexos. Permite configurar:
    *   **Timeouts:** Limite de tempo para a requisição não travar o programa.
    *   **Cookies:** Para manter sessões ou enviar dados de autenticação.
    *   **Redirecionamentos:** Controlando como o cliente lida com respostas 3xx.
*   **`type Request`:** Representa a requisição que será enviada. Com `http.NewRequest`, você pode construir uma requisição detalhadamente, definindo headers (`Authorization`, `Content-Type`), método HTTP (`GET`, `POST`) e corpo (`Body`).

### 2. Lado do Servidor (Server-Side) - *Criando APIs*

É a parte que permite ao seu programa Go receber e responder a requisições, agindo como um servidor web.

*   **`func HandleFunc(pattern string, handler func(ResponseWriter, *Request))`:** Registra uma função handler para uma rota específica (ex: `/submit`).
*   **`func ListenAndServe(addr string, handler Handler) error`:** Inicia o servidor HTTP e o coloca para "escutar" requisições em uma porta (ex: `:8080`).
*   **`type Handler` e `HandlerFunc`:** Conceitos centrais. `Handler` é uma interface com o método `ServeHTTP`. `HandlerFunc` é um tipo que adapta uma função comum com a assinatura `func(w ResponseWriter, r *Request)` para se tornar um `Handler`.
*   **`type ResponseWriter`:** A interface usada para construir a resposta. É com ela que você define o status da resposta (`WriteHeader`) e escreve o corpo da resposta (`Write` ou `fmt.Fprintf`).

### 3. Constantes e Funções Auxiliares

Tornam o código mais legível e seguro, substituindo números e textos brutos.

*   **Constantes de Método HTTP:** `MethodGet`, `MethodPost`, `MethodPut`, etc.
*   **Constantes de Status HTTP:** `StatusOK` (200), `StatusNotFound` (404), `StatusInternalServerError` (500).
*   **`func Error(w ResponseWriter, error string, code int)`:** Um atalho para responder com uma mensagem de erro e um código HTTP específico.

---

## 🎯 Os 4 Pilares do `net/http` que Você Precisa Dominar

1.  **O Padrão `Handler`:** A alma do servidor.
    Qualquer função com a assinatura `func(w http.ResponseWriter, r *http.Request)` pode ser usada como um manipulador (handler) de rotas.
    ```go
    // Esta assinatura específica é tudo que o pacote http precisa.
    func minhaFuncaoHandler(w http.ResponseWriter, r *http.Request) {
        // Lógica da sua rota aqui.
    }