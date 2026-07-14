


## 🏛️ O que são Microserviços?

A arquitetura de Microserviços é uma abordagem de design de software onde uma aplicação grande e complexa é dividida em vários serviços menores, independentes e especializados. Cada microserviço roda seu próprio processo, possui seu próprio banco de dados isolado e cuida de apenas uma funcionalidade de negócio (como "Pagamentos", "Autenticação" ou "Estoque"). Eles se comunicam entre si através de protocolos leves, como APIs HTTP/REST, gRPC ou mensageria (filas).

O oposto de microserviços é o Monólito, onde todo o código do sistema (frontend, backend, banco de dados, regras de negócio) mora dentro de um único projeto e deploy.

Para visualizar como o sistema deixa de ser um bloco único e se transforma em uma rede conectada, veja a diferença estrutural abaixo:

![alt text](licensed-image.jpeg)

## 🛠️ O que ela resolve? (As Dores do Mercado)

Em empresas que crescem muito, o Monólito começa a dar problemas graves que os Microserviços resolvem:

Escalabilidade Seletiva: Se o seu sistema é um e-commerce e chega a Black Friday, o serviço de "Busca de Produtos" e "Pagamentos" vai sofrer muito mais acessos do que o serviço de "Alteração de Cadastro". No Monólito, você precisa duplicar o sistema inteiro na nuvem para aguentar o tranco (gastando muito dinheiro). Com microserviços, você escala (cria mais instâncias) apenas o serviço de Pagamentos e Busca.

Gargalo de Equipes (Deploy Independente): Em um monólito com 50 desenvolvedores mexendo no mesmo código, um time pode quebrar o código do outro. Para subir uma alteração boba no estoque, é preciso gerar o deploy do sistema inteiro. Com microserviços, o time de pagamentos faz deploy do seu serviço na terça-feira sem nem avisar o time de estoque, pois os projetos são totalmente separados.

Resiliência: Se o serviço de "Recomendações de Produtos" cair por falta de memória, o e-commerce não para de funcionar. O cliente ainda consegue ver o carrinho, fazer login e pagar. O sistema falha de forma graciosa.

## ⏳ Quando começou?

O termo "Microservices" ganhou força e foi formalizado entre 2011 e 2012, impulsionado por pioneiros da indústria como Martin Fowler e James Lewis, além de grandes empresas como Netflix, Amazon e SoundCloud, que atingiram o limite do que seus monólitos conseguiam aguentar em termos de tráfego mundial.

A Netflix, por exemplo, começou sua transição em 2008 após uma falha grave em seu banco de dados monolítico que parou a empresa por dias, concluindo a migração completa por volta de 2016.

##💻 Para quais linguagens ela é usada?

Os microserviços são agnósticos a linguagens de programação. Como cada serviço é independente e se comunica por protocolos de rede padrão (como JSON sobre HTTP ou Protocol Buffers sobre gRPC), você pode construir um ecossistema poliglota:

Go (Golang): Tornou-se uma das linguagens mais populares e usadas para microserviços no mundo (usada por Uber, Twitch, Mercado Livre). Por ser compilada, iniciar em milissegundos, consumir pouquíssima memória e ter concorrência nativa (goroutines), ela é perfeita para serviços que precisam de performance extrema e baixo custo de nuvem.

Java / Kotlin (Spring Boot): Muito forte em grandes corporações e bancos devido à maturidade do ecossistema.

Node.js (TypeScript): Amplamente usada por times que querem unificar a linguagem do Frontend (React) com o Backend, ideal para APIs de I/O intensivo.

Python / C#: Também muito presentes em cenários específicos (Python para IA/Dados e C# em ambientes corporativos Microsoft).

Em resumo: você poderia ter o serviço de Pedidos em Go, o de Pagamentos em Java e o de Notificações em Node.js, e todos conversariam entre si perfeitamente.

## 🗺️ Vamos inaugurar a Fase 1: Fundações de Microserviços.

Como vimos na reintrodução acima, o grande desafio dessa arquitetura não é criar os serviços em si, mas sim fazer com que eles conversem de forma eficiente e resiliente.

Para o nosso cenário prático, vamos construir dois microserviços independentes em Go:

- Microserviço de Pedidos (orders-service): Recebe a requisição do usuário criando um pedido e salva no banco.

- Microserviço de Pagamentos (payments-service): Processa o pagamento desse pedido.

## 🗺️ Como eles vão se comunicar?

Nesta primeira etapa, vamos implementar a Comunicação Síncrona via gRPC.

### O que é gRPC?

Em vez de usar o REST tradicional (onde o Go precisa transformar uma Struct em JSON, mandar pela rede, e o outro serviço precisa ler o JSON e transformar em Struct de novo), o gRPC (criado pelo Google) usa Protocol Buffers.

Os dados são transmitidos em formato binário super compactado através do protocolo HTTP/2. Ele chega a ser até 10 vezes mais rápido que o REST tradicional, sendo o padrão ouro de mercado para microserviços conversarem entre si por baixo dos panos.

Para ver o fluxo completo de como as requisições vão transitar no nosso sistema, desde a chamada externa até a comunicação binária interna, observe o diagrama abaixo:

### 🏗️ Passo 1: Preparando o Terreno e as Dependências

Como desenvolvedor pleno, você sabe que cada microserviço precisa ser um projeto Go completamente isolado, com seu próprio arquivo go.mod.

```bash
lab-microservices/
├── orders-service/     # Projeto Go do serviço de Pedidos
└── payments-service/   # Projeto Go do serviço de Pagamentos
```

Para trabalharmos com gRPC no Go, precisamos instalar o compilador de protocolo (protoc) na sua máquina e os plugins do Go.

Instalação das ferramentas no sistema:

```python sudo apt install -y protobuf-compiler``` 

Instalando os plugins no Go:

Abra o seu terminal e execute estes comandos globais para que o Go aprenda a gerar códigos gRPC:

```bash 
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```


## 📄 Passo 2: Definindo o Contrato (.proto)

No gRPC, o contrato nasce antes do código. Nós escrevemos um arquivo .proto que descreve quais funções o serviço de pagamentos terá e quais dados ele aceita.

Crie um arquivo chamado payment.proto dentro de uma pasta comum ou dentro do payments-service:

⚡ Passo 3: Gerando o Código Go (protoc)

Agora vem a mágica! Com o arquivo corrigido e salvo, vamos rodar o comando que vai ler esse .proto e criar toda a estrutura que o Go precisa para rodar o gRPC (as structs, funções e o cliente/servidor).

Abra o terminal.

- Vá até a pasta onde está o seu arquivo (dentro de payments-services).

- Execute o seguinte comando:

```bash protoc --go_out=. --go-grpc_out=. payment.proto```

## 🧐 O que esse comando faz?

    --go_out=.: Diz para o compilador gerar as structs normais do Go (Request e Response) na pasta atual baseada no go_package (que definimos como ./pb).

    --go-grpc_out=.: Diz para gerar as funções específicas de servidor e cliente gRPC do Go.

Após rodar esse comando, você vai notar que uma nova pasta chamada pb vai aparecer no seu projeto com dois arquivos dentro: payment.pb.go e payment_grpc.pb.go. Nunca mexa nesses arquivos gerados, eles são gerenciados pelo gRPC.

Rode o comando no seu terminal. Deu tudo certo e a pasta pb foi criada? Se sim, me avise para criarmos o servidor de Pagamentos em Go que vai ler esse contrato!

Agora que a pasta pb foi criada com os arquivos gerados automaticamente, nós temos o "esqueleto" do gRPC pronto.

O nosso próximo passo é dar vida a esse esqueleto. Vamos criar o Servidor do Microserviço de Pagamentos (payments-service). Ele vai escutar em uma porta lógica (geralmente usamos a 50051 para gRPC) esperando que o serviço de Pedidos envie uma ordem de pagamento.

## 🏗️ Estrutura de Pastas do payments-services

O seu projeto de pagamentos vai ficar organizado dessa forma agora: 

```bash payments-services/
├── pb/
│   ├── payment.pb.go       # Gerado pelo protoc
│   └── payment_grpc.pb.go  # Gerado pelo protoc
├── main.go                 # Inicia o servidor gRPC na porta 50051
├── payment.proto           # O contrato que você criou

```
## 💻 Escrevendo o Servidor (main.go)

Antes de colar o código, precisamos inicializar o módulo do Go e baixar a dependência oficial do gRPC. No terminal, dentro da pasta payments-services, rode:

``` bash 
go mod init payments-services
go get google.golang.org/grpc
go mod tidy
``` 
Execute go run main.go e deixa o servidor rodando.

___

### Servidor gRPC de pé e escutando na porta :50051. Metade da Fase 1 está concluída.

Agora vamos criar o Microserviço de Pedidos (orders-services). Ele será o responsável por receber uma requisição, simular a criação de um pedido e fazer uma chamada gRPC interna para o serviço de Pagamentos que você acabou de ligar.

## 🏗️ Passo 1: Copiar o Contrato (payment.proto)

Para que o serviço de Pedidos saiba como falar com o serviço de Pagamentos, ele precisa conhecer o mesmo contrato

- Crie uma pasta chamada pb dentro de orders-services.
- Copie o arquivo payment.proto que você usou no outro projeto e cole dentro de orders-services.
- Abra o terminal, navegue até a pasta orders-services e rode o mesmo comando do protoc para gerar os arquivos Go locais dele:

```bash cd ../orders-services
protoc --go_out=. --go-grpc_out=. payment.proto
```
📦 Passo 2: Inicializar o Módulo de Pedidos

Ainda dentro da pasta orders-services, vamos inicializar o módulo do Go e baixar as dependências do gRPC para esse serviço:

```bash
go mod init orders-services
go get google.golang.org/grpc
go mod tidy
```

## 💻 Escrevendo o Cliente (main.go de Pedidos)

Agora, crie o arquivo main.go dentro de orders-services. Este código vai se conectar no servidor :50051, disparar um pedido de teste e ler a resposta binária.

Perceba o uso do context.WithTimeout. Em microserviços, se o serviço de pagamentos estiver lento, o serviço de pedidos não pode ficar travado para sempre. O Context cancela a chamada se ela demorar mais de 3 segundos!

## 🏃‍♂️ O Teste do Ecossistema

Agora você tem os dois lados da moeda. Vamos fazê-los conversar:

- Deixe o terminal do payments-services ligado e rodando.
- Abra um segundo terminal e vá para a pasta do orders-services:

## 🎯 O que vai acontecer:

No terminal de Pedidos, você verá a resposta impressa na tela dizendo que o pagamento foi Aprovado. E se você olhar o terminal de Pagamentos, verá o log pipocando na hora informando que ele recebeu a chamada e processou o valor!

Faça esse disparo. 

![alt text](<Screenshot From 2026-07-13 17-36-00.png>)

___

![alt text](<Screenshot From 2026-07-13 17-36-25.png>)
