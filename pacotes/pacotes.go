package main

import (
	"fmt"
	"math" // Importando um pacote da standard library
)

func main() {
	// Usando uma função do pacote math
	raizQuadrada := math.Sqrt(16)
	fmt.Println("Raiz quadrada de 16:", raizQuadrada)

	// Usando o pacote fmt
	fmt.Printf("O valor de Pi é %.2f\n", math.Pi)
}
