/* Scrivere un programma max_num_cifre_pari.go che, data una sequenza di numeri (da leggere come stringhe),
terminata dalla stringa "000", stampa il numero di cifre pari contenute nel numero che ne contiene più di tutti.
E` possibile risolvere il problema senza tenere in memoria sequenze di numeri. Che tipo di composizione occorrerà
per mettere insieme il conteggio delle cifre pari e il calcolo del massimo di tali conteggi?
*/

package main

import (
	"fmt"
	"io"
)

func main() {
	var numeri string

	fmt.Print("Inserire una sequenza di numeri: ")

	var massimoPari int
	for {
		_, err := fmt.Scan(&numeri)
		if err == io.EOF {
			break
		}
		counterPari := 0
		for _, cifra := range numeri {
			if (cifra-'0')%2 == 0 { //cifra è una runa. Facendo la sottrazione tra il valore in ASCII di cifra e il valore in ASCII del carattere 0, si trova il valore numerico inserito(circa)
				counterPari++
			}
		}

		if counterPari > massimoPari {
			massimoPari = counterPari
		}
	}

	fmt.Println("Massimo numero di cifre pari all'interno di un numero:", massimoPari)
}

/* ancora, in questo caso non capisco l'utilità di far concludere il programma con 000 se tanto poi
i prof scelgono di far concludere il programma quando dai in input la fine del file (Ctrl+D). Se dovessi
fermare il programma a 000 dovrei dividere la stringa iniziale in tante stringhe quante sono i numeri, in seguito
ci sarebbe bisogno dell'analisi della singola stringa: se si trova il numero 000, allora si ferma l'analisi e si esce
dal programma. Alla fine ho optato per la soluzione dei docenti perché + semplice, nonostante non credo si attenga al
100% alla richiesta fatta nella consegna*/
