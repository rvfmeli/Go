package contas

type verificarConta interface {
	VerificarConta(valor float64) string
}
