package contas

import "golang/Go/cmd/core/repositories"

type sacarUseCase interface {
	Sacar(valor float64) string
}

type sacarUseCaseImpl struct {
	contaRepositories repositories.ContaRepositories
}

func (useCase *sacarUseCaseImpl) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		ContaCorrente.ObterSaldo()
		return "Saque realizado com sucesso"
	} else {
		return "saldo insuficiente"
	}
}
