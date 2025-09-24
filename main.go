package main

import "fmt"

func main() {

	nome := "Mundo"
	ano := 2025

	fmt.Printf("Olá, %s! Estamos no ano de %d! \n", nome, ano)

	fds := false

	if fds {
		fmt.Println("Opa, chegou o fim de semana! Hora de descansar.")
	} else {
		fmt.Println("Foco! Ainda é dia de semana.")
	}
	dizOla()
}

func dizOla() {
	fmt.Println("Olá de novo, Renan!")
}
