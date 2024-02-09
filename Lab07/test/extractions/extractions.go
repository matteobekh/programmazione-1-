/* Scrivere un programma extractions.go con due funzioni:


estraiPari(in []int) (out []int) che prende una slice di interi e ne restituisce una che contiene solo i numeri
di quella in ingresso che sono numeri pari.

rimuoviMultipli(m int, in []int) (out []int) che prende un intero e una slice di interi e ne restituisce una
senza i multipli del numero passato di quella in ingresso.
Es.: rimuoviMultipli(5, in) restituisce, a partire da in, una slice senza i multipli di 5.

Nota: non ci sono specifiche per la funzione main.
*/

package main

import "fmt"

func estraiPari(in []int) (out []int) {
	for _, numero := range in {
		if numero%2 == 0 {
			out = append(out, numero)
		}
	}
	return
}

func rimuoviMultipli(m int, in []int) (out []int) {
	for _, numero := range in {
		if numero%m != 0 {
			out = append(out, numero)
		}
	}
	return
}

func main() {

	numeri := []int{4, 6, 8, 12, 55, 60, 100, 54, 17}
	multiplo := 9

	fmt.Println("i numeri pari sono:", estraiPari(numeri))

	fmt.Println("slice senza multipli di", multiplo, ":", rimuoviMultipli(multiplo, numeri))

}
