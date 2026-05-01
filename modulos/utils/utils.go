package utils

import (
	"fmt"
)

func Mensagem(nome string) {
	fmt.Println("Olá seja bem vindo", nome)
}

func StatusSistema() string {
	return "ONLINE"
}
