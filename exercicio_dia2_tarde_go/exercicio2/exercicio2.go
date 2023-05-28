// Exercício 2 - Produtos de e-commerce
// Diversas lojas de e-commerce precisam realizar funcionalidades no Go para gerenciar
// produtos e devolver o valor do preço total.
// As empresas têm 3 tipos de produtos:
// - Pequeno, Médio e Grande.
// Existem custos adicionais para manter o produto no armazém da loja e custos de envio.

// Lista de custos adicionais:
// - Pequeno: O custo do produto (sem custo adicional)
// - Médio: O custo do produto + 3% pela disponibilidade no estoque
// - Grande: O custo do produto + 6% pela disponibilidade no estoque + um custo
// adicional pelo envio de $2500.

// 2

// Requisitos:
// - Criar uma estrutura “loja” que guarde uma lista de produtos.
// - Criar uma estrutura “produto” que guarde o tipo de produto, nome e preço
// - Criar uma interface “Produto” que possua o método “CalcularCusto”

// - Criar uma interface “Ecommerce” que possua os métodos “Total” e “Adicionar”.
// - Será necessário uma função “novoProduto” que receba o tipo de produto, seu nome
// e preço, e devolva um Produto.
// - Será necessário uma função “novaLoja” que retorne um Ecommerce.

// - Interface Produto:
// - Deve possuir o método “CalcularCusto”, onde o mesmo deverá calcular o
// custo adicional segundo o tipo de produto.

// - Interface Ecommerce:
// - Deve possuir o método “Total”, onde o mesmo deverá retornar o preço total com
// base no custo total dos produtos + o adicional citado anteriormente (caso a categoria
// tenha)
// - Deve possuir o método “Adicionar”, onde o mesmo deve receber um novo produto
// e adicioná-lo a lista da loja

package main

import "fmt"

type loja struct {
	listaProdutos []produto
}

func (l *loja) Adicionar(prod produto) {
	l.listaProdutos = append(prod)
}

type Ecommerce interface {
}

func Total(e Ecommerce) float64 {
	total := 0.0
	for _, produto := range e.listaProdutos {
		total += produto.calcularCusto()
	}
	return total
}

type produto struct {
	nome  string
	tipo  string
	preco float64
}

type Produto interface {
}

func calcularCusto(p Produto) float64 {
	switch p.tipo {
	case "Pequeno":
		return p.preco
	case "Medio":
		return p.preco + (p.preco * 0.03)
	case "Grande":
		return p.preco + (p.preco * 0.06) + 2500.00
	default:
		return 0.0
	}
}

func main() {

	var produto1, produto2, produto3 Produto

	produto1 = produto{
		nome:  "copo",
		tipo:  "Pequeno",
		preco: 9.90,
	}
	produto2 = produto{
		nome:  "Televisão",
		tipo:  "Medio",
		preco: 2589.60,
	}
	produto3 = produto{
		nome:  "cama",
		tipo:  "Grande",
		preco: 658.50,
	}

	var lojaDoJoao Ecommerce
	lojaDoJoao = loja{
		listaProdutos: []produto{produto1, produto2, produto3},
	}

	fmt.Printf("Valor total do produto 1 R$%f", produto1.calcularCusto())

	fmt.Println(lojaDoJoao)

}
