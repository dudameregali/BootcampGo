// Exercício 2 - Calcular média
// Um colégio precisa calcular a média das notas (por aluno). Precisamos criar uma função na
// qual possamos passar N quantidade de números inteiros e devolva a média.
// Obs: Caso um dos números digitados seja negativo, a aplicação deve retornar um erro.

package main

import (
	"errors"
	"fmt"
)

func calMedia(notas ...float64) (float64, error) {
	soma := 0.0
	cont := 0.0
	for _, nota := range notas {
		if nota < 0 {
			return 0, errors.New("Uma das notas possui valor negativo, não sendo possível calcular a média")
		} else {
			soma += nota
			cont++
		}

	}
	return soma / cont, nil
}

func main() {
	res, err := calMedia(10, 9, 8, 10, 7)

	if err == nil {
		fmt.Println("A média do aluno é: ", res)
	} else {
		fmt.Println("Erro ao calcular a média - ", err.Error())
	}

}
