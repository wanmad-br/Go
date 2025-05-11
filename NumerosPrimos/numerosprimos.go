package main

import (
	"fmt" // formata entrada e saida de dados
	"os"  // recursos do S.O.
	"strconv" // conversao de strings
)

//recebe um inteiro e devolve boleano
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		} 
	}
	return true
}

//sempre tem que ter um main para ser um executavel
func main() {
	var input string
	fmt.Print("Entre com um numero: ")
	fmt.Scanln(&input) //le o que foi digitado

	num, err := strconv.Atoi(input) //Atoi = ASCII To Integer
	if err != nil {
		fmt.Println("Numero invalido.")
		os.Exit(1) //sai do programa
	}

	if isPrime(num) {
		fmt.Printf("%d eh primo.\n", num)
	} else {
		fmt.Printf("%d NAO eh primo.\n", num)
	}
}
