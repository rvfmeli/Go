package repositories

type ContaRepositories interface {
	Sacar(valor float64) string
	VerificarConta(valor float64) string
	Saldo(float642 float64)
}
