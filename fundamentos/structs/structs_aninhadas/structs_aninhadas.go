package main

import "fmt"

type Endereco struct {
	Cidade string
	Estado string
}

type Pessoa struct {
	Nome     string
	Idade    int
	Endereco Endereco
}

func main() {
	fmt.Println("Structs aninhadas")
	p := Pessoa{
		Nome:  "João Paulo",
		Idade: 45,
		Endereco: Endereco{
			Cidade: "Curitiba",
			Estado: "PR",
		},
	}
	fmt.Println(p)
	fmt.Println(p.Endereco)
	fmt.Println(p.Endereco.Cidade)

	pessoas := map[int]Pessoa{
		1: {Nome: "Carlos Erig", Idade: 40, Endereco: Endereco{Cidade: "Marechal Cândido Rondon", Estado: "PR"}},
		2: p,
	}
	fmt.Println(pessoas)

	//percorrendo com for
	for id, pessoa := range pessoas {
		fmt.Println(id, pessoa.Nome, pessoa.Idade, pessoa.Endereco.Cidade)
	}
}
