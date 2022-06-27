package main

import (
	"fmt"

	"golang/Go/cmd/core/navigate"
)

func main() {

	//var contaDoDenis = clientes.Titular{}
	//
	//repositories.ContaRepositories.Depositar()
	//PagarBoleto(&contaDoDenis, 60)
	//
	//fmt.Println(contaDoDenis.ObterSaldo())
	//
	//contaDaLuisa := contas.ContaCorrente{}
	//contaDaLuisa.Depositar(500)
	//PagarBoleto(&contaDaLuisa, 1000)
	//
	//fmt.Println(contaDaLuisa.ObterSaldo())

	for {
		navigate.ExibeIntroducao()
		navigate.ExibeMenu()
		comando := navigate.LerComando()

		switch comando {
		case 1:
			navigate.IniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 3:
			navigate.ExibirNomes()
		case 0:
			fmt.Println("Saindo do programa...")
		default:
			fmt.Println("Não conheço este comando")
		}
	}
}
