package banco

import (
    "bancario/internal/conta"
    "fmt"
)

// ============================================
// TIPO BANCO
// ============================================

type Banco struct {
    nome          string
    contas        map[int]*conta.Conta
    proximoNumero int
}

// 0. NovoBanco (funcao construtora, não é método)
func NovoBanco(nome string) *Banco {
    return &Banco{
        nome:          nome,
        contas:        make(map[int]*conta.Conta), // ← OBRIGATÓRIO!
        proximoNumero: 1,                          // ← Primeira conta será número 1
    }
}

// 1. AbrirConta (método do banco)
func (b *Banco) AbrirConta(titular string) *conta.Conta {
    novaConta := conta.NovaConta(b.proximoNumero, titular)
    b.contas[b.proximoNumero] = novaConta
    b.proximoNumero++
    return novaConta
}

// 2. BuscarConta (método do banco) - busca por TITULAR
func (b *Banco) BuscarConta(titular string) (*conta.Conta, error) {
    for _, conta := range b.contas {
        if conta.GetTitular() == titular {  // ← SEM parâmetro!
            return conta, nil
        }
    }
    return nil, fmt.Errorf("conta não encontrada para o titular %s", titular)
}

// 3. FecharConta (método do banco)
func (b *Banco) FecharConta(numero int) error {
    // Verifica se a conta existe
    conta, existe := b.contas[numero]
    if !existe {
        return fmt.Errorf("conta %d não existe", numero)
    }

    // Verifica se o saldo é zero
    if conta.GetSaldo() != 0 {  // ← Use GetSaldo() ou Getsaldo() conforme seu conta.go
        return fmt.Errorf("conta %d tem saldo R$%.2f. Não é possível removê-la", 
            numero, conta.GetSaldo())
    }

    // Remove a conta do mapa
    delete(b.contas, numero)
    
    return nil
}

// 4. TotalDepositado (método do banco) - SEM parâmetro!
func (b *Banco) TotalDepositado() float64 {
    var total float64 = 0.0

    for _, conta := range b.contas {
        total += conta.GetSaldo()  // ← Use GetSaldo() ou Getsaldo()
    }
    return total
}

// 5. ContasPorTitular (método do banco)
func (b *Banco) ContasPorTitular(titular string) []*conta.Conta {
    var contasEncontradas []*conta.Conta

    for _, conta := range b.contas {
        if conta.GetTitular() == titular {  // ← SEM parâmetro!
            contasEncontradas = append(contasEncontradas, conta)
        }
    }
    return contasEncontradas
}

// ============================================
// MÉTODOS ADICIONAIS ÚTEIS
// ============================================

// GetNome retorna o nome do banco
func (b *Banco) GetNome() string {
    return b.nome
}

// TotalContas retorna quantas contas o banco possui
func (b *Banco) TotalContas() int {
    return len(b.contas)
}

// BuscarContaPorNumero busca uma conta pelo número
func (b *Banco) BuscarContaPorNumero(numero int) (*conta.Conta, error) {
    conta, existe := b.contas[numero]
    if !existe {
        return nil, fmt.Errorf("conta %d não encontrada", numero)
    }
    return conta, nil
}

// ListarTodasContas exibe todas as contas do banco
func (b *Banco) ListarTodasContas() {
    if len(b.contas) == 0 {
        fmt.Println("   Nenhuma conta ativa")
        return
    }
    
    fmt.Printf("\n=== CONTAS DO %s ===\n", b.nome)
    for numero, conta := range b.contas {
        fmt.Printf("   Conta %d: %s - R$%.2f\n", 
            numero, conta.GetTitular(), conta.GetSaldo())
    }
}