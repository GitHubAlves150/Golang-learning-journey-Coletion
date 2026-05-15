# 🎯 Exercício 1: Contador Simples com Race Condition

## 🎯 Objetivo do Exercício
Entender o problema básico de race condition quando múltiplas goroutines acessam uma variável compartilhada sem sincronização.
## 📝 Enunciado

Você tem uma variável contador que começa em 0. Você precisa criar 100 goroutines que cada uma vai incrementar o contador 1000 vezes.  
No final, o valor esperado é 100.000 (100 × 1000).  
Implemente sem usar nenhum mecanismo de sincronização (sem Mutex, sem canal).  

## 🧪 Execute e Observe  
## 🔍 O que está acontecendo?
