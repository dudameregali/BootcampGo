// Exercício 1 - Letras de uma palavra

// A Real Academia Brasileira quer saber quantas letras tem uma palavra e então ter cada uma
// das letras separadamente para soletrá-la. Para isso terão que:
// 1. Crie uma aplicação que tenha uma variável com a palavra e imprima o número de
// letras que ela contém.
// 2. Em seguida, imprimi cada uma das letras.

package main

import "fmt"

func main() {
	palavra := "Golang" // Palavra exemplo

	// Imprime o número de letras da palavra
	numLetras := len(palavra)
	fmt.Printf("A palavra tem %d letras.\n", numLetras)

	// Imprime cada uma das letras da palavra separadamente
	fmt.Println("As letras da palavra são:")
	for _, letra := range palavra {
		fmt.Printf("%c\n", letra)
	}
}
