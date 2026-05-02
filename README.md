# Golang
## Introdução
Este README, é uma breve anotações que fiz ao longo do mês de Abril de 2026 sobe Golang e suas particularidade.  





começa com  




















## Goroutine

Imagine que você tem um único funcionário (a CPU) num restaurante. Ele precisa:

    Atender telefone ☎️

    Levar pratos aos clientes 🍽️

    Anotar pedidos ✍️

Sem goroutine: ele faz uma coisa de cada vez. Se atender telefone, todos os clientes esperam. Isso é lento e ineficiente.

Com goroutine: é como se esse funcionário multiplicasse (ganhasse "cópias leves" de si mesmo). Cada goroutine cuida de uma tarefa:

    Uma atende telefone

    Outra leva pratos

    Outra anota pedidos

E o "gerente" (scheduler do Go) decide quem faz o quê e por quanto tempo, trocando muito rápido entre elas. Parece que tudo acontece ao mesmo tempo, mas na verdade é intercalado (concorrente).

    🧠 Resumo leigo: goroutine é uma tarefa super leve que roda "ao mesmo tempo" que outras, sem precisar de muitos recursos.

<br><br>

---



1. O que é goroutine? 📚
🔹 Nível leigo (analogia do restaurante)

Imagine que você tem um único funcionário (a CPU) num restaurante. Ele precisa:

    Atender telefone ☎️

    Levar pratos aos clientes 🍽️

    Anotar pedidos ✍️

Sem goroutine: ele faz uma coisa de cada vez. Se atender telefone, todos os clientes esperam. Isso é lento e ineficiente.

Com goroutine: é como se esse funcionário multiplicasse (ganhasse "cópias leves" de si mesmo). Cada goroutine cuida de uma tarefa:

    Uma atende telefone

    Outra leva pratos

    Outra anota pedidos

E o "gerente" (scheduler do Go) decide quem faz o quê e por quanto tempo, trocando muito rápido entre elas. Parece que tudo acontece ao mesmo tempo, mas na verdade é intercalado (concorrente).

    🧠 Resumo leigo: goroutine é uma tarefa super leve que roda "ao mesmo tempo" que outras, sem precisar de muitos recursos.

🔹 Nível Engenharia Elétrica/Computação

Do ponto de vista de sistemas e arquitetura:

    Goroutine é uma thread leve gerenciada pelo runtime do Go, não pelo sistema operacional.

    Cada goroutine ocupa ~2KB de stack inicial (vs. 1-8MB de uma thread OS).

    O scheduler do Go (M:N scheduling) mapeia N goroutines para M threads do SO.

    Ele usa cooperação preemptiva: em Go 1.14+, goroutines são preemptíveis em pontos seguros (ex: chamadas a função ou quando o scheduler decide).

    Context switching entre goroutines é muito mais barato que entre threads do SO, pois não envolve chamadas de sistema (syscalls) e preserva menos contexto.

Fundamento elétrico subjacente: o scheduler usa timers e interrupções virtuais baseadas em sinais do sistema operacional (ex: SIGURG no Linux) para forçar preempção. O hardware (CPU) continua executando instruções sequenciais, mas a ilusão de paralelismo vem da alternância extremamente rápida.
2. No back-end, o que a goroutine faz?

No desenvolvimento back-end (APIs, microsserviços, workers):

    Manipulação de requisições HTTP concorrentes
    Cada requisição http.Request pode rodar em sua própria goroutine. O servidor web do Go (net/http) já faz isso por padrão.

    I/O não-bloqueante
    Enquanto uma goroutine aguarda dados do banco de dados, rede ou disco, o scheduler automaticamente pausa ela e executa outra. Em outras linguagens, isso exigiria async/await manual.

    Processamento assíncrono
    Ex: salvar um arquivo, enviar e-mail, processar fila — tudo rodando em goroutines separadas, sem travar a resposta imediata ao cliente.

    Padrões concorrentes
    Como pipelines (fan-in/fan-out), workers pools, timeouts com select e canais.

Exemplo simples:
go

go processarPedido(pedido) // roda em background, sem travar a resposta

3. Existe apenas no Go?

Não. O conceito de threads leves gerenciadas pelo runtime existe em outras linguagens:
Linguagem	Nome	Notas
Erlang/Elixir	Processos	Isolados, comunicação via mensagens
Kotlin	Coroutines	Leves, suportadas a nível de linguagem
Java	Virtual Threads (Project Loom)	Desde Java 21, similar a goroutines
C++	Coroutines (C++20)	Leves, mas sem scheduler automático
Rust	Tokio/async-std (tasks)	Green threads via bibliotecas assíncronas

O diferencial do Go: goroutines vêm embutidas na linguagem desde o início, com canais (chan) nativos e scheduler automático, sem necessidade de bibliotecas externas ou palavras-chave async/await.
4. O que é concorrência?

Concorrência é a capacidade de lidar com várias tarefas ao mesmo tempo, mas não necessariamente executá-las simultaneamente no exato mesmo instante.

    Concorrência ≠ Paralelismo

        Concorrência: estruturação do programa em tarefas que podem ser executadas em ordem arbitrária (intercaladas).

        Paralelismo: execução simultânea real (múltiplos núcleos/CPUs).

Exemplo:

    Concorrente: você cozinha e responde mensagens no celular, alternando entre as duas.

    Paralelo: você cozinha e alguém responde mensagens ao mesmo tempo, cada um numa CPU.

5. É chamado de concorrência apenas no Go?

Não. Concorrência é um conceito geral da ciência da computação, muito anterior ao Go:

    Anos 1960: C. A. R. Hoare (CSP – Communicating Sequential Processes, que inspirou os canais do Go).

    Sistemas operacionais: processos concorrentes, threads.

    Banco de dados: controle de concorrência (transações).

O que o Go fez foi popularizar o termo no contexto de linguagens modernas, principalmente com a famosa frase de Rob Pike:

    "Concorrência é sobre lidar com muitas coisas ao mesmo tempo. Paralelismo é sobre fazer muitas coisas ao mesmo tempo."

Portanto: concorrência não é exclusividade Go, mas o Go foi desenhado desde o início para torná-la simples e eficiente com goroutines e canais.
Resumo rápido 🧠
Pergunta	Resposta
Goroutine é só no Go?	Não, mas é nativa e com scheduler próprio
Concorrência é só no Go?	Não, conceito universal
Goroutine no back-end?	Roda requisições, I/O, tarefas assíncronas
Leigo → goroutine?	Vários atendentes leves no restaurante
Engenharia → goroutine?	Contexto leve (2KB stack), mapeamento M:N, preempção

Se quiser, posso mostrar exemplos práticos de concorrência vs paralelismo com código Go.
