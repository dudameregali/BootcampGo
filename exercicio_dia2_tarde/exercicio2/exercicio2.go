// Exercício 2 - Empréstimo

// Um banco deseja conceder empréstimos a seus clientes, mas nem todos podem acessá-los.
// Para isso, possui certas regras para saber a qual cliente pode ser concedido. Apenas
// concede empréstimos a clientes com mais de 22 anos, empregados e com mais de um ano
// de atividade. Dentro dos empréstimos que concede, não cobra juros para quem tem um
// salário superior a US$ 100.000.
// É necessário fazer uma aplicação que possua essas variáveis e que imprima uma mensagem
// de acordo com cada caso.
// Dica: seu código deve ser capaz de imprimir pelo menos 3 mensagens diferentes.

package main

import "fmt"

func main() {
	maria := map[string]int{"idade": 23, "tempo_atividade": 1, "salario": 100000}

	if maria["idade"] >= 22 && maria["tempo_atividade"] >= 1 {
		if maria["salario"] >= 100000 {
			fmt.Println("Empréstimo concedido sem juros!")
		} else {
			fmt.Println("Empréstimo concedido com juros!")
		}
	} else {
		fmt.Println("Empréstimo negado!")
	}
}
