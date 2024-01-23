/* Scrivere un programma cifra_pari.go che, dato un intero da standard input, determina e stampa in che posizione
(procedendo da destra a sinistra) si trova la prima cifra pari del numero. Se il numero non contiene cifre pari,
il programma stampa -1.
nomefile: cifra_pari.go*/

package main

import "fmt"

func main() {
	var num, cifra int

	const divisore = 10 //facendo num % 10 isolo il numero + a destra

	fmt.Println("Inserire un numero:")
	fmt.Scan(&num)

	rimanente := num
	posizione := 1 //contatore di fatto
	nonPari := -1  //che eventualmente sostituisco con la vera cifra pari

	for rimanente != 0 {
		cifra = rimanente % divisore

		if cifra%2 == 0 {
			nonPari = posizione
			break
		}
		rimanente /= divisore
		posizione++
	}
	fmt.Print("Partendo da destra, la prima cifra pari di ", num, " si trova in posizione numero ", nonPari)
}

/* note: questo esercizio si può fare benissimo senza usare sia num che rimanente,
io l'ho usata per poter dire da quale numero siamo partiti, ma si può tranquillamente tagliare e dire semplicemente:
si trova in posizione x e fine. Il mio è, quindi, un semplice accorgimento*/
