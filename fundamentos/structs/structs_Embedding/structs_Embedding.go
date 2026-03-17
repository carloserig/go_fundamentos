package main

import "fmt"

type Endereco struct {
	Cidade string
	Estado string
}

type Pessoa struct {
	Nome  string
	Idade int
	Endereco
}

func main() {
	fmt.Println("Structs Embedding(Acesso direto)")
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
	fmt.Println(p.Cidade)

}
