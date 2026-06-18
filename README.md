```markdown
# Processamento de Pagamento com Contexto em Go

Este documento explica um programa Go que simula o processamento de pagamento em um gateway (Pargar.me), utilizando o pacote `context` para controlar tempo de execução, cancelamentos e transporte de dados.

---

## Visão Geral

O programa demonstra:

- **Transporte de dados** via `Value()` (ID do cliente)
- **Controle de tempo** com `Deadline()` e `WithTimeout()`
- **Cancelamento manual** com `WithCancel()` e `Done()`
- **Identificação do motivo** do cancelamento com `Err()` e `errors.Is()`

É um exemplo educacional de como proteger APIs de **timeouts** e **cancelamentos de usuário** em operações de rede (como chamadas bancárias).

---

## Imports

```go
import (
    "context"   // Contextos para cancelamento, deadlines e valores
    "errors"    // Para comparar erros com errors.Is()
    "fmt"       // Saída formatada
    "time"      // Durações e tempos
)
```

---

## Função `ProcessarPagamento(ctx context.Context, valorCompra float64)`

### 1. Recuperar ID do cliente com `Value()`

```go
if clienteID, ok := ctx.Value("usuario_id").(int); ok {
    fmt.Printf("\n👤 [GATEWAY] Identificando CPF/Dados do cliente ID: %d..\n", clienteID)
}
```

- `ctx.Value("usuario_id")` recupera o valor associado à chave.
- `. (int)` converte para `int`; `ok` indica sucesso.

---

### 2. Verificar deadline com `Deadline()`

```go
if limite, ok := ctx.Deadline(); ok {
    tempoRestante := time.Until(limite)
    fmt.Printf("\n⏱️ [GATEWAY] Janela de tempo segura; a requisição expira em %v\n", tempoRestante)

    if tempoRestante < 50*time.Millisecond {
        fmt.Println("⚠️ [Gateway] Tempo insuficiente. Abortando antes de cobrar!")
        return
    }
}
```

- Se tempo restante < **50 ms**, aborta antes de chamar o banco.

---

### 3. Simular chamada ao banco (goroutine + canal)

```go
chBanco := make(chan string, 1)
go func() {
    time.Sleep(300 * time.Millisecond)
    chBanco <- "PAGAMENTO_APROVADO_TOKEN_9988"
}()
```

- Goroutine simula chamada de rede: **300 ms** de delay + token de aprovação.

---

### 4. Monitorar com `select` (banco vs contexto cancelado)

```go
select {
case resposta := <-chBanco:
    fmt.Printf("✅ [Gateway] Sucesso! Cartão cobrado: R$ %.2f. Código: %s\n", valorCompra, resposta)

case <-ctx.Done():
    motivo := ctx.Err()
    fmt.Printf("\n🚨 [Gateway] OPERAÇÃO ABORTADA PELO SISTEMA!\n")

    if errors.Is(motivo, context.DeadlineExceeded) {
        fmt.Println("❌ Motivo: Timeout — servidor do banco demorou demais")
    } else if errors.Is(motivo, context.Canceled) {
        fmt.Println("❌ Motivo: Cliente cancelou ou fechou a aba")
    }

    fmt.Println("🛡️ [Gateway] Estorno garantido. Nenhuma cobrança foi feita.")
}
```

- `select` monitora:
  - **Case 1**: resposta do banco (sucesso)
  - **Case 2**: contexto cancelado/expirado
- `ctx.Err()` retorna:
  - `context.DeadlineExceeded` → timeout
  - `context.Canceled` → cancelamento manual

---

## Função `main()`

### Contexto base

```go
ctxPAI := context.Background()
ctxUsuario := context.WithValue(ctxPAI, "usuario_id", 4042)
```

---

### CASO 1: Fluxo perfeito (timeout 500 ms, banco 300 ms)

```go
fmt.Println("--- SIMULAÇÃO 1: FLUXO PERFEITO ---")
ctxSucesso, cancel1 := context.WithTimeout(ctxUsuario, 500*time.Millisecond)
ProcessarPagamento(ctxSucesso, 150.90)
cancel1()
```

- **Resultado**: pagamento aprovado (500 ms > 300 ms).

---

### CASO 2: Banco lento (timeout 100 ms, banco 300 ms)

```go
fmt.Println("--- SIMULAÇÃO 2: BANCO LENTO ---")
ctxTimeout, cancel2 := context.WithTimeout(ctxUsuario, 100*time.Millisecond)
ProcessarPagamento(ctxTimeout, 89.90)
cancel2()
```

- **Resultado**: timeout expirado (`context.DeadlineExceeded`).

---

### CASO 3: Cliente cancela manualmente

```go
fmt.Println("--- SIMULAÇÃO 3: CANCELAMENTO MANUAL ---")
ctxCancelamento, cancelManual := context.WithCancel(ctxUsuario)

go func() {
    time.Sleep(50 * time.Millisecond)
    fmt.Println("💻 [Navegador] Usuário fechou a aba do e-commerce!")
    cancelManual()
}()

ProcessarPagamento(ctxCancelamento, 450.00)
```

- **Resultado**: cancelamento manual (`context.Canceled`).

---

## Métodos do Contexto Utilizados

| Método     | Descrição                                | Retorna            |
|------------|------------------------------------------|--------------------|
| `Value()`  | Recupera dado do contexto                | `(valor, ok)`      |
| `Deadline()` | Tempo limite do contexto               | `(time.Time, ok)`  |
| `Done()`   | Canal fechado quando contexto cancela    | `chan struct{} `   |
| `Err()`    | Motivo do cancelamento (após `Done()`)   | `error`            |

---

## Boas Práticas Demonstradas

| Prática                              | Por que é importante                          |
|--------------------------------------|------------------------------------------------|
| `cancel()` após uso                  | Libera recursos do contexto (evita leaks)     |
| Check de tempo antes de chamar banco | Evita cobrar sem confirmar transação          |
| `select` com `ctx.Done()`            | Cancela operações se usuário desistir         |
| `errors.Is(err, context.DeadlineExceeded)` | Diferencia timeout de cancelamento manual |

---

## Correções Recomendadas

| Problema                                               | Correção                                      |
|--------------------------------------------------------|-----------------------------------------------|
| `fmt.Printf("... ID: %s", clienteID)` com `int`        | `fmt.Printf("... ID: %d", clienteID)`         |
| `motivo:=ctx.Err()` (ausência de espaço)               | `motivo := ctx.Err()`                         |
| Nenhum `cancel()` no CASO 3                            | Adicionar `defer cancelManual()` ou `cancel()`|

---

## Conclusão

Este código é um **exemplo completo** de como usar `context` em operações de rede sensíveis ao tempo:

- Protege contra **timeouts** longos
- Cancela operações quando usuário **desiste**
- Informa o **motivo** do cancelamento
- Garante **segurança**: "nenhuma cobrança foi feita" se operação abortada

Padrão essencial para APIs de e-commerce, gateways de pagamento e serviços que chamam bancos/serviços externos.
```