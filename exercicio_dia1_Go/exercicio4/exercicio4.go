// Um estudante de programação tentou fazer declarações de variáveis com seus respectivos
// tipos em Go mas teve várias dúvidas ao fazê-lo. A partir disso, ele nos deu seu código e
// pediu a ajuda de um desenvolvedor experiente que pode:
// ● Revisar o código e realizar as devidas correções.

// var sobrenome string = "Silva"
// var idade int = "25"
// boolean := "false";
// var salario string = 4585.90
// var nome string = "Fellipe"

package main

import "fmt"

var sobrenome string = "Silva"
var idade int = 25
var salario float32 = 4585.90
var nome string = "Fellipe"

func main() {
	varboolean := false

	fmt.Println(sobrenome, idade, salario, nome, varboolean)
}
