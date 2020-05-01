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

const qtdeMonitoramento = 5
const delay = 4
const statusOk = 200

func main() {
	exibirIntroducao()
	for {
		exibeMenu()
		comando := capturaComando()
		realizaComando(comando)
	}
}

func exibirIntroducao() {
	nome := "Gabriel"
	fmt.Println("Olá mundo", nome)
}

func exibeMenu() {
	fmt.Println("Selecione uma das opções")
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func capturaComando() int {
	var comando int
	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi", comando)
	return comando
}

func realizaComando(comando int) {
	switch comando {
	case 1:
		monitoraSites()
	case 2:
		imprimirConteudoArquivo("log.txt")
	case 0:
		fmt.Println("Saindo do sistema")
		os.Exit(0)
	default:
		fmt.Println("Comando não reconhecido pelo sistema")
		os.Exit(255)
	}
}

func monitoraSites() {
	fmt.Println("Iniciando monitoramento ...")
	sites := inicializaListaSites()
	for i := 0; i < 5; i++ {
		for i, site := range sites {
			verificaStatusSite(i, site)
		}
		fmt.Println("")
		time.After(delay * time.Second)
	}
	fmt.Println("")
}

func inicializaListaSites() []string {
	sites := lerLinhasArquivo("sites.txt")
	return sites
}

func lerLinhasArquivo(nome string) []string {
	linhas := []string{}

	arquivo, err := os.Open(nome)
	if err != nil {
		fmt.Println("Houve um erro na leitura do arquivo:", nome, "Erro:", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		linhas = append(linhas, linha)

		if err == io.EOF {
			break
		}
	}

	arquivo.Close()
	return linhas
}

func verificaStatusSite(index int, url string) {
	fmt.Println("Verificando site index:", index, ":", url)
	response, _ := http.Get(url)
	if response.StatusCode == statusOk {
		fmt.Println("Site:", url, "foi carregada com sucesso!")
		registraLogs(url, true, "log.txt")
	} else {
		fmt.Println("Site:", url, "esta com problemas, Status Code:", response.StatusCode)
		registraLogs(url, false, "log.txt")
	}
}

func registraLogs(url string, online bool, nomeArquivo string) {
	arquivo, err := os.OpenFile(nomeArquivo, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " O site : " + url + " Online: " + strconv.FormatBool(online) + "\n")
	arquivo.Close()
}

func imprimirConteudoArquivo(nome string) {
	fmt.Println("Exibindo", nome)
	arquivo, err := ioutil.ReadFile(nome)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(arquivo))
}
