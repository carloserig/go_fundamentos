package utils

import "fmt"

// Função com primeira letra Maiúscula (pública)
func Mensagem(nome string) {
	fmt.Println("Olá, seja bem vindo,", nome)
}

func StatusSistema() string {
	return "ONLINE"
}

func Somar(n1 float64, n2 float64) float64 {
	return n1 + n2
}
