package contas

import (
	"fmt"
	"golang/Go/cmd/core/repositories"
)

type verificarContaUseCase interface {
	Sacar(valor float64) string
}

type verificarContaUseCaseImpl struct {
	contaRepositories repositories.ContaRepositories
}

func (useCase *verificarContaUseCaseImpl) VerificarConta(saldo float64) {

	useCase.contaRepositories.Saldo(saldo)

	}
}
