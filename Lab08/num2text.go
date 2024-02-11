/* Scrivere un programma num2text.go per convertire un numero intero non negativo nella sequenza delle parole corrispondenti alle sue cifre.
Il programma legge un intero non negativo da standard input, per ogni nuova (non incontrata finora) cifra del numero chiede il nome corrispondente
(e alimenta un dizionario), e infine stampa la sequenza delle parole corrispondenti alle sue cifre.
Ad esempio, per il numero 203, il programma stampa:

due - zero - tre

Esempio di esecuzione

$ go run num2text.go
un numero: 622
parola per 2 ? due
parola per 6 ? sei
sei - due - due*/

package main

import "fmt"

func main() {
	var num int

	mappa := make(map[int]string)

	fmt.Print("un numero: ")
	fmt.Scan(&num)
	numcopia := num

	for num > 0 {
		cifra := num % 10

		numStr, isPresent := mappa[cifra]

		if !isPresent {
			fmt.Print("parola per ", cifra, " ? ")
			fmt.Scan(&numStr)

			mappa[cifra] = numStr
		}
		num /= 10
	}

	var parole []string
	for numcopia > 0 {
		cifra := numcopia % 10
		parole = append(parole, mappa[cifra])
		numcopia /= 10
	}
	//stampo al contrario per avere l'ordine corretto
	for i := len(parole) - 1; i >= 0; i-- {
		fmt.Print(parole[i])
		if i > 0 { // serve per evitare il trattino extra alla fine
			fmt.Print(" - ")
		}
	}
}
