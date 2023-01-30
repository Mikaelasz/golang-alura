package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {
	exibeIntroducao()

	for {
		exibeMenu()

		comando := leComando()

		if comando == 1 {
			fmt.Printf("Monitorando ...")
		} else if comando == 2 {
			fmt.Println("Exibindo log ...")
		} else if comando == 0 {
			fmt.Println("Saindo do programa ...")
		} else {
			fmt.Printf("Não conheçe este comando")
		}

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo log ...")
		case 0:
			fmt.Println("Saindo do programa ...")
			os.Exit(0)
		default:
			fmt.Println("Não conheçe este comando")
			os.Exit(-1)
		}
	}

}

func exibeIntroducao() {
	var nome string = "Mikaela"
	var versao float32 = 1.1
	fmt.Println("Ola, Sra", nome)
	fmt.Println("Este programa esta na versao", versao)
}
func exibeMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir log")
	fmt.Println("0 - Sair do programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", &comandoLido)

	return comandoLido
}
func iniciarMonitoramento() {

	fmt.Println("Monitorando ...")
	//declarando slice
	sites := []string{"https://www.alura.com.br", "https://www.google.com", "https://www.gmail.com"}

	//fmt.Println(sites)
	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando SITE", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")

	}

	fmt.Println("")
}

func testaSite(site string) {
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site: ", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site: ", site, "esta com problemas. Status code: ", resp.StatusCode)
	}
}

/*func exibeNomes() {
	nomes := []string{"Mikaela", "Zéca", "Murilo"}
	nomes = append(nomes, "Joao")
	fmt.Println(nomes)
	fmt.Println(reflect.TypeOf(nomes))
	fmt.Println("O meu slice tem", len(nomes))
}*/
