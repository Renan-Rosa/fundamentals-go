package variables

import "fmt"

func VariableFunction() (string, string, int, float64) {
	// Declaração de variáveis com tipo explícito
	var nome string = "John"
	var idade int = 30

	// Declaração de variáveis com inferência de tipo
	sobrenome := "Doe"
	altura := 1.80

	// Impressão das variáveis
	fmt.Println("Nome:", nome)
	fmt.Println("Sobrenome:", sobrenome)
	fmt.Println("Idade:", idade)
	fmt.Println("Altura:", altura)

	return nome, sobrenome, idade, altura
}
