/* Scrivere un programma maxQuadrato.go che legge da standard input un intero n positivo e stampa,
utilizzando solo variabili di tipo int, il massimo quadrato (k^2) <= n.
Per calcolare il quadrato di un numero n usate n*n.*/

package main

import "fmt"

func main() {
	var num int

	fmt.Print("Inserire un numero: ")
	fmt.Scan(&num)

	maxQuadrato := 0
	for k := 0; k*k <= num; k++ {
		maxQuadrato = k * k
	}
	fmt.Print("Il massimo quadrato minore di ", num, " è: ", maxQuadrato)
}

// necessario un confronto tra l'ultimo numero uscito e quello appena generato, dato che va sempre avanti di 1 extra
