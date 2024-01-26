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

	for i := 0; i < h-1; i++ {
		for k := 0; k < i; k++ {
			fmt.Print(" ")
		}
		fmt.Print("*")
		//con la roba precedente ho disegnato il backslash, escludendo però l'ultima riga del vertice!

		for j := 0; j < 2*(h-i-2)+1; j++ {
			fmt.Print(" ")
		}
		fmt.Println("*")
	}

	for w := 0; w < h-1; w++ {
		fmt.Print(" ")
	}
	fmt.Println("*")
}
