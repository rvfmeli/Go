package contas

import (
	clientes "golang/Go/cmd/core/entities"
	"golang/Go/cmd/core/repositories"
)

type DepositarUseCase interface {
	Execute(valorDoDeposito float64) (string, float64)
}

type DepositarUseCaseImpl struct {
	repositories.ContaRepositories
}

func (useCase *DepositarUseCaseImpl) Depositar(valorDoDeposito float64) (string, float64) {
	var saldoConta clientes.ContaCorrente
	saldo := saldoConta.Saldo

	if valorDoDeposito > 0 {

		return "Deposito realizado com sucesso", saldo
	} else {
		return "Valor do Deposito menor que zero", saldo
	}
}
