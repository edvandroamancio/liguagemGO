// Vetor e Matriz
package main

import "fmt"

func main() {

	//criando vetor
	var vetor [3]int
	var matriz [3][3]int

	for i := 0; i < 3; i++ {
		vetor[i] = i
	}
	fmt.Println(vetor)

	//imprimindo matriz
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			matriz[i][j] = j
		}
	}

	fmt.Println(matriz)

	//proxima etapa ler do teclado
	//vetor

	for i := 0; i < 3; i++ {
		fmt.Println("Digite o", i+1, "numero")
		fmt.Scan(&vetor[i])
	}
	fmt.Println(vetor)

}
