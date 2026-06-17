Claro — esse exemplo não é sobre “cancelar uma goroutine” diretamente, e sim sobre **como uma goroutine fica ouvindo um sinal externo para parar**. O `context` aqui funciona como um canal de controle de vida útil da operação, algo que o pacote Go usa para cancelamento, prazos e dados de requisição. [aprendagolang.com](https://aprendagolang.com.br/o-que-e-e-como-utilizar-o-package-context/)

## O que o código faz

- `context.Background()` cria um contexto base, vazio, sem cancelamento associado. [pkg.go](https://pkg.go.dev/context)
- `context.WithCancel(...)` cria um contexto filho e uma função `cancel()` que dispara o cancelamento desse contexto. [aprendagolang.com](https://aprendagolang.com.br/o-que-e-e-como-utilizar-o-package-context/)
- A goroutine entra em `select` esperando `ctx.Done()`, que é um canal fechado quando o contexto é cancelado. [pt.linux-console](https://pt.linux-console.net/?p=28755)
- Depois de 1 segundo, `cancel()` é chamado e a goroutine recebe esse sinal e imprime a mensagem. [pkg.go](https://pkg.go.dev/context)

## Outro ângulo: pense em “alarme de saída”

Uma boa forma de entender é imaginar que o `context` é um **alarme de saída**.  
A goroutine não para “por mágica”; ela só está programada para olhar o alarme e decidir parar quando ele tocar. [dneto](https://dneto.me/posts/contextos-em-go/)

Ou seja:

- `ctx` = o alarme.
- `cancel()` = apertar o botão que toca o alarme.
- `<-ctx.Done()` = a goroutine esperando o alarme disparar.

## Linha por linha

```go
ctx, cancel := context.WithCancel(context.Background())
```

Aqui você cria um contexto cancelável. O `cancel` é importante porque libera o sinal de cancelamento e deve ser chamado explicitamente. [aprendagolang.com](https://aprendagolang.com.br/o-que-e-e-como-utilizar-o-package-context/)

```go
go func() {
    select {
    case <-ctx.Done():
        fmt.Println("✅ Goroutine recebeu sinal de cancelamento")
    }
}()
```

Essa goroutine está “escutando” o contexto. Quando `ctx.Done()` é fechado, o `select` destrava e a mensagem é exibida. [pt.linux-console](https://pt.linux-console.net/?p=28755)

```go
time.Sleep(1 * time.Second)
cancel()
```

Depois de esperar 1 segundo, o código cancela o contexto. Isso não mata a goroutine na força; apenas sinaliza que ela deve encerrar o trabalho. [dneto](https://dneto.me/posts/contextos-em-go/)

```go
time.Sleep(100 * time.Millisecond)
```

Esse pequeno atraso dá tempo para a goroutine imprimir antes do programa principal terminar.

## O ponto mais importante

O `context` **não é um mecanismo de controle direto de goroutines**; ele é um **sinal padronizado** para dizer “pare o que estiver fazendo”. [pt.linux-console](https://pt.linux-console.net/?p=28755)
Por isso ele aparece muito em servidores, APIs, consultas a banco e tarefas assíncronas que precisam respeitar timeout ou cancelamento. [pkg.go](https://pkg.go.dev/context)

## Uma versão mental mais clara

Pense assim:

1. O programa cria uma tarefa.
2. A tarefa fica trabalhando ou esperando.
3. Se alguém chamar `cancel()`, o contexto avisa.
4. A tarefa, ao perceber isso, encerra sua execução.

Esse é o padrão que você vai ver muito em Go: **passar o contexto para funções e obedecer ao sinal dele**. [aprendagolang.com](https://aprendagolang.com.br/o-que-e-e-como-utilizar-o-package-context/)

## Detalhe útil para estudo

Se você quiser deixar o exemplo mais realista, normalmente a goroutine teria mais de um `case` no `select`, por exemplo para continuar trabalhando ou sair ao cancelar. Isso ajuda a entender que `context` faz mais sentido quando há alguma operação em andamento e não apenas uma espera vazia. [dneto](https://dneto.me/posts/contextos-em-go/)

Se quiser, eu posso reescrever esse mesmo exemplo com uma explicação bem visual, estilo “passo a passo da execução”, ou transformar em um exemplo com `WithTimeout`.