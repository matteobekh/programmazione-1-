/*
In un file operazioni.go definire le seguenti funzioni:

- operazioni(n1, n2 int) (int, int, int)
che accetta due interi e restituisce somma, prodotto e differenza.
Scrivine una versione con variabili di ritorno nominate e una senza (puoi commentare una delle due versioni per caricare un file solo)

Scrivi un main per invocare e testare la funzione. Il programma legge da standard input due int.

*/

package main

import "fmt"

func operazioni(n1, n2 int) (int, int, int) {
	somma := n1 + n2
	prodotto := n1 * n2
	differenza := n1 - n2

	return somma, prodotto, differenza
}

/* func operazioni(n1, n2 int) (somma, prodotto, differenza int) {
	somma = n1 + n2
	prodotto = n1 * n2
	differenza = n1 - n2

	return
}*/

func main() {
	var num1, num2 int

	//fmt.Print("Inserire 2 numeri e ne calcolerò somma, prodotto e differenza in questo ordine: ")
	fmt.Scan(&num1, &num2)
	//fmt.Println("eccone i risultati:")
	fmt.Println(operazioni(num1, num2))
}

//giusto, ma manca libreria dei test :/
