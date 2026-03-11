package main

import "fmt"

func main() {
	fmt.Println("Manipulando Arrays")
	var numeros [4]int
	numeros[0] = 10
	numeros[1] = 20
	numeros[2] = 30
	numeros[3] = 40
	fmt.Println(numeros)

	fmt.Println()
	//var idades = [3]int{65, 34, 89} ou
	idades := [3]int{65, 34, 89}
	fmt.Println(idades)

	fmt.Println()
	//                    0          1          2
	nomes := [3]string{"CARLOS", "EDUARDO", "HELOISA"}
	fmt.Println(nomes)
	fmt.Println(nomes[1])

	fmt.Println()
	fmt.Println("Percorrendo array com for")
	for n := 0; n < len(nomes); n++ {
		fmt.Println(nomes[n])
	}

	fmt.Println()
	fmt.Println("Percorrendo array com range, somente valores")
	for i := range nomes {
		fmt.Println(nomes[i])
	}

	fmt.Println()
	fmt.Println("Percorrendo array com range, indice e valor")
	for i, nome := range nomes {
		fmt.Println(i, nome)
	}

	fmt.Println()
	fmt.Println("Percorrendo array com range, ocultar indice")
	for _, nome := range nomes {
		fmt.Println(nome)
	}

}
