package main

import (
	"fmt"
	"main/contas"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	fmt.Println(conta.Sacar(valorDoBoleto))
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {

	// contaGui := contas.ContaCorrente{titular: "Guilherme", saldo: 200.32}
	// fmt.Println(contaGui)

	// contaGui2 := contas.ContaCorrente{titular: "Guilherme", saldo: 200.33}
	// fmt.Println(contaGui2)

	// fmt.Println(contaGui == contaGui2)

	// conta := contas.ContaCorrente{titular: "Jose", numeroAgencia: 1, numeroConta: 1, saldo: 23.33}
	// fmt.Println(conta)

	// contaBruna := contas.ContaCorrente{"Bruna", 123, 321, 200.33}
	// fmt.Println(contaBruna)

	// var contaDaCris *contas.ContaCorrente
	// contaDaCris = new(contas.ContaCorrente)
	// contaDaCris.titular = "Cris"
	// contaDaCris.saldo = 500.00

	// var contaDaCris2 *contas.ContaCorrente
	// contaDaCris2 = new(contas.ContaCorrente)
	// contaDaCris2.titular = "Cris"
	// contaDaCris2.saldo = 500.00

	// fmt.Println(*contaDaCris == *contaDaCris2)

	// contaSilvia := contas.ContaCorrente{}
	// contaSilvia.titular = "Silvia"
	// contaSilvia.saldo = 500

	// // fmt.Println(contaSilvia)

	// // valorDoSaque := -501

	// // fmt.Println(contaSilvia.sacar(float64(valorDoSaque)))

	// // fmt.Println(contaSilvia)

	// status, valor := contaSilvia.depositar(1000)
	// fmt.Println(status, " Saldo: R$", valor)

	// contaDaSilvia := contas.ContaCorrente{Titular: clientes.Titular{Nome: "Silvia", Cpf: "123.123.123-12", Profissao: "Analista"}}
	// contaDaSilvia.Depositar(1000)
	// contaDoGustavo := contas.ContaCorrente{Titular: clientes.Titular{Nome: "Gustavo", Cpf: "123.123.123-11", Profissao: "Analista"}}

	// status := contaDaSilvia.Transferir(200, &contaDoGustavo)
	// fmt.Println(status)
	// fmt.Println(contaDaSilvia, contaDoGustavo)
	// fmt.Println(contaDaSilvia.ObterSaldo(), contaDoGustavo.ObterSaldo())

	contaDoDenis := contas.ContaPoupanca{}
	contaDoDenis.Depositar(1000)
	PagarBoleto(&contaDoDenis, 60)
	fmt.Println(contaDoDenis.ObterSaldo())

	contaDaLuiza := contas.ContaCorrente{}
	contaDaLuiza.Depositar(50)
	PagarBoleto(&contaDaLuiza, 60)
	fmt.Println(contaDaLuiza.ObterSaldo())
}
