package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {
	exibeIntroducao()
	//leSitesDoArquivo()
	registraLog("site-falso", false)
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
			imprimeLog()
		case 0:
			fmt.Println("Saindo do programa ...")
			os.Exit(0)
		default:
			fmt.Println("Não conheçe este comando")
			os.Exit(-1)
		}
	}

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

func exibeIntroducao() {
	var nome string = "Mikaela"
	var versao float32 = 1.1
	fmt.Println("Ola, Sra", nome)
	fmt.Println("Este programa esta na versao", versao)
}

func iniciarMonitoramento() {

	fmt.Println("Monitorando ...")

	sites := leSitesDoArquivo()

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
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro ao testar sites:", err)
	}
	if resp.StatusCode == 200 {
		fmt.Println("Site: ", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site: ", site, "esta com problemas. Status code: ", resp.StatusCode)
		registraLog(site, false)
	}
}

func leSitesDoArquivo() []string {
	var sites []string

	arquivo, err := os.Open("sites.txt")
	//arquivo, err := ioutil.ReadFile("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
		//return
	}
	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)
		fmt.Println(linha)

		if err == io.EOF {
			//fmt.Println("Ocorreu um erro:", err)
			break
		}

	}

	fmt.Println(sites)

	arquivo.Close()
	//fmt.Println(string(arquivo))
	return sites
}

func registraLog(site string, status bool) {
	//criando arquivo
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "- online: " + strconv.FormatBool(status) + "\n")
	//fmt.Println(arquivo)
	arquivo.Close()
}

func imprimeLog() {

	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))
}
