/* Scrivere un programma stampaAlternata.go che legge da standard input del testo su più righe
(terminato da EOF) e stampa prima le righe pari e poi le righe dispari
(considerate la prima riga del testo la riga 1 (e non 0)).
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	righe := []string{}
	righeContatore := 0
	for scanner.Scan() {
		righe = append(righe, scanner.Text()) //scanner.Text() è la singola riga
		righeContatore++
	}
	pari := []string{}
	dispari := []string{}
	for i := 0; i < righeContatore; i++ {
		if i%2 == 0 {
			dispari = append(dispari, righe[i]) // invertito altrimenti stampa le righe dispari in pari e viceversa
		} else {
			pari = append(pari, righe[i])
		}
	}
	//fmt.Println("pari:")
	for i := 0; i < len(pari); i++ {
		fmt.Println(pari[i])
	}

	//fmt.Println("dispari:")
	for i := 0; i < len(dispari); i++ {
		fmt.Println(dispari[i])
	}
}

/* ad occhio sembra funzionare, tuttavia non supera i test :/ */
