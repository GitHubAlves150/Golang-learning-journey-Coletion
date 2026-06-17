# Explicação do Código Go (Markdown Formatado)

## Visão Geral

Este programa demonstra o uso do pacote `context` em Go para:
- Transportar um valor (`operador_id`) via contexto
- Criar um contexto com **timeout**
- Cancelar/expirar uma operação e verificar o motivo do cancelamento

A função `InspecionaOperacao` lê o valor do contexto, consulta o deadline e executa um trabalho "pesado" simulado com `select` entre conclusão e cancelamento.

***

## Imports

```go
import (
    "context"  // Suporte a contextos (propagar cancelamento, deadlines, valores)
    "fmt"      // Saída formatada
    "time"     // Manipulação de tempo e durações
)
```

***

## Função `InspecionaOperacao(ctxFinal context.Context)`

### 1) Recuperar valor do contexto

```go
if operador, ok := ctxFinal.Value("operador_id").(string); ok {
    fmt.Printf("\n👤 Operador responsável: %s\n", operador)
} else {
    fmt.Printf("\n👤 Nenhum operador identificado no conexto.")
}
```

- `ctxFinal.Value("operador_id")` tenta recuperar o valor associado à chave `"operador_id"`.
- A assertiva `.(string)` converte para `string`; `ok` indica sucesso.
- **Observação**: usar strings literais como chave não é recomendado — os exemplos oficiais recomendam tipos não exportados para evitar colisões de chave entre pacotes.

***

### 2) Deadline()

```go
if horarioLimite, ok := ctxFinal.Deadline(); ok {
    tempoRestante := time.Until(horarioLimite)
    fmt.Printf("\n⏱️ Horário limite: %s (Resta exatamente: %v)\n", horarioLimite.Format("15:04:10"), tempoRestante)
} else {
    fmt.Printf("\n⏱️ Este horário é externo, não tem horário limite\n")
}
```

- `horarioLimite, ok := ctxFinal.Deadline()` retorna o tempo limite (`time.Time`) e um `bool ok` que indica se existe deadline.
- Se `ok` for true, calcula `tempoRestante := time.Until(horarioLimite)` e imprime horário e tempo restante.
- Caso contrário, indica que não há horário limite.

***

### 3) Simulação de processamento pesado e tratamento de cancelamento

```go
fmt.Printf("\n⏳ Iniciando processamento pesado....\n")

select {
case <-time.After(500 * time.Millisecond):
    fmt.Printf("\n ✅ Processamento concluído com sucesso\n")
case <-ctxFinal.Done():
    fmt.Printf("\n❌ERROR: \n", ctxFinal.Err())
}
```

- Imprime mensagem de início.
- `select` com dois cases:
  - **Case 1**: `<-time.After(500 * time.Millisecond)` aguarda **500 ms** e considera processamento concluído com sucesso.
  - **Case 2**: `<-ctxFinal.Done()` espera o canal `Done()` do contexto ser fechado (cancelamento ou timeout).
    - No branch de cancelamento, o código tenta imprimir o erro, mas **há um erro de formatação**:
      ```go
      // ❌ INCORERTO (não exibe ctxFinal.Err())
      fmt.Printf("\n❌ERROR: \n", ctxFinal.Err())
      
      // ✅ CORRETO
      fmt.Printf("\n❌ERROR: %v\n", ctxFinal.Err())
      ```
    - `ctxFinal.Err()` retorna:
      - `"context deadline exceeded"` → timeout
      - `"context canceled"` → cancelamento manual

***

## Função `main()`

```go
func main() {
    ctx := context.Background()
    operador := "operador_id"
    driver := "Lucas_Dev_2026"

    // Injetamos um dado usando value
    ctxComvalor := context.WithValue(ctx, operador, driver)

    // Criamos um timeout curto de 200ms apartir do contexto que já tinha o valor
    ctxFinal, cancel := context.WithTimeout(ctxComvalor, 200*time.Millisecond)
    defer cancel()

    // Executa a inspeção
    InspecionaOperacao(ctxFinal)
}
```

### Fluxo passo a passo:

1. `ctx := context.Background()` → cria contexto base.
2. `ctxComvalor := context.WithValue(ctx, operador, driver)` → injeta o par chave-valor no contexto.
3. `ctxFinal, cancel := context.WithTimeout(ctxComvalor, 200*time.Millisecond)` → cria contexto filho com **timeout de 200 ms**.
4. `defer cancel()` → garante que o contexto será cancelado ao final da função.
5. `InspecionaOperacao(ctxFinal)` → executa a inspeção.

***

## Comportamento Esperado na Execução

| Componente | Valor |
|------------|-------|
| Operador injetado | `Lucas_Dev_2026` |
| Timeout do contexto | `200 ms` |
| Tempo do processamento | `500 ms` |

### Resultado:

```
👤 Operador responsável: Lucas_Dev_2026

⏱️ Horário limite: 17:31:45 (Resta exatamente: 199.xxx ms)

⏳ Iniciando processamento pesado....

❌ERROR: context deadline exceeded
```

**Por que?**
- O timeout é **200 ms**, mas o processamento leva **500 ms**.
- O contexto expira antes do processamento terminar → `ctxFinal.Done()` é fechado primeiro.
- `ctxFinal.Err()` retorna `"context deadline exceeded"`.

***

## Correções Recomendadas

| Linha | Problema | Correção |
|-------|----------|----------|
| `fmt.Printf("\n❌ERROR: \n", ctxFinal.Err())` | Especificador de formato faltante | `fmt.Printf("\n❌ERROR: %v\n", ctxFinal.Err())` |
| `context.WithValue(ctx, operador, driver)` | Chave é string (pode colidir) | Use tipo não exportado: `type ctxKey string; const operadorKey ctxKey = "operador_id"` |

***

## Métodos do Contexto Utilizados

| Método | Descrição | Retorna |
|--------|-----------|---------|
| `Value(key)` | Recupera dados do "crachá" do contexto | `(valor, ok)` |
| `Deadline()` | Devolve quando o contexto vai expirar | `(time.Time, ok)` |
| `Done()` | Canal fechado quando contexto é cancelado | `chan struct{}` |
| `Err()` | Só devolve algo **depois** que `Done()` fecha | `error` |

***

## Conclusão

Este código é um **exemplo educacional** de como usar `context` para:
- Passar dados entre funções
- Controlar tempo de execução com timeout
- Cancelar operações e tratar erros corretamente

A correção principal é adicionar `%v` no `fmt.Printf` para exibir o erro do contexto.