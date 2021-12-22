package main

import "fmt"
import (
	"strconv"
)

type Conta struct {
	titular string
	numAgencia int
	numConta string
	saldo float64
}



func (c *Conta) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Depositado", c.saldo
	} else {
		return  "Valor insuficiente", c.saldo
	}
}



func main ()  {
	contaRafa := Conta{titular: "Rafa", numAgencia: 001, numConta: "100", saldo: 1000}
	contaJully := Conta{titular: "July", numAgencia: 001, numConta: "100", saldo: 1000}
	var saldoTeste int

	fmt.Println( "Titular: " + contaRafa.titular + "\nNumero da Conta: " + contaRafa.numConta)
	fmt.Println( "Titular: " + contaJully.titular + "Saldo: " + strconv.Itoa(saldoTeste))
	status, valor := contaJully.Depositar(2000)
	fmt.Println(status, valor)
}

