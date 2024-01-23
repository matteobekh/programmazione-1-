/* Scrivere un programma primo.go che, dato un numero intero su standard input, stampa "primo" se il numero è primo,
 "non primo" altrimenti.
Suggerimento: occorre determinare il primo numero che è un divisore (se c'è).
Domanda: dato in input n, quante iterazioni dovrò fare al massimo?
nomefile: primo.go

I primi metodi per il calcolo dei numeri primi sono chiamati test di primalità e si basano sul test di divisione
per tutti i numeri inferiori alla radice quadrata del numero scelto: Se è divisibile per uno di essi, è composto,
se non è divisibile per uno di essi, è il primo.*/

package main

import "fmt"

func main() {
	var num int

	fmt.Print("Dammi un numero e scopriamo se è primo: ")
	fmt.Scan(&num)

	var isPrime bool
	for divisore := 2; divisore < num; divisore++ {
		if num%divisore == 0 {
			isPrime = true
			break
		}
	}

	if isPrime {
		fmt.Println(num, "non è primo")
	} else {
		fmt.Println(num, "è primo")
	}
}

/* questa è una versione semplificata, tuttavia, cercando su interent, si scopre che ci si può fermare alla radice
quadrata del numero chiesto. SI può aggiungere con la libreria math e aggiungi il valore della radice quadrata. Nel
for invece, piuttosto che arrivare fino al numero inserito, si modifica per farlo arrivare solamente fino alla radice,
per il resto il problema è lo stesso :D */
