package main

import "fmt"

func main() {
	num1 := 20
	num2 := 25
	operacao := "soma"
	var resultado int

	if operacao == "soma" {
		resultado = num1 + num2
		fmt.Printf("A soma de %d + %d é: %d\n", num1, num2, resultado)
	} else if operacao == "subtracao" {
		resultado = num1 - num2
		fmt.Printf("A subtração de %d - %d é: %d\n", num1, num2, resultado)
	} else if operacao == "divisao" {
		if num2 != 0 {
			resultado = num1 / num2
			fmt.Printf("A divisao de %d / %d é: %d\n", num1, num2, resultado)
		} else {
			fmt.Println("Erro: Divisão por zero não é permitida")
		}

	} else if operacao == "multiplicacao" {
		resultado = num1 * num2
		fmt.Printf("A multiplicacao de %d * %d é: %d\n", num1, num2, resultado)
	} else {
		fmt.Println("Operação Inválida. Por favor escolha entre 'soma', 'subtracao', 'divisao' e 'multiplicacao'.")
	}

}
