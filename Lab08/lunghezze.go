/* Scrivere un programma lunghezze.go che legge riga per riga un testo da standard input (potete usare la ridirezione), terminato da EOF,
e stampa quante parole ci sono di lunghezza 1, quante di lunghezza 2, ecc.
In particolare sperimentare quattro tipi di stampa:

una stampa della mappa
una stampa degli elementi (le coppie k:v) della mappa, non importa in che ordine (for range)
una stampa degli elementi in ordine dalla lunghezza minima a quella massima, comprese le lunghezze eventualmente non presenti nella mappa
(che tipo di for?)
una stampa degli elementi in ordine dalla lunghezza minima a quella massima, escluse le lunghezze non presenti nella mappa
(come evito di stampare le lunghezze che hanno 0 parole associate?).

Il programma deve essere dotato di:

una funzione

func aggiornaConteggio(m map[int]int, riga string)

	che aggiorna la mappa dei conteggi (e che deve essere usata dal main);


una funzione

func findMaxKey(m map[int]int) int

	che restituisce la massima chiave presente nella mappa.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func aggiornaConteggio(m map[int]int, riga string) {
	parole := strings.Fields(riga)
	for _, parola := range parole {
		m[len(parola)]++
	}
}

func findMaxKey(m map[int]int) int {
	massimo := 0
	for key := range m {
		if key > massimo {
			massimo = key
		}
	}
	return massimo
}

func main() {

	mappa := make(map[int]int)

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		aggiornaConteggio(mappa, scanner.Text())
	}

	fmt.Println("stampa della mappa:")
	for lunghezza, conteggio := range mappa {
		fmt.Printf("Lunghezza %d, %d\n", lunghezza, conteggio)
	}

	fmt.Println("stampa degli elementi della mappa:")
	for _, valore := range mappa {
		fmt.Printf("%d ", valore)
	}
	fmt.Println()

	fmt.Println("stampa degli elementi in ordine crescente di lunghezza:")
	chiavi := make([]int, 0, len(mappa))
	for k := range mappa {
		chiavi = append(chiavi, k)
	}
	sort.Ints(chiavi)
	for _, lunghezza := range chiavi {
		fmt.Printf("%d", lunghezza) //non sto stampando i valori all'interno della mappa però dettagli
	}
	fmt.Println()
	fmt.Println("\nStampa degli elementi in ordine crescente di lunghezza, escludendo le lunghezze non presenti:")
	maxKey := findMaxKey(mappa)
	for i := 1; i <= maxKey; i++ {
		if mappa[i] > 0 {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}

/* questo programma non funge molto bene, anzi in realtà (specialmente il main) fungono abbastanza male, però sto ancora
capendo come funzionano ste cose. Eventualmente tornerò qui a correggere le boiate scritte qua sopra */
