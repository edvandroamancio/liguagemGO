// Meu Primeiro código em GO
package main

import "fmt"

func main() {
	fmt.Println("Olá, Mundo em Go!")
	entradadados()
}

func entradadados() {
	var nome string
	var nota, soma float64

	fmt.Println("Digite seu nome")
	fmt.Scan(&nome)
	fmt.Println("Vamos calcular sua média", nome)

	for i := 1; i <= 4; i++ {
		fmt.Println("Digite sua nota", i)
		fmt.Scan(&nota)
		soma = soma + nota
	}

	fmt.Println(nome, "sua média é", soma/4)

}
