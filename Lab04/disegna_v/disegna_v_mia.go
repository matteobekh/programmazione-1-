/* Scrivere un programma disegna_v.go che legge un intero positivo n e stampa una v di altezza (vertice compreso)
n di asterischi.
Esempio di esecuzione:
dimensione v: 4

*     *
 *   *
  * *
   *

nomefile: disegna_v.go*/

package main

import "fmt"

func main() {
	var h int

	fmt.Print("Inserire altezza della V:")
	fmt.Scan(&h)

	for i := 0; i < h; i++ {
		for k := 0; k <= 2*h-1-i; k++ { //for che stampa spazi
			if k == i || k == 2*h-1-i {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
