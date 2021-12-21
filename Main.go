package main

import "fmt"

type ContaCorrente struct {
	titular string
	numAgencia int
	numConta int
	saldo float64
}



type Conta struct {
	saldo float64
}

func (c *Conta) depositarDezReais() float64 {
	  return c.saldo + 10
	}

func main ()  {
	contaTeste := Conta{saldo: 10}
	contaTeste.depositarDezReais()

	fmt.Println(contaTeste)

}