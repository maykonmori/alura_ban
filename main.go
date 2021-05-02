package main

import (
	"fmt"

	"github.com/maykonmori1993/alura_ban/contas"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaDoDenis := contas.ContaPoupanca{}
	contaDoDenis.Depositar(100)
	PagarBoleto(&contaDoDenis, 60)

	fmt.Println(contaDoDenis.ObterSaldo())

	contaDaPati := contas.ContaCorrente{}
	contaDaPati.Depositar(500)
	PagarBoleto(&contaDaPati, 400)
	fmt.Println(contaDaPati.ObterSaldo())
}
