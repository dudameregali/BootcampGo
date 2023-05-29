// Exercício 2 - Ler arquivo
// A mesma empresa precisa ler o arquivo armazenado, para isso exige que:
// - Seja impresso na tela os valores tabelados, com título ( à esquerda para o ID e à
// direita para o Preço e Quantidade), o preço, a quantidade e abaixo do preço o total
// deve ser exibido (somando preço por quantidade)

// Exemplo:
// ID        Preco        Quantidade
// 111223    30012.00     1
// 444321    1000000.00   4
// 434321    50.50        1

// 4030062.50

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	data, err := os.ReadFile("../exercicio1/estoque.txt")
	resposta := "ID\tPreço\tQuantidade\t\n"
	elemento := ""
	// cont := 0
	for _, caracter := range data {
		if string(caracter) == ";" {
			resposta = fmt.Sprintf("%s%s\t", resposta, elemento)
			elemento = ""
		} else {
			elemento = fmt.Sprintf("%s%s", elemento, string(caracter))
		}
	}

	if err == nil {
		fmt.Printf("%s", resposta)
	} else {
		log.Print("erro ao gerar o arquivo", err)
	}
}
