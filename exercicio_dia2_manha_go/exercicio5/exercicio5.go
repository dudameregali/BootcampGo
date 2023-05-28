// Exercício 5 - Cálculo da quantidade de alimento
// Um abrigo de animais precisa descobrir quanta comida comprar para os animais de
// estimação. No momento eles só têm tarântulas, hamsters, cachorros e gatos, mas a previsão
// é que haja muito mais animais para abrigar.
// 1. Cães precisam de 10 kg de alimento
// 2. Gatos 5 kg
// 3. Hamster 250 gramas.
// 4. Tarântula 150 gramas.

// Precisamos:
// 1. Implementar uma função Animal que receba como parâmetro um valor do tipo texto
// com o animal especificado e que retorne uma função com uma mensagem (caso não
// exista o animal)
// 2. Uma função para cada animal que calcule a quantidade de alimento com base na
// quantidade necessária do animal digitado.
// Exemplo:

// const (
// 	dog = "dog"
// 	cat = "cat"
// 	)

// 	...
// 	animalDog, msg := Animal(dog)
// 	animalCat, msg := Animal(cat)

// ...
// var amount float64
// amount+= animaldog(5)
// amount+= animalCat(8)
package main

import (
	"errors"
	"fmt"
)

func quantComida(animal string) (func(int) int, error) {
	switch animal {
	case "cao":
		return funcCao, nil
	case "gato":
		return funcGato, nil
	case "hamster":
		return funcHamster, nil
	case "tarantula":
		return funcTarantula, nil
	default:
		return nil, errors.New("função inválida")
	}
}

func funcCao(quant int) int {
	return quant * 10000
}

func funcGato(quant int) int {
	return quant * 5000
}

func funcHamster(quant int) int {
	return quant * 250
}

func funcTarantula(quant int) int {
	return quant * 150
}

func main() {

	animalDog, _ := quantComida("cao")
	animalCat, _ := quantComida("gato")

	var total int
	total += animalDog(5)
	total += animalCat(8)

	quilos := total / 1000

	fmt.Printf("Total de comida necessária são %dgramas \n", total)
	fmt.Printf("Total de comida necessária são %dkg \n", quilos)
}
