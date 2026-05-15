# 🎯 Exercício 2: Map Compartilhado com Race Condition

## 🎯 Objetivo do Exercício
Entender o problema de race condition quando múltiplas goroutines acessam um map compartilhado sem sincronização.

## 📝 Enunciado

Você tem um map vazio que vai guardar o número de votos para 3 candidatos (A, B, C).

Você precisa criar 1000 goroutines que cada uma vai:

1.  Escolher um candidato aleatório (A, B ou C)

2.  Incrementar o voto desse candidato no map

No final, o total de votos deve ser 1000 (um por goroutine).

Implemente sem usar nenhum mecanismo de sincronização (sem Mutex, sem canal). 

## 🧪 Execute e Observe  
## 🔍 O que está acontecendo?

O programa quebrou porque duas goroutines tentaram escrever no map ao mesmo tempo.  

**Remova os comentários e veja o resultado. Tenta debugar com F5 para ver o comportamento do código**
