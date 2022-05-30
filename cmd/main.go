package main

import (
	"fmt"
	"golang/Go/cmd/core/contas"
	clientes "golang/Go/cmd/core/entities"
	"golang/Go/cmd/core/repositories"
)

func main() {

	var contaDoDenis = clientes.Titular{}
depositar:
	repositories.ContaRepositories.Depositar()
	PagarBoleto(&contaDoDenis, 60)

	fmt.Println(contaDoDenis.ObterSaldo())

	contaDaLuisa := contas.ContaCorrente{}
	contaDaLuisa.Depositar(500)
	PagarBoleto(&contaDaLuisa, 1000)

	fmt.Println(contaDaLuisa.ObterSaldo())

}
