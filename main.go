package main

import (
	"fmt"

	"github.com/maykonmori1993/alura_ban/contas"
)

func main() {

	contaDaSilvia := contas.ContaCorrente{
		Titular:       "Silvia",
		NumeroAgencia: 0,
		NumeroConta:   0,
		Saldo:         300,
	}
	contaDoGustavo := contas.ContaCorrente{
		Titular:       "Gustavo",
		NumeroAgencia: 0,
		NumeroConta:   0,
		Saldo:         100,
	}

	status := contaDaSilvia.Transferir(-200, &contaDoGustavo)

	fmt.Println(status)
	fmt.Println(contaDaSilvia)
	fmt.Println(contaDoGustavo)

}
