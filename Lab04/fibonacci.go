/* Scrivere un programma fibonacci.go che legge un intero positivo n e stampa i numeri di fibonacci dal
primo all'n-esimo, rappresentandoli come righe di asterischi, ciascuna lunga quanto il numero da rappresentare.
Esempio di esecuzione:
un numero: 6
*
*
**
***
*****
********
Nota. I primi due numeri della serie di Fibinacci sono 1 e 1, dal terzo in poi sono ciascuno la somma dei due numeri
che lo precedono nella serie. Quindi i primi numeri della serie sono: 1, 1, 2, 3, 5, 8, 13, 21, ...
nomefile: fibonacci.go*/

package main

import "fmt"

func main() {
	var num int

	fmt.Print("Inserire numero per sequenza di fibonacci: ")
	fmt.Scan(&num)

	precedente := 0
	attuale := 1

	for i := 1; i <= num; i++ {
		precedente, attuale = attuale, precedente+attuale //modo + semplice per evitare diverse assegnazioni in diverse righe
		for j := 1; j <= precedente; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
