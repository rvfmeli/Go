package navigate

import (
	"fmt"
	"net/http"
	"reflect"
	"time"
)

func IniciarMonitoramento() {

	const monitoramentos = 3
	const delay = 5

	sites := []string{"https://www.mercadolivre.com.br/"}

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			response, _ := http.Get(site)
			fmt.Println(response)

			if response.StatusCode == 200 {
				fmt.Println("monitorando...")
				fmt.Println("testando sites", i, ":", site)

			}
			time.Sleep(delay * time.Second)
			fmt.Println("")
		}
		fmt.Println("")
	}
}

func ExibirNomes() {
	nomes := []string{"rafa", "july", "kapih"}
	fmt.Println(nomes)
	fmt.Println(reflect.TypeOf(nomes))
	fmt.Println("meu slice tem ", len(nomes), "nomes")
	fmt.Println("tem capacidade de: ", cap(nomes))
	fmt.Println("")
}
