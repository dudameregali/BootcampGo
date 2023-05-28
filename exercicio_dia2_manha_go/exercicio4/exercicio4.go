// Os professores de uma universidade na Colômbia precisam calcular algumas estatísticas de
// notas dos alunos de um curso, sendo necessário calcular os valores mínimo, máximo e médio
// de suas notas.
// Será necessário criar uma função que indique que tipo de cálculo deve ser realizado (mínimo,
// máximo ou média) e que retorna outra função (e uma mensagem caso o cálculo não esteja
// definido) que pode ser passado uma quantidade N de inteiros e retorne o cálculo que foi
// indicado na função anterior
// Exemplo:

// const (
// 	minimum = "minimum"
// 	average = "average"
// 	maximum = "maximum"
// 	)

// 	...
// 	minhaFunc, err := operation(minimum)
// 	averageFunc, err := operation(average)
// 	maxFunc, err := operation(maximum)

// 	...
// 	minValue := minhaFunc(2, 3, 3, 4, 10, 2, 4, 5)
// 	averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
// 	maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

package main

import (
	"errors"
	"fmt"
)

// const (
// 	minimo = "minimo"
// 	maximo = "maximo"
// 	media = "media"
// )

func calcula(operacao string) (func(valores ...float64) float64, error) {
	switch operacao {
	case "minimo":
		return minimo, nil
	case "maximo":
		return maximo, nil
	case "media":
		return media, nil
	default:
		return nil, errors.New("Função inválida")
	}
}

func minimo(valores ...float64) float64 {
	min := 0.0
	for i, nota := range valores {
		if i == 0 {
			min = nota
		} else if min > nota {
			min = nota
		}
	}
	return min
}

func maximo(valores ...float64) float64 {
	max := 0.0
	for i, nota := range valores {
		if i == 0 {
			max = nota
		} else if max < nota {
			max = nota
		}
	}
	return max
}

func media(valores ...float64) float64 {
	sun := 0.0
	cont := 0.0
	for _, nota := range valores {
		sun += nota
		cont++
	}
	return sun / cont
}

func main() {

	minhaFunc, _ := calcula("minimo")
	averageFunc, _ := calcula("media")
	maxFunc, _ := calcula("maximo")

	minValue := minhaFunc(2, 3, 3, 4, 10, 2, 4, 5)
	averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	fmt.Println("Mínimo: ", minValue)
	fmt.Println("Média: ", averageValue)
	fmt.Println("Máximo: ", maxValue)
}
