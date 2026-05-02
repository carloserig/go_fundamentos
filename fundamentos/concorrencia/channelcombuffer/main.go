package main

import (
	"fmt"
	"math/rand"
	"time"
)

func prepareOrder(orderID int, completedOrders chan string) {
	fmt.Printf("Preparando o pedido #%d... \n", orderID)

	// Simulação de tempo de preparação (1 a 3 seg)
	time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second)

	// enviando o pedido para o channel
	completedOrders <- fmt.Sprintf("Pedido #%d concluido!!", orderID)

}

func main() {
	rand.NewSource(time.Now().UnixNano())

	orderCount := 4
	completedOrders := make(chan string, orderCount)

	// Lança uma goroutine para cada pedido
	for i := 1; i <= orderCount; i++ {
		go prepareOrder(i, completedOrders)
	}

	for i := 1; i <= orderCount; i++ {
		fmt.Println(<-completedOrders)
	}

	//Fechando o canal após o uso
	close(completedOrders)

}
