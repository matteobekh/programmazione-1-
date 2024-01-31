/* Scrivere un programma cesare.go che legge da standard input un numero intero non negativo k
(la chiave di cifratura), uno spazio e una sequenza di lettere ASCII minuscole consecutive
(sulla stessa riga e senza spazi) terminate da invio e  cifra le lettere lette con il cifrario di Cesare,
usando come chiave k (quella fornita dall'utente) e stampa il testo cifrato:
ogni lettera del testo in chiaro è sostituita nel testo cifrato dalla lettera che si trova k posizioni
dopo nell'alfabeto, ritornando alla lettera 'a' dopo la 'z'.
Esempi di esecuzione:
2 zaprb
bcrtd
100 abcd
wxyz */

package main

import "fmt"

func main() {
	var chiave int

	fmt.Print("Chiave: ")
	fmt.Scan(&chiave)

	fmt.Print("Testo: ")

	var char, charCifrato rune

	for {
		fmt.Scanf("%c", &char)

		if char == '\n' { //come richiesto da consegna, quando fai invio, il programma conclude
			break
		}

		charCifrato = rune((int(char-'a')+chiave)%26 + 'a')
		fmt.Print(string(charCifrato))
	}
	fmt.Println("\necco il file cifrato...")
}

/* sembra funzionare, fare attenzione però all'ultimo carattere che stampa che per qualche motivo ne stampa uno in
+ e lo sbatti di controllare perché fa così manca. arrivederci*/
