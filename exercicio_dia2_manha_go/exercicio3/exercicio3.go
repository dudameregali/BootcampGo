// Exercício 3 - Calcular salário

// Uma empresa marítima precisa calcular o salário de seus funcionários com base no número
// de horas trabalhadas por mês e na categoria do funcionário.
// Se a categoria for C, seu salário é de R$1.000 por hora
// Se a categoria for B, seu salário é de $1.500 por hora mais 20% caso tenha passado de 160h
// mensais
// Se a categoria for A, seu salário é de $3.000 por hora mais 50% caso tenha passado de 160h
// mensais

// Calcule o salário dos funcionários conforme as informações abaixo:
// - Funcionário de categoria C: 162h mensais
// - Funcionário de categoria B: 176h mensais
// - Funcionário de categoria A: 172h mensais

package main

import "fmt"

func calSalario(categoria string, horas float64) float64 {
	switch categoria {
	case "a":
		if horas > 160 {
			return (horas * 3000) + (horas*3000)*0.5
		} else {
			return horas * 3000
		}
	case "b":
		if horas > 160 {
			return (horas * 1500) + (horas*3000)*0.2
		} else {
			return horas * 1500
		}
	case "c":
		return horas * 1000
	default:
		return 0.0
	}
}

func main() {

	fmt.Println("Salario do Funcionário 1: ", calSalario("a", 170.0))
	fmt.Println("Salario do Funcionário 2: ", calSalario("b", 170.0))
	fmt.Println("Salario do Funcionário 3: ", calSalario("c", 170.0))

}
