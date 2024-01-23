/* Scrivere un programma vicino.go che legge da standard input una serie di 5 valori interi tra 0 e 20
(ignorando eventuali valori fuori da questo intervallo)
 e stampa il valore piu' vicino a TARGET, tra quelli letti, dove TARGET è una costante definita nel programma.
Nota. Il package math ha una funzione Abs che calcola (e restituisce) il valore assoluto.*/

package main

import (
	"fmt"
	"math"
)

func main() {
	const TARGET = 15
	var num, vicino int

	fmt.Println("Inserire 5 valori...")
	fmt.Scan(&num)

	vicino = num
	for i := 1; i < 5; i++ {
		fmt.Scan(&num)
		if num < 0 || num > 20 {
			i--
			continue
		}
		if math.Abs(float64(TARGET-num)) < math.Abs(float64(TARGET-vicino)) {
			vicino = num
		}
	}

	fmt.Print("Il valore + vicino a ", TARGET, " è ", vicino)
}
