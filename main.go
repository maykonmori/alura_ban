package main

import (
	"banco/contas"
	"fmt"
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
