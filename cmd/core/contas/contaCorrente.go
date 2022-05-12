package contas

//func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
//	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
//	if podeSacar {
//		c.saldo -= valorDoSaque
//		return "Saque realizado com sucesso"
//	} else {
//		return "saldo insuficiente"
//	}
//}

func (useCase *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		saldo.
		return "Deposito realizado com sucesso", c.saldo
	} else {
		return "Valor do deposito menor que zero", c.saldo
	}
}

func (useCase *ContaCorrente) Tranferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia < c.saldo && valorDaTransferencia > 0 {
		useCase.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}
}

func (useCase *ContaCorrente) ObterSaldo() float64 {
	return useCase.saldo
}
