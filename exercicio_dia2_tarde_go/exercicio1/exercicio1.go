// Exercício 1 - Registro de estudantes

// Uma universidade precisa cadastrar os alunos e gerar uma funcionalidade para imprimir os
// detalhes dos dados de cada um deles, conforme o exemplo abaixo:
// Nome: [Nome do aluno]
// Sobrenome: [Sobrenome do aluno]
// RG: [RG do aluno]
// Data de admissão: [Data de admissão do aluno]

// Os valores que estão entre parênteses devem ser substituídos pelos dados fornecidos pelos
// alunos.
// Para isso é necessário gerar uma estrutura Alunos com as variáveis Nome, Sobrenome, RG,
// Data e que tenha um método de detalhameto.

package main

import (
	"fmt"
)

type Aluno struct {
	nome         string
	sobrenome    string
	rg           string
	dataAdmissao string
}

func (a Aluno) detalhe() {
	fmt.Printf("Nome:  \t\t%s\nSobrenome:  \t%s\nRG:  \t\t%s\nData de Admissão: %s\n", a.nome, a.sobrenome, a.rg, a.dataAdmissao)
}

func main() {
	maria := Aluno{
		nome:         "Maria",
		sobrenome:    "Silva",
		rg:           "00000000",
		dataAdmissao: "26//05//2023",
	}
	maria.detalhe()
}
