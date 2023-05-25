// Exercício 1 - Imprimindo o nome na tela
// 1. Crie uma aplicação que tenha uma variável “nome” e outra “idade”.
// 2. Imprima no terminal o valor de cada variável.

package main

import "fmt"

var name string
var age int

func main() {
	name = "Maria Eduarda"
	age = 22

	fmt.Println(name, age)
}
