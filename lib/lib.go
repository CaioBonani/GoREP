package lib

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var Substring string
var nome_arquivo string

func Run() (string) {
	Substring = os.Args[1]
	nome_arquivo = os.Args[2]
	caminho := "/home/bonani/GoREP/" + nome_arquivo

	conteudo, err := ioutil.ReadFile(caminho)
	if err != nil {
		panic(err)
	}

	var Texto = string(conteudo)
	// fmt.Printf("tipo do Texto: %T", Texto)

	return Texto
}

func Search(texto string, Substring string) {

	linhas := strings.Split(texto, "\n")

	for _, linha := range linhas {
		if strings.Contains(linha, Substring) {
			fmt.Println(linha)
		}
	}

}
