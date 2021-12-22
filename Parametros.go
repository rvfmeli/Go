package main

import "fmt"


// explicando parametro
func semParametro() string {
	return "exemplo de funcao sem paramentro!"
}

func umParametro(texto string) string {
	return texto
}

func doisParametros (texto string, numero int) (string, int) {
	return texto, numero
}

func parametro ()  {
	fmt.Println(semParametro())
	fmt.Println(umParametro("um parametro"))
	fmt.Println(doisParametros("dois parametros: eessa string e o numero", 10))

}
