package repositories

type ContaRepositories interface {
	Sacar(valor float64) string
	VerificarConta(valor float64) string
	Saldo(saldo float64)
	Depositar(valorDoDeposito string) (string, float64)
}
