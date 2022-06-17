package navigate

import (
	"fmt"
	"net/http"
)

func IniciarMonitoramento() {
	site := "https://www.mercadolivre.com.br/"
	response, _ := http.Get(site)
	fmt.Println(response)

	if response.StatusCode == 200 {
		fmt.Println("monitorando...")
	}
}
