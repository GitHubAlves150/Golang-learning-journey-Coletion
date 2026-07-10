package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	//=============================================================================================
	// PASSO 1. DEFINIR A CHAVE SECRETA
	//=============================================================================================

	//È como uma "senha" para assinar o token. Em produção, coloque em uma variavel de ambiente!

	chaveSecreta := []byte("minha-chave-secreta-super-secreta")

	//=============================================================================================
	//2. Criar os dados que vão dentro do token
	//=============================================================================================
	//São as "informações" que você quer guardar
	//EX: ID do usuário, email, etc

	dados := jwt.MapClaims{
		"user_id": 4455,                                 //ID do usuario
		"email":   "joao@hotmail.com",                   //email do usuario
		"exp":     time.Now().Add(1 * time.Hour).Unix(), //expira em um hora
		"iat":     time.Now().Unix(),                    //criado em (issued at)
	}

	//=============================================================================================
	//3. Criar o token com os dados e o método de assinatura
	//=============================================================================================

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, dados)

	//=============================================================================================
	//4. Assinar o token com a chave secreta
	//=============================================================================================

	//È aqui que o token é "criptografado"
	tokenString, err := token.SignedString(chaveSecreta)
	if err != nil {
		fmt.Errorf("Erro ao gerar o token: ", err)
		return
	}

	//=============================================================================================
	// PASSO 5: Exibir o token gerado
	//=============================================================================================

	fmt.Println("========================================")
    fmt.Println("🔑 TOKEN JWT GERADO:")
    fmt.Println("========================================")
    fmt.Println(tokenString)
    fmt.Println("========================================")
    fmt.Println()
    fmt.Println("📝 O token tem 3 partes separadas por pontos:")
    fmt.Println("   HEADER.PAYLOAD.SIGNATURE")
}
