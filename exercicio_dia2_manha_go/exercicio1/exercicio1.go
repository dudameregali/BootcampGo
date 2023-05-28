// Exercício 1 - Impostos de salário

// Uma empresa de chocolates precisa calcular o imposto de seus funcionários no momento de
// depositar o salário, para cumprir seu objetivo será necessário criar uma função que retorne o
// imposto de um salário.
// Temos a informação que um dos funcionários ganha um salário de R$50.000 e será
// descontado 17% do salário. Um outro funcionário ganha um salário de $150.000, e nesse
// caso será descontado mais 10%.

package main

import (
	"fmt"
)

func calImposto(salario, imposto float64) float64 {
	return salario - (salario * (imposto / 100))
}

func main() {
	fmt.Println("Salário líquido: R$ ", calImposto(50000.00, 17))
	fmt.Println("Salário líquido: R$ ", calImposto(150000.00, 10))
}
