// Exercício 1 - Guardar arquivo
// Uma empresa que vende produtos de limpeza precisa de:
// 1. Implementar uma funcionalidade para guardar um arquivo de texto, com a informação
// de produtos comprados, separados por ponto e vírgula (csv).
// 2. Deve possuir o ID do produto, preço e a quantidade.
// 3. Estes valores podem ser hardcodeados ou escritos manualmente em uma variável.

package main

import (
	"fmt"
	"log"
	"os"
)

var (
	id    int
	preco float64
	quant int
)

func main() {

	fmt.Print("ID do Produto: ")
	fmt.Scan(&id)
	fmt.Print("Valor do Produto: ")
	fmt.Scan(&preco)
	fmt.Print("Quantidade Comprada: ")
	fmt.Scan(&quant)

	produtoAdd := fmt.Sprintf("%v;%v;%v;\n", id, preco, quant)

	err := os.WriteFile("./estoque.csv", []byte(produtoAdd), 0644)

	if err == nil {
		fmt.Print("Arquivo gerado!")
	} else {
		log.Print("erro ao gerar o arquivo", err)
	}

}
