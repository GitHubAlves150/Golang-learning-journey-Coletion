# 🎯 Operações em arquivos  

```go
 Conceito	               O que é	                        Por que é importante  
 defer file.Close()	       Fecha arquivo automaticamente	EVITA VAZAMENTO DE MEMÓRIA  
os.Open()	               Abre arquivo para leitura	    Retorna *File e error
bufio.Scanner	           Lê linha por linha	            Eficiente para arquivos grandes
os.ReadFile()	           Lê arquivo inteiro de uma vez	Simples para arquivos pequenos 


``` 

**Verificar erros	Sempre if err != nil OBRIGATÓRIO em Go**

# 📂 Estrutura de Arquivo para Testes  
**Crie um arquivo dados.txt para testar:**

```go
linha 1: Hello World
linha 2: Go é incrível
linha 3: Aprendendo arquivos
linha 4: Última linha
```

## 1. 📖 Método 1: os.ReadFile() - Ler arquivo INTEIRO  

- Quando usar: Arquivos PEQUENOS (< 10MB) que cabem em memória 

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    // 1. Lê o arquivo inteiro de uma vez
    dados, err := os.ReadFile("dados.txt")
    if err != nil {
        fmt.Println("Erro ao ler arquivo:", err)
        return
    }
    
    // 2. Converte para string e exibe
    fmt.Println(string(dados))
}
```
Vantagens: ✅ Simples,
✅ Uma linha
✅ Fecha automaticamente  
**Desvantagens:** ❌ Carrega tudo na memória, ❌ Ruim para arquivos gigantes

# 2. 📖 Método 2: os.Open() + defer - Abrir e LER COM SEGURANÇA  

**Quando usar: PRÁTICA RECOMENDADA para a maioria dos casos**  

```go 

package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    // 1. Abrir o arquivo
    file, err := os.Open("dados.txt")
    if err != nil {
        fmt.Println("Erro ao abrir arquivo:", err)
        return
    }
    defer file.Close()  // ← ESSENCIAL! Fecha automaticamente no final
    
    // 2. Criar scanner para ler linha por linha
    scanner := bufio.NewScanner(file)
    
    // 3. Ler cada linha
    for scanner.Scan() {
        linha := scanner.Text()
        fmt.Println(linha)
    }
    
    // 4. Verificar se houve erro na leitura
    if err := scanner.Err(); err != nil {
        fmt.Println("Erro ao ler arquivo:", err)
    }
}

```


