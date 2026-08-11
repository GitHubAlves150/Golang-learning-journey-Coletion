# 📡 Exemplo de Channels em Go

## 🎯 O que este código demonstra

Comunicação básica entre **goroutines** usando **channels** (canais).

## 🧠 Conceito

> **Channel é um "tubo" que conecta goroutines. O que entra por um lado, sai do outro.**

## 📝 Código

```go
package main

import (
	"fmt"
	"time"
)

// ============================================
// channels
// ============================================

func main() {

	ch := make(chan string, 3) //declaração do channel

	go func()  {
		fmt.Println("escrevendo no channel com buffer")
		ch <- "menssagem 1"
		fmt.Println("Menssagem 1 enviada")
		ch <- "menssagem 2"
		fmt.Println("Menssagem 2 enviada")
		ch <- "menssagem 3"
		fmt.Println("Menssagem 3 enviada")
		ch <- "menssagem 4"
		fmt.Println("Menssagem 4 enviada")
	}()

	time.Sleep(2 *time.Second)

	go func ()  {
		for i:=0; i<4;i++{
			msg:= <-ch
			fmt.Println("Recebido: ", msg)
			time.Sleep(1*time.Second)
		}

	}()


	time.Sleep(10*time.Second)


	fmt.Println("....FIM.....")
}


```

# ⚡ Sincronização Automática  


**Sem time.Sleep, sem WaitGroup!**

- A escrita bloqueia até alguém ler  
- A leitura bloqueia até alguém escrever  

```go

ch <- "mensagem"  // ⏸️ Trava até alguém ler
msg := <-ch       // ⏸️ Trava até alguém escrever

```  
# 🖼️ Visualização  

```go

Goroutine 1                    Goroutine 2 (main)
     │                              │
     ├── ch <- "Olá!"               │
     │      │                        │
     │      └──────[ 📦 ]──────────→ │
     │                              ├── msg := <-ch
     │                              │
     └── (desbloqueia)              └── (desbloqueia)

```
```go
✅ Resumo
Operação	            Símbolo	O que faz
Criar	                make(chan tipo)	Cria um novo canal
Escrever	            ch <- valor	Envia valor para o canal
Ler	                    var := <-ch	Recebe valor do canal
``` 

