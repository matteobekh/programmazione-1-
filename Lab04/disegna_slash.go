/* Scrivere un programma disegna_slash.go che legge un intero positivo n e stampa un backslash (\) di altezza n
composto da asterischi.
Esempio di esecuzione:
dimensione \: 3

*
 *
  *

*/
//secondo me posso contare gli spazi in base al contatore:
//i = 0, 0 spazi, i= 1, 1 spazio, i = 2, 2 spazi e così via
package main

import "fmt"

func main() {
	var h int
	fmt.Println("Si richiede altezza dello slash (\\)")
	fmt.Scan(&h)

	for i := 0; i < h; i++ {
		for k := 0; k < i; k++ { //for che stampa spazi
			fmt.Print(" ")
		}

		fmt.Print("*")
		fmt.Println()
	}
}
