# 📚 Código Comentado Linha por Linha - API de Produtos com Echo + Viper

## 🎯 Visão Geral

Este código cria uma **API REST** para gerenciar produtos, utilizando:
- **Echo** → Framework web para criar o servidor e rotas
- **Viper** → Gerenciador de configurações (porta do servidor)
- **Memória** → Banco de dados temporário (slice de produtos)

---

## 📦 PARTE 1: IMPORTS (Bibliotecas)

```go
package main

