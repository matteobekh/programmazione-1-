/* Scrivere un programma array.go dotato di:

- una costante DIM = 5 per la dimensione dell'array, dichiarata a livello di package

- una funzione main che inizializza a piacere un array di int di dimensione DIM (ad esempio con numeri 0, 1, 2, ...)
e testa le due seguenti funzioni che lavorano **direttamente sull'array stesso**, senza quindi dichiarare e usare
altri array. Il programma stampa l'array iniziale, l'array dopo aver invocato reverse e l'array dopo aver invocato
switchFirstLast.

- una funzione reverse che inverte l'ordine dei valori in un array di dimensione DIM, mettendo il primo in fondo,
il secondo in penultima posizione e così via, nell'array stesso

- una funzione switchFirstLast che scambia il primo con l'ultimo dei valori in un array di dimensione DIM nell'array
 stesso


Esempio di esecuzione

$ go run array.go

array: [0 1 2 3 4]
reverse: [4 3 2 1 0]
switchFirstLast: [0 3 2 1 4]
*/

package main

import "fmt"

const DIM = 5

func reverse(array *[DIM]int) {
	for i := 0; i < (DIM / 2); i++ {
		array[i], array[DIM-i-1] = array[DIM-i-1], array[i]
	}
}

func switchFirstLast(array *[DIM]int) {
	array[0], array[DIM-1] = array[DIM-1], array[0]
}

func main() {
	var numeri [DIM]int

	for i := 0; i < DIM; i++ {
		fmt.Scan(&numeri[i])
	}
	fmt.Print("array: ", numeri)

	reverse(&numeri) //passo un puntatore, così da poter modificare direttamente lo stesso array, altrimenti modificherei solo una copia dell'array
	fmt.Print("\nreverse: ", numeri)

	switchFirstLast(&numeri)
	fmt.Print("\nswitchFirstLast: ", numeri)

}
