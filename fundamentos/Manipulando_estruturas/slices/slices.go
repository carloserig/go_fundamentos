package main

import "fmt"

func main() {
	// slice não tem tamanho fixo, permitindo que você adicione
	// ou remova elementos dinamicamente
	// append -> adicionar ou concatenar

	// Inicializar slice(nil)
	var numeros []int
	// adicionar depois
	numeros = append(numeros, 50, 10, 99, 30)
	numeros = append(numeros, 1000)
	fmt.Println("Conteudo do slice numeros:", numeros)
	fmt.Println("Pegando posiçao 0 do slice:", numeros[0])

	// inicializando e adicionando valroes
	frutas := []string{"banana", "maça", "mamao"}
	fmt.Println("Slice de frutas:", frutas)
	frutas = append(frutas, "pêra")
	fmt.Println("Slice de frutas após o append:", frutas)

	// Inicializar slice vazio(not nil)
	nomes := make([]string, 0)
	fmt.Println("Slice nomes:", nomes)

	nomes = append(nomes, "carlos", "Eduardo", "Heloisa", "Antonio")
	fmt.Println("Slice de nomes:", nomes)

	//adicionar novo elemento posição x
	nomes = append(nomes[:1], append([]string{"Maria"}, nomes[1:]...)...)
	fmt.Println("Novo Slice de nomes:", nomes)

	fmt.Println()
	fmt.Println("Percorrendo slice com for")
	for n := 0; n < len(nomes); n++ {
		fmt.Println(nomes[n])
	}

	// remover primeiro
	nomes = nomes[1:]
	fmt.Println(nomes)

	// remover último -1
	nomes = nomes[:len(nomes)-1]
	fmt.Println(nomes)

	// remover segundo item
	nomes = append(nomes[:1], nomes[2:]...)
	fmt.Println(nomes)

}
