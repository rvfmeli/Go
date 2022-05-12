package clientes

type Titular struct {
	Nome      string
	cpf       string
	profissao string
}

type ContaCorrente struct {
	titular         Titular
	NumeroDaAgencia int64
	NumeroDaConta   int64
	Saldo           float64
}

type ContaPoupanca struct {
	Titular       Titular
	NumeroAgencia int64
	NumeroDaConta int64
	operacao      int64
	saldo         float64
}
