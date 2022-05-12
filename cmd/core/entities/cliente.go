package clientes

type Titular struct {
	Nome, CPF, Profissao string
}

type ContaCorrente struct {
	titular         Titular
	numeroDaAgencia int64
	numeroDaConta   int64
	saldo           float64
}
