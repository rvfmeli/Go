package contas

import (
	clientes "golang/Go/cmd/core/entities"
	"golang/Go/cmd/core/repositories"
)

type SacarUseCase interface {
	Sacar(valor float64) string
}

type SacarUseCaseImpl struct {
	contaRepositories repositories.ContaRepositories
}

func (useCase *SacarUseCaseImpl) Sacar(valorDoSaque float64) string {

	var saldo clientes.ContaCorrente

	var saldoNaConta float64

	saldo.Saldo = saldoNaConta

	podeSacar := valorDoSaque > 0 && valorDoSaque <= saldoNaConta
	if podeSacar {
		saldoNaConta -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "saldo insuficiente"
	}
}
