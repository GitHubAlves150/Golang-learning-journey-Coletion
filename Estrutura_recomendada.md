# Estrutura Recomendada
```go
exercicios/dia5-sistema-bancario/
│
├── go.mod                          # Gerenciador de módulo
├── go.sum                          # Checksums (auto-gerado)
├── README.md                       # Documentação do projeto
│
├── cmd/                            # Pontos de entrada
│   └── bancario/
│       └── main.go                 # Programa principal
│
├── internal/                       # Código interno (não importável)
│   ├── conta/
│   │   ├── conta.go                # Struct Conta e seus métodos
│   │   └── conta_test.go           # Testes da Conta
│   │
│   └── banco/
│       ├── banco.go                # Struct Banco e seus métodos
│       └── banco_test.go           # Testes do Banco
│
├── pkg/                            # Código reutilizável (opcional)
│   └── utils/
│       └── validacoes.go           # Funções auxiliares
│
└── .vscode/                        # Configurações do VSCode
    ├── launch.json                 # Debug config
    └── tasks.json                  # Tarefas automatizadas


```  
## Explicação de Cada Pasta

```go  

Pasta	Propósito	Quando Usar
cmd/	Executáveis principais	Cada subpasta é um programa diferente
internal/	Código privado do projeto	Não pode ser importado por outros projetos
pkg/	Código público reutilizável	Pode ser importado por outros projetos
.vscode/	Configurações do editor	Debug, tasks, snippets

``` 

# Criar pasta principal  
- mkdir -p exercicios/dia5-sistema-bancario
- cd exercicios/dia5-sistema-bancario  

# Criar estrutura de pastas  
- mkdir -p cmd/bancario
- mkdir -p internal/conta
- mkdir -p internal/banco
- mkdir -p pkg/utils
- mkdir -p .vscode

# Criar arquivos
- touch go.mod README.md
- touch cmd/bancario/main.go
- touch internal/conta/conta.go
- touch internal/banco/banco.go
- touch pkg/utils/validacoes.go
- touch .vscode/launch.json
- touch .vscode/tasks.json

# Sistema Bancário em Go

## Estrutura do Projeto
- `cmd/bancario/` - Ponto de entrada
- `internal/conta/` - Lógica de Conta
- `internal/banco/` - Lógica do Banco
- `pkg/utils/` - Funções utilitárias

## Como executar
```bash
go run ./cmd/bancario   
``` 

# Como testar
```bash
go test ./...
```


# Nota !!!

**Comece por conta.go primeiro! É a base de todo o sistema.**

📋 Ordem Recomendada (Do Menos para o Mais Complexo)
text
```go
1️⃣ conta.go           (Tipo base - não depende de ninguém)
   ↓
2️⃣ validacoes.go      (Utilitários - opcional, pode ir depois)
   ↓
3️⃣ banco.go           (Depende de conta.go)
   ↓
4️⃣ main.go            (Depende de conta.go e banco.go)
   ↓
5️⃣ *_test.go          (Testes - podem vir depois ou junto)

``` 
# 🎯 Por que Começar por conta.go?  
```go

Arquivo	Dependências	Por que ordem?
conta.go	Nenhuma (só pacotes padrão)	✅ BASE - Pode ser escrito sozinho
validacoes.go	Nenhuma	✅ Pode ser escrito a qualquer momento
banco.go	conta.go	⚠️ PRECISA de conta.go pronto
main.go	conta.go + banco.go	🚨 PRECISA de tudo pronto

```