/* Scrivere un programma ultima_pioggia.go che legge da standard input una sequenza di numeri interi
(terminata da EOF, da tastiera prodotto con ctrl-D) che indicano i mm di pioggia caduti (0 se non ha piovuto)
ogni giorno in una sequenza successiva di giorni e stampa (l'indice del)l'ultimo giorno in cui ha piovuto.
Nota su come "monitorare" la lettura con Scan e interrompere un ciclo for quando si intercetta EOF:

_, err := fmt.Scan(&myVar)
		if err == io.EOF {
			break
		}
*/

package main

import (
	"fmt"
	"io"
)

func main() {
	var num int

	fmt.Println("Inserire mm di pioggia degli ultimi giorni...")
	counterPioggia := 0
	//counter := 1
	for {
		_, err := fmt.Scanf("%d", &num)
		if err == io.EOF {
			break
		}
		if num > 0 {
			counterPioggia++
		}
		//counter++
	}
	fmt.Println("l'ultimo giorno di pioggia è stato il giorno", counterPioggia)
}

/* per qualche motivo il numero stampato non funge bene, ma non trovo altre soluzioni. Quella del prof invece
è una soluzione che tecnicamente non coincide con la consegna, quindi ritengo oppportuno caricare questa versione
nonostante non sia giusta*/
