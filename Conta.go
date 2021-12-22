package main

import (
	"fmt"
)
import (
	"strconv"
)

type Conta struct {
	titular string
	numAgencia int
	numConta int
	saldo float64

}

var valorConta float64 = 0



func (c *Conta) Depositar(valorDoDeposito float64) (string, float64, string, float64) {
	if valorDoDeposito >= valorConta {
		c.saldo += valorDoDeposito
		return "Saldo agora é", c.saldo, "Valor depositado foi de: ", valorDoDeposito
	} else {
		return  "Valor insuficiente", c.saldo, "Insucesso", valorDoDeposito
	}
}

func (c *Conta) Sacar (valorDoSaque float64) (string, float64, string, float64) {
	if valorDoSaque >= valorConta {
		c.saldo =- valorDoSaque
		return "Seu saldo agora é: ", c.saldo, "Valor do saque foi de: ", valorDoSaque
	} else {
		return "Valor insuficiente: ", c.saldo, "Valor na conta é: ", valorDoSaque
	}
}





func main ()  {
	contaRafa := Conta{titular: "Rafa", numAgencia: 001, numConta: 100, saldo: 0}
	//contaJoiceChata := Conta{titular: "Joice", numAgencia: 001, numConta: 101, saldo: 10}
	contaJully := Conta{titular: "July", numAgencia: 001, numConta: 102, saldo: 0}
	//contaKapih := Conta{titular: "Kapih", numAgencia: 001, numConta: 103, saldo: 1000}

	fmt.Println( "Titular: " + contaRafa.titular + "\nSaldo: " + strconv.FormatFloat(float64(contaRafa.saldo), 'f', 1,32))
	saldoAtual, valor, status, valorDoDeposito  := contaRafa.Depositar(100 )
	fmt.Println(saldoAtual, valor, status, valorDoDeposito)


	fmt.Println( "Titular: " + contaJully.titular + "\nSaldo: " + strconv.FormatFloat(float64(contaJully.saldo), 'f', 1,32))
	saldoAtual, valor, status, valorDoDeposito  = contaJully.Depositar(1500 )
	fmt.Println(saldoAtual, valor, status, valorDoDeposito)

	fmt.Println( "Titular: " + contaJully.titular + "\nSaldo: " + strconv.FormatFloat(float64(contaJully.saldo), 'f', 1,32))
	saldoAtual, valor, status, valorDoSaque := contaJully.Sacar(300 )
	fmt.Println(saldoAtual, valor, status, valorDoSaque)
}

