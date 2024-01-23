/* Scrivere un programma indovina_numero.go che chiede all'utente di indovinare un numero intero random n
tra 1 e MAX, (dove MAX è una costante definita nel programma) e ripete la richiesta fino a che l'utente indovina,
oppure fino a un massimo di MAX/2 tentativi.
Il programma stampa il numero di tentativi che sono stati necessari ("t tentativi") per indovinare, oppure il
messaggio "hai perso, il numero era n".
Se il numero digitato dall'utente è fuori dall'intervallo [1,MAX], il tentativo non viene conteggiato e viene
visualizzato il messaggio "fuori intervallo!", senza interrompere l'esecuzione.
Utilizzare la funzione rand.Intn del package "math/rand" per fissare il numero da indovinare.


io scelgo di dare all'utente MAX / 2 tentativi

(scelta personale, si ricorda*/

package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {
	const (
		MAX       = 50 // 50 è un valore preso così, puoi scegliere il numero che preferisci
		tentativi = MAX / 2
	)
	indovina := rand.Intn(MAX)
	var numero int
	fmt.Println("cercare numero nell'intervallo [1,50]") // da cambiare 50 se si vuole cambiare l'intervallo
	numTentativi := 1
	for numTentativi <= tentativi {
		fmt.Print("Inserire numero: ")
		fmt.Scan(&numero)

		if numero < 1 || numero > MAX { //sistemiamo così il caso speciale
			fmt.Println("fuori intervallo!!!")
			numTentativi--
		}

		if numero == indovina {
			fmt.Println(numTentativi, "tentativi sono stati necessari")
			os.Exit(0)
		}

		numTentativi++
	}

	fmt.Println("hai perso, il numero era:", indovina)
}
