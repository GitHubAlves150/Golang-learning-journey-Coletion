## Autenticação com JWT em Golang
🎯 O que é JWT?

JWT (JSON Web Token) é como um "cartão de acesso" que o servidor dá para o cliente depois que ele faz login.

Pense como um crachá de identificação:

```bash
┌─────────────────────────────────────────────────────┐
│  🪪 CARTÃO DE ACESSO                                │
│  Nome: João Silva                                   │
│  ID: 123                                            │
│  Cargo: Admin                                       │
│  Válido até: 2026-07-05 14:30:00                    │
│  Assinatura: 🔒 (criptografada)                     │
└─────────────────────────────────────────────────────┘
``` 

📝 Como o JWT é Estruturalmente

```bash
┌─────────────────────────────────────────────────────────────────────┐
│  JWT = HEADER + PAYLOAD + SIGNATURE                                 │
│                                                                     │
│  HEADER: {"alg": "HS256", "typ": "JWT"}                             │
│  PAYLOAD: {"user_id": 123, "exp": 1720000000}                       │
│  SIGNATURE: HMAC-SHA256(header + "." + payload, "chave-secreta")    │
└─────────────────────────────────────────────────────────────────────┘

Exemplo completo:
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMjMsImV4cCI6MTcyMDAwMDAwMH0.abc123def456xyz789
└─────────────────┘ └────────────────────┘ └─────────────────────┘
     HEADER                PAYLOAD              SIGNATURE
```

``` bash
┌─────────────────────────────────────────────────────────────────────┐
│  1. USUÁRIO FAZ LOGIN                                               │
│     POST /login                                                     │
│     Body: {"email":"joao@email.com", "senha":"123456"}              │
└─────────────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────────┐
│  2. SERVIDOR VALIDA CREDENCIAIS                                     │
│     ✅ Email existe                                                 │
│     ✅ Senha está correta                                           │
└─────────────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────────┐
│  3. SERVIDOR GERA O JWT                                             │
│     Payload: {"user_id": 123, "email": "joao@email.com"}            │
│     Assina com uma chave secreta                                    │
│     Resultado: "eyJhbGciOiJIUzI1NiIs...xyz" (token)                 │
└─────────────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────────┐
│  4. SERVIDOR ENVIA O TOKEN PARA O CLIENTE                           │
│     Resposta: {"token": "eyJhbGciOiJIUzI1NiIs...xyz"}               │
└─────────────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────────┐
│  5. CLIENTE ARMAZENA O TOKEN                                        │
│     Ex: localStorage.setItem('token', token)                        │
└─────────────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────────┐
│  6. CLIENTE FAZ REQUISIÇÃO PROTEGIDA                                │
│     GET /api/dados-secretos                                         │
│     Header: Authorization: Bearer eyJhbGciOiJIUzI1NiIs...xyz        │
└─────────────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────────┐
│  7. SERVIDOR VALIDA O TOKEN                                         │
│     ✅ Token não expirou                                            │
│     ✅ Assinatura é válida                                          │
│     ✅ Usuário existe                                               │
└─────────────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────────────┐
│  8. SERVIDOR RETORNA OS DADOS                                       │
│     Resposta: {"dados": "informação secreta"}                       │
└─────────────────────────────────────────────────────────────────────┘

```

💻 Implementação JWT em Go


1. Instalar a biblioteca JWT
``` bash
 go get github.com/golang-jwt/jwt/v5
```