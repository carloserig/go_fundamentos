package main

import "fmt"

// * = pegar o valor da variável da memoria
// & = posição na memória
func f(x *int) {
	*x = 20
	fmt.Println("Print na func f", *x)
}

func main() {
	a := 10
	f(&a) // ✔️ passa o endereço
	fmt.Println("Print na func main", a)
}
