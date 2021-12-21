package estudandoGo

import (
	"fmt"
)

func semParametro() string {
	return "exempl de funcao sem paramentro!"
}

func umParametro(texto string) string {
	return texto
}

func doisParametros (texto string, numero int) (string, int) {
	return texto, numero
}

func parametros()  {
	fmt.Println(semParametro())
	fmt.Println(umParametro("um parametro"))
	fmt.Println(doisParametros("dois parametros: eessa string e o numero", 10))


}