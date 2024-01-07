package main

import "fmt"

func main() {
	//variaveis
	var nome string
	var idade int
	var altura float64

	fmt.Print("Digite seu nome: ")
	fmt.Scanln(&nome)

	fmt.Print("Digite sua idade: ")
	fmt.Scanln(&idade)

	fmt.Print("Digite sua altura (em metros): ")
	fmt.Scanln(&altura)

	fmt.Println("Inofrmações do Usuário:")
	fmt.Printf("Nome: %s\n", nome)
	fmt.Printf("Idade: %d\n", idade)
	fmt.Printf("Altura: %.2f metros\n", altura)
}
