# 🎯 Exercício 3: Deadlock com Mutex

## 🎯 Objetivo do Exercício
Entender o que é DEADLOCK (abraceamento) - uma situação onde as goroutines ficam travadas para sempre esperando umas pelas outras.

- Deadlock = "O programa trava e nunca mais sai"  


## 📝 Enunciado

Você tem duas contas bancárias e vai fazer uma transferência entre elas. Mas o código tem um erro que causa deadlock!

Duas goroutines tentam transferir dinheiro uma para a outra ao mesmo tempo, cada uma segurando um mutex e esperando pelo outro.

Sua missão: Identificar por que o programa trava e corrigir o deadlock.  


## 🧪 Execute e Observe  
## 🔍 O que está acontecendo?

O programa quebrou porque duas goroutines tentaram escrever no map ao mesmo tempo.  

**Remova os comentários e veja o resultado. Tenta debugar com F5 para ver o comportamento do código**
