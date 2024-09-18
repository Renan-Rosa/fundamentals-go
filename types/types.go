package main

import "fmt"

// Definindo um novo tipo
type Pessoa struct {
	Nome  string
	Idade int
}

// Função que recebe um tipo Pessoa
func exibirPessoa(p Pessoa) {
	fmt.Printf("Nome: %s, Idade: %d\n", p.Nome, p.Idade)
}

func main() {
	// Tipos primitivos

	// Booleano
	var verdade bool = true
	var falso bool = false
	fmt.Println("Booleano verdadeiro:", verdade)
	fmt.Println("Booleano falso:", falso)

	// Inteiros
	var inteiro int = -42
	var inteiroPos uint = 42 // Tipo uint (unsigned int)
	fmt.Println("Inteiro:", inteiro)
	fmt.Println("Inteiro positivo:", inteiroPos)

	// Números de ponto flutuante
	var float32Num float32 = 3.14
	var float64Num float64 = 2.71828
	fmt.Println("Número de ponto flutuante (float32):", float32Num)
	fmt.Println("Número de ponto flutuante (float64):", float64Num)

	// Strings
	var texto string = "Olá, Go!"
	fmt.Println("Texto:", texto)

	// Usando o tipo definido
	p := Pessoa{"Maria", 28}
	exibirPessoa(p)

	// Exemplos adicionais
	// Definindo uma variável de tipo int8
	var int8Var int8 = 127
	fmt.Println("Valor de int8Var:", int8Var)

	// Definindo uma variável de tipo uint64
	var uint64Var uint64 = 18446744073709551615
	fmt.Println("Valor de uint64Var:", uint64Var)

	// Definindo uma variável de tipo rune (equivalente a int32)
	var runeVar rune = 'A'
	fmt.Println("Valor de runeVar:", runeVar, "Caracter:", string(runeVar))
}
