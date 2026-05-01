package main

import "fmt"

type Pagamento interface {
	Pagar(valor float64)
}

type Cartao struct{}

func (c Cartao) Pagar(valor float64) {
	fmt.Println("Pagando com cartão:", valor)
}

type Pix struct{}

func (p Pix) Pagar(valor float64) {
	fmt.Println("Pagando com pix:", valor)
}

func ProcessarPagamento(p Pagamento, valor float64) {
	p.Pagar(valor)
}

func main() {
	ProcessarPagamento(Cartao{}, 100)
	ProcessarPagamento(Pix{}, 50)
	ProcessarPagamento(Pix{}, 500.50)
}
