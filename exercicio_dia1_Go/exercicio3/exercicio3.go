// Exercício 3 - Declaração de variáveis

// Um professor de programação está corrigindo as avaliações de seus alunos na disciplina de
// Programação I para fornecer-lhes o feedback correspondente. Um dos pontos do exame é
// declarar diferentes variáveis.
// Ajude o professor com as seguintes questões:
// 1. Verifique quais dessas variáveis declaradas pelo aluno estão corretas.
// 2. Corrigir as incorrectas.

// var 1nome string - errada
// var sobrenome string - certa
// var int idade - errado
// 1sobrenome := 6 - errado
// var licenca_para_dirigir = true - certo
// var estatura da pessoa int - errado
// quantidadeDeFilhos := 2 - certo

package main

import "fmt"

var nome1 string
var sobrenome string
var idade int
var licenca_para_dirigir = true
var estatura_da_pessoa int

func main() {
	sobrenome1 := 6
	quantidadeDeFilhos := 2

	fmt.Println(nome1, sobrenome, idade, licenca_para_dirigir, estatura_da_pessoa, sobrenome1, quantidadeDeFilhos)
}
