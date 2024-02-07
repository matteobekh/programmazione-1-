/* Scrivere una programma quadrati.go che legge da linea di comando una sequenza di numeri interi non negativi
e stampa solo quelli che sono dei quadrati (1, 4, 9, ecc.). Il programma deve essere dotato di una funzione
isSquare(n int) bool
che restituisce true se n è un quadrato, false altrimenti.
*/

package main

import (
	"fmt"
	"os"
	"strconv"
)

func isSquare(n int) bool {

	for i := 1; i*i <= n; i++ {
		if i*i == n {
			return true
		}
	}
	return false
}

func main() {

	numeri := os.Args[1:]

	quadrati := []int{}
	for _, numeroStr := range numeri {
		numero, _ := strconv.Atoi(numeroStr)
		if isSquare(numero) {
			quadrati = append(quadrati, numero)
		}
	}
	fmt.Print(quadrati)
}

/* al momento passa 2/3 test. per farlo passare mi sembra di capire che sia necessario usare uno scanner perché
dai test sembra che i numeri possano anche essere divisi tra di loro da + righe:
esempio: 5 8       31

45 645 33*/
