package main

import "fmt"

type Animal struct {
	Nome string
	Tipo string
}

func main() {
	fmt.Println("Struct Animais")
	a1 := Animal{Nome: "Rex", Tipo: "Cachorro"}
	a2 := Animal{Nome: "Bella", Tipo: "Gato"}
	a3 := Animal{Nome: "Luna", Tipo: "Cavalo"}
	a4 := Animal{Nome: "Loro", Tipo: "Papagaio"}

	animais := make(map[int]Animal)
	animais[1] = a1
	animais[2] = a2
	animais[3] = a3
	animais[4] = a4

	fmt.Println(animais)
}
