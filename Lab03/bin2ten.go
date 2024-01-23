/* Scrivere un programma bin2ten.go che converte un intero composto di soli 0 e 1 nel valore corrispondente al
numero binario rappresentato (es. 101 -> 5) e lo stampi. Nel caso il numero contenga altre cifre,
il programma stampa un messaggio di errore.

questo programma funge in maniera molto simile a somma_cifre:
basta estrarre le cifre una alla volta, analizzarle in base al contatore del for e poi fare la somma delle potenze
(se hai dubbi, vedi come trasformare un numero binario in decimale)!!!!!*/

package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var numBin int

	fmt.Print("Inserire un numero BINARIO: ")
	fmt.Scan(&numBin)

	binario := numBin
	numDec := 0.0
	cifra := 0
	potenza := 0
	for numBin > 0 {
		cifra = numBin % 10
		if cifra != 1 && cifra != 0 {
			fmt.Println("Il numero non è binario o non è stato scritto correttamente!")
			os.Exit(1)
		}
		if cifra == 1 {
			numDec = numDec + math.Pow(2.0, float64(potenza))
		}
		potenza++
		numBin /= 10
	}
	fmt.Print("Il numero ", binario, " in decimale è ", numDec)
}
