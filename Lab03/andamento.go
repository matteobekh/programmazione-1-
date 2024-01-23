/*
Scrivere un programma andamento.go che legge da tastiera
una serie (di almeno un elemento) di numeri > -1 e stampa "+" ogni volta che il nuovo valore
e` maggiore o uguale al precedente e "-" altrimenti.
Si ferma quando il numero in input e` -1 e stampa la somma di tutti i numeri letti (escluso -1).

Esempio
-------
go run andamento.go
2 4 7 3 9 1 5 -1
++-+-+
somma: 31
*/

package main

import "fmt"

func main() {
	var attuale, precedente, somma int

	fmt.Scan(&attuale)

	for {
		somma += attuale
		precedente = attuale
		fmt.Scan(&attuale)

		if attuale == -1 {
			break
		}
		if attuale >= precedente {
			fmt.Print("+")
		} else {
			fmt.Print("-")
		}
	}
	fmt.Println()
	fmt.Println("somma", somma)
}
