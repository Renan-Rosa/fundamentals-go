package main

import "fmt"

func main() {
	// Loop for básico
	for i := 1; i <= 5; i++ {
		fmt.Println("Contagem:", i)
	}

	// Loop com condição
	x := 10
	for x > 0 {
		fmt.Println("Valor de x:", x)
		x--
	}

	// Loop infinito
	// for {
	//     fmt.Println("Loop infinito!")
	// }
}
