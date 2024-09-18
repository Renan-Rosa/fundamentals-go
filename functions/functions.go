package main

import (
	"fmt"
)

// Função simples: Soma dois números
func somar(a, b int) int {
	return a + b
}

// Função com múltiplos retornos: Divide dois números e retorna o quociente e um erro
func dividir(a, b int) (int, error) {
	if b == 0 {
		// Retorna um erro se o divisor for zero
		return 0, fmt.Errorf("divisão por zero")
	}
	// Retorna o quociente e um valor nil para o erro
	return a / b, nil
}

// Função de ordem superior que retorna uma função: Cria uma função que soma um valor fixo
func highOrderExample(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}

// Função de ordem superior que recebe uma função como parâmetro: Aplica a função fornecida a um valor
func applyFunction(fn func(int) int, valor int) int {
	return fn(valor)
}

func main() {
	// Chamando a função somar
	resultado := somar(10, 5)
	fmt.Println("Resultado da soma:", resultado)

	// Chamando a função dividir com um divisor válido
	quociente, err := dividir(10, 2)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Resultado da divisão:", quociente)
	}

	// Chamando a função dividir com divisor zero para demonstrar erro
	quociente, err = dividir(10, 0)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Resultado da divisão:", quociente)
	}

	// Usando a função highOrderExample para criar uma função que soma 10
	somarCom10 := highOrderExample(10)

	// Aplicando a função criada
	resultadoHighOrder := somarCom10(5)
	fmt.Println("Resultado da soma com High Order Function (10 + 5):", resultadoHighOrder)

	// Usando a função applyFunction para aplicar a função de soma com 10
	resultadoAplicado := applyFunction(somarCom10, 8)
	fmt.Println("Resultado da aplicação da função High Order (10 + 8):", resultadoAplicado)
}
