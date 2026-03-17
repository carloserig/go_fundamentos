package main

func main() {
	/*println("Maps")
	meses := map[int]string{
		1:  "Janeiro",
		2:  "Fevereiro",
		3:  "Março",
		4:  "Abril",
		5:  "Maio",
		6:  "Junho",
		7:  "Julho",
		8:  "Agosto",
		9:  "Setembro",
		10: "Outubro",
		11: "Novembro",
	}
	meses[12] = "Dezembro"
	println(meses[4])
	println(meses[12])
	*/
	idades := map[string]int{
		"carlos":  30,
		"eduardo": 20,
		"helisa":  10,
	}
	println(idades["carlos"])
	println(idades["eduardo"])
	println(idades["helisa"])

	frutas := map[string]string{
		"banana": "amarela",
		"maça":   "vermelha",
		"mamao":  "laranja",
	}

	for chave, valor := range frutas {
		println(chave, valor)
	}
}
