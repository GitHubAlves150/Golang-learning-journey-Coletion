// 5_interface_vazia.go
package main

import "fmt"

// Interface vazia = aceita QUALQUER tipo
type QualquerCoisa interface{}

func main() {
    // ========================================
    // Exemplo 1: interface{} aceita qualquer tipo
    // ========================================
    var x interface{}
    
    x = 42
    fmt.Printf("x é %v, tipo %T\n", x, x)
    
    x = "hello"
    fmt.Printf("x é %v, tipo %T\n", x, x)
    
    x = true
    fmt.Printf("x é %v, tipo %T\n", x, x)
    
    x = struct{ Nome string }{Nome: "João"}
    fmt.Printf("x é %v, tipo %T\n", x, x)
    
    // ========================================
    // Exemplo 2: Type Assertion (converter de volta)
    // ========================================
    fmt.Println("\n--- Type Assertion ---")
    
    var y interface{} = 100
    
    // Tentar converter para int
    valorInteiro, ok := y.(int)
    if ok {
        fmt.Printf("É um int! Valor: %d\n", valorInteiro)
    } else {
        fmt.Println("Não é um int")
    }
    
    // Tentar converter para string (vai falhar)
    valorString, ok := y.(string)
    if ok {
        fmt.Printf("É uma string! Valor: %s\n", valorString)
    } else {
        fmt.Println("Não é uma string")
    }
    
    // ========================================
    // Exemplo 3: Type Switch
    // ========================================
    fmt.Println("\n--- Type Switch ---")
    
    tipos := []interface{}{42, "texto", true, 3.14, struct{ Nome string }{Nome: "João"}}
    
    for _, t := range tipos {
        switch v := t.(type) {
        case int:
            fmt.Printf("Inteiro: %d\n", v)
        case string:
            fmt.Printf("String: %s\n", v)
        case bool:
            fmt.Printf("Booleano: %t\n", v)
        case float64:
            fmt.Printf("Float: %.2f\n", v)
        default:
            fmt.Printf("Tipo desconhecido: %T\n", v)
        }
    }
    
    // ========================================
    // Exemplo 4: Aplicação prática
    // ========================================
    fmt.Println("\n--- Aplicação prática: Processador genérico ---")
    
    dados := []interface{}{
        42,
        "latitude: -23.5505",
        true,
        struct{ Lat, Lon float64 }{Lat: -23.5505, Lon: -46.6333},
    }
    
    for _, dado := range dados {
        ProcessarDado(dado)
    }
}

func ProcessarDado(dado interface{}) {
    switch v := dado.(type) {
    case int:
        fmt.Printf("📊 Processando número: %d\n", v)
    case string:
        fmt.Printf("📝 Processando texto: %s\n", v)
    case bool:
        fmt.Printf("✅✅ Processando booleano: %t\n", v)
    case struct{ Lat, Lon float64 }:
        fmt.Printf("📍 Processando coordenada: [%.6f, %.6f]\n", v.Lat, v.Lon)
    default:
        fmt.Printf("❓ Tipo desconhecido: %T\n", v)
    }
}