# Fundamentos de Go

## Índice

- [Introdução](#introdução)
- [Configuração do Ambiente](#configuração-do-ambiente)
- [Pacotes e Importações](#pacotes-e-importações)
- [Função `main`](#função-main)
- [Tipos de Dados](#tipos-de-dados)
- [Variáveis](#variáveis)
- [Constantes](#constantes)
- [Funções](#funções)
  - [Funções Simples](#funções-simples)
  - [Funções Higher Order](#funções-higher-order)
  - [Funções Variáticas](#funções-variáticas)
- [Estruturas de Controle](#estruturas-de-controle)
  - [Loops](#loops)
  - [Condicionais](#condicionais)
- [Nomes Públicos e Privados](#nomes-públicos-e-privados)
- [Pacotes Internos](#pacotes-internos)

## Introdução

Go, também conhecido como Golang, é uma linguagem de programação desenvolvida pela Google que é conhecida por sua simplicidade e eficiência. Ela é frequentemente usada para desenvolvimento de software em sistemas distribuídos e aplicações de backend.

## Configuração do Ambiente

Antes de começar a programar em Go, você precisa configurar o ambiente de desenvolvimento:

1. **Verifique a Instalação do Go**:

    ```sh
    go version
    ```

2. **Inicialize um Novo Módulo**:

    ```sh
    go mod init <nome_do_projeto>
    ```

    Isso criará um arquivo `go.mod` com as especificações do projeto.

## Pacotes e Importações

Todo arquivo Go deve declarar um pacote. Se o pacote é `main`, o arquivo contém a função `main` que é o ponto de entrada da aplicação.

### Exemplo:

```go
package main

import "fmt"
```

### Importação de Pacotes

Os pacotes são importados com o comando `import` e devem ser usados no código. É possível importar múltiplos pacotes em uma única declaração.

### Exemplo:

```go
import (
    "fmt"
    "os"
)
```

## Função `main`

A função `main` é o ponto de entrada da aplicação Go. Todo programa Go executável deve ter uma função `main`.

### Exemplo:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello World!")
}
```

## Tipos de Dados

Go possui vários tipos de dados embutidos, incluindo:

- **Inteiros**: `int`, `int8`, `int16`, `int32`, `int64`
- **Pontos Flutuantes**: `float32`, `float64`
- **Booleanos**: `bool`
- **Strings**: `string`
- **Unsigned Integers**: `uint`, `uint8`, `uint16`, `uint32`, `uint64`
- **Complexos**: `complex64`, `complex128`

### Exemplo:

```go
var a int = 10
var b float64 = 3.14
var c bool = true
var d string = "Hello"
var e uint = 42
```

## Variáveis

Você pode declarar variáveis de várias formas:

1. **Usando `var`**:

    ```go
    var nome string = "Pedro"
    var idade int = 30
    ```

2. **Usando `:=`** (Declaração curta):

    ```go
    nome := "Pedro"
    idade := 30
    ```

**Nota**: Na declaração curta `:=`, o tipo é inferido e não pode ser definido explicitamente.

## Constantes

Constantes são valores que não podem ser alterados após a inicialização.

### Exemplo:

```go
const pi = 3.14
const greeting = "Hello, Go!"
```

## Funções

As funções são blocos de código que realizam uma tarefa específica. Elas podem receber parâmetros e retornar valores.

### Funções Simples

```go
func somar(a int, b int) int {
    return a + b
}
```

### Funções Higher Order

Funções que retornam outras funções.

```go
func somarHigher(a int) func(int) int {
    return func(b int) int {
        return a + b
    }
}

func main() {
    x := somarHigher(1)(3)
    fmt.Println(x)
}
```

### Funções Variáticas

Funções que aceitam um número variável de argumentos.

```go
func somar(nums ...int) int {
    var out int
    for _, n := range nums {
        out += n
    }
    return out
}

func main() {
    fmt.Println(somar(1, 3, 4))
}
```

## Estruturas de Controle

### Loops

Existem três formas de loops em Go: `for` básico, `for` com condição e `for` infinito.

#### Exemplo:

```go
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
```

### Condicionais

Condicionais permitem executar diferentes blocos de código com base em condições.

#### Exemplo:

```go
a := 10
if a > 5 {
    fmt.Println("a é maior que 5")
} else if a == 5 {
    fmt.Println("a é igual a 5")
} else {
    fmt.Println("a é menor que 5")
}
```

## Nomes Públicos e Privados

A visibilidade de funções e variáveis é determinada pela primeira letra do nome:

- **Público**: Começa com letra maiúscula.
- **Privado**: Começa com letra minúscula.

### Exemplo:

```go
func Foo() {} // Público
func foo() {} // Privado

var Foo string = "Olá" // Público
var foo string = "Olá" // Privado
```

## Pacotes Internos

Pacotes na pasta `internal` são acessíveis apenas dentro do mesmo projeto. Eles não podem ser importados de fora.

### Estrutura de Pacotes Internos

```plaintext
pacote/
  └── internal/
      └── pacote_2/
```

Dentro do pacote `pacote`, você pode acessar o pacote `pacote_2`, mas não de fora.

---

Este arquivo Markdown serve como uma introdução abrangente aos fundamentos de Go. Certifique-se de adicionar exemplos práticos e detalhados para cada conceito para facilitar a compreensão.