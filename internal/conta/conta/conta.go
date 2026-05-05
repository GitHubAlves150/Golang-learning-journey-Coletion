package conta

import (
	"fmt"
	"regexp"
)

//====================
//Tipo de conta
//====================

type Conta struct{
	numero int
	titular string
	saldo float64
	movimentações []string
}

//construtor da conta
func NovaConta(numero int, titular string) *Conta{
	return &Conta{
		numero: numero,
		titular: titular,
		saldo: 0,
		movimentações: []string{},
	}
}

//1. Depositar(ponteiro)
func (c *Conta) Depositar(valor float64)error{
	if valor <0 {
		fmt.Println("ERRO!", " Valor negativo")
		return nil;
	}
	c.saldo +=valor;
	c.movimentações = append(c.movimentações, fmt.Sprintf("Depósito de R$.2f ", valor));
	return nil;
}
//2. sacar (ponteiro)
func (s *Conta) Sacar(valor float64) error{
	if valor < 0{
		fmt.Println("ERRO!", " Valor negativo- Não dá pra sacar")		
		return nil;
	}
	s.saldo -= valor;
	s.movimentações= append(s.movimentações, fmt.Sprintf("Saque: R$.2f ", valor))
	return nil;
}
//3. Get Saldo
func (G_saldo *Conta) GetSaldo() error{
	fmt.Println("Saldo: R$.3f", G_saldo.saldo);
	return nil;
}
//4. GetTitular(ponteiro)
func (t *Conta) GetTitular()error{
	fmt.Println(">", t.titular);
	return nil;
}
//5. SetTitular (ponteiro)
func (s *Conta) SetTitular(name string) error{
	temNumeros,_ := regexp.MatchString(`\d`, name)


	if temNumeros{
		fmt.Println("Contém números na string")
		return nil;
	}
	s.titular = name;
	return nil
}
//6. GetNUmero (valor)
func (g *Conta) GetNumero(numero int)int{
	if numero < 0{
		fmt.Println("Numero negativo")
		return 0;
	}
	
	return g.numero;
}
//7. Exibir saldo (valor)
func (d *Conta) ExibirSaldo()error{
	fmt.Println("Saldo:", d.saldo);
	return nil;
}
//8. Transferir (ponteiro)
func (t *Conta) Transferir(valor float64, o *Conta, d *Conta)error{
	if valor < 0{
		fmt.Println("ERRO", " Valor negativo para tranasferência");
		return nil;
	}
	d.saldo = o.saldo;
	d.numero = o.numero;
	return nil;
}

//9. String (valor)
func (c Conta) String()string{
	return fmt.Sprintln("Conta %d: %s -R$.2f", c.numero, c.titular, c.saldo);
}