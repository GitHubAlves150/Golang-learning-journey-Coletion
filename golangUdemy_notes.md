# Exemplo: Processamento de dados com e sem concorrẽncia   

📋 Cenário: Processamento de Pedidos

imagine que você tem 3 etapas para processar um pedido:  
-  1. Validar(verificarse está tudo ok)  
-  2. Processar(Calcular total, aplicar descontos) 
-  3. Notificar(envia email, SMS)

---

**Versão 2: Nesta branch será mostrada COM gouroutines + waitgroup (correto)** 


**Saída**  

```go
Iniciando processamento COM goroutine+waitgroup....

Pedido : 123 Processamento concluído 123
Pedido  123 : Notificacão enviada
Pedido : 123 Validação concluída 123
Tempo total:  1.000495267s
....FIM.....

```