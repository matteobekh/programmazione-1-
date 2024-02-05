/*
Scrivete un programma conta_cifre.go dotato di:

- una funzione `main` che legge una sequenza di parole da standard input fino a incontrare la parola "stop",
analizza una parola alla volta utilizzando la funzione `contaCifre` (vedi sotto) e alla fine stampa:
	- quante parole contengono almeno una cifra
	- per ogni cifra (0, 1, ..., 9), il numero di volte che quella cifra appare nell'intera sequenza di parole

- una funzione
     contaCifre(s string, numCifre *[10]int) (haCifre bool)
che, data una stringa, aggiorna il conteggio delle cifre incontrate e restituisce true se la stringa
contiene almeno una cifra, false altrimenti.

Nota. Le parole in input possono contenere caratteri di qualsiasi tipo
(cifre, lettere, simboli, punteggiatura, lettere accentate, ecc.).

Il programma NON deve fare uso di variabili globali.

Domanda: che prototipo (signature) deve avere la funzione `contaCifre`?
- ha parametri? se sì, quanti, di che tipi e con che finalità?
- restituisce valori? se sì, quanti, di che tipi e con che finalità?
Soffermatevi su questi punti per progettare la funzione prima di scriverla.

Esempio di esecuzione:

$ go run conta_cifre.go
c140 c140
c0m3 t1 ch14m1?
bye bye
stop
5 stringhe con cifre
[0 1 2 3 4 5 6 7 8 9]
[3 5 0 1 3 0 0 0 0 0]
*/

package main

import (
	"fmt"
	"unicode"
)

func contaCifre(s string, numCifre *[10]int) (haCifre bool) {
	/* che, data una stringa, aggiorna il conteggio delle cifre incontrate e restituisce true
	se la stringacontiene almeno una cifra, false altrimenti.*/
	numero := 0
	for _, char := range s {
		if unicode.IsDigit(char) {
			numero = int(char - '0') // trasformo il numero in int e poi faccio numcifre di numero ++
			(*numCifre)[numero]++
			haCifre = true
		}
	}
	return
}

func main() {
	var str string
	//var numeri = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var numeri [10]int
	countCifre := 0 //conta il numero di parole con almeno una cifra
	for {
		fmt.Scan(&str)
		if str == "stop" {
			break
		}
		if contaCifre(str, &numeri) {
			countCifre++
		}
	}
	fmt.Print(countCifre, " stringhe con cifre")
	fmt.Print("\n[0 1 2 3 4 5 6 7 8 9]")
	fmt.Print("\n", numeri)
}
