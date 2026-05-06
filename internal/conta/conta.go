package conta

import (
    "bancario/pkg/utils"
    "fmt"
)

// ============================================
// TIPO CONTA
// ============================================

type Conta struct {
    Numero        int
    Titular       string
    Saldo         float64
    Movimentacoes []string
}

// NovaConta é o construtor
func NovaConta(numero int, titular string) *Conta {
    return &Conta{
        Numero:        numero,
        Titular:       titular,
        Saldo:         0,
        Movimentacoes: []string{},
    }
}

// 1. Depositar
func (c *Conta) Deposito(valor float64) error {
    if valor <= 0 {
        return fmt.Errorf("valor do depósito deve ser positivo. Recebido: R$%.2f", valor)
    }
    
    c.Saldo += valor
    c.Movimentacoes = append(c.Movimentacoes, 
        fmt.Sprintf("Depósito: +R$%.2f", valor))
    return nil
}

// 2. Sacar
func (c *Conta) Sacar(valor float64) error {
    if valor <= 0 {
        return fmt.Errorf("valor do saque deve ser positivo. Recebido: R$%.2f", valor)
    }
    
    if valor > c.Saldo {
        return fmt.Errorf("saldo insuficiente. Saldo atual: R$%.2f", c.Saldo)
    }
    
    c.Saldo -= valor
    c.Movimentacoes = append(c.Movimentacoes, 
        fmt.Sprintf("Saque: -R$%.2f", valor))
    return nil
}

// 3. GetSaldo (com 'S' maiúsculo - padrão Go)
func (c Conta) GetSaldo() float64 {
    return c.Saldo
}

// 4. GetTitular (SEM parâmetros!)
func (c Conta) GetTitular() string {
    return c.Titular
}

// 5. SetTitular
func (c *Conta) SetTitular(nome string) error {
    if !utils.EhLetra(nome) {
        return fmt.Errorf("nome deve conter apenas letras. Recebido: %s", nome)
    }
    c.Titular = nome
    return nil
}

// 6. GetNumero
func (c Conta) GetNumero() int {
    return c.Numero
}

// 7. ExibirExtrato
func (c Conta) ExibirExtrato() string {
    if len(c.Movimentacoes) == 0 {
        return "Nenhuma movimentação realizada"
    }
    
    resultado := fmt.Sprintf("=== EXTRATO DA CONTA %d ===\n", c.Numero)
    resultado += fmt.Sprintf("Titular: %s\n", c.Titular)
    resultado += fmt.Sprintf("Saldo atual: R$%.2f\n", c.Saldo)
    resultado += "\nMOVIMENTAÇÕES:\n"
    
    for _, mov := range c.Movimentacoes {
        resultado += mov + "\n"
    }
    
    return resultado
}

// 8. Transferir
func (c *Conta) Transferir(destino *Conta, valor float64) error {
    if destino == nil {
        return fmt.Errorf("conta destino não existe")
    }
    
    // Tenta sacar da conta origem
    if err := c.Sacar(valor); err != nil {
        return fmt.Errorf("transferência falhou: %v", err)
    }
    
    // Tenta depositar na conta destino
    if err := destino.Deposito(valor); err != nil {
        // Rollback: devolve o dinheiro para a conta origem
        c.Deposito(valor)
        return fmt.Errorf("transferência falhou no destino: %v", err)
    }
    
    // Registra transferência nas duas contas
    c.Movimentacoes = append(c.Movimentacoes,
        fmt.Sprintf("Transferência enviada: -R$%.2f para conta %d", valor, destino.Numero))
    destino.Movimentacoes = append(destino.Movimentacoes,
        fmt.Sprintf("Transferência recebida: +R$%.2f da conta %d", valor, c.Numero))
    
    return nil
}

// String - implementa fmt.Stringer para exibir a conta
func (c Conta) String() string {
    return fmt.Sprintf("Conta %d: %s - R$%.2f", c.Numero, c.Titular, c.Saldo)
}