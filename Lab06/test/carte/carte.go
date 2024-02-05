/*
Scrivere un programma carte.go dotato delle seguenti parti:

- definite un tipo struct Carta, per rappresentare carte da poker, con due campi, valore e seme (in quest'ordine), di tipo string,
il primo per il valore ("A", "1", ..., "10", "J", "Q", "K") e il secondo per il seme
("C", "Q", "F", "P" o, volendo, i simboli dei semi "\u2665", "\u2666", "\u2663", "\u2660").

- definite 3 costanti: per il numero di semi (4), di valori (13) e di carte in un mazzo (52).

- definite una funzione
carta(n int) (Carta, bool)
che, dato un int nell'intervallo [0,52), restituisce la corrispondente Carta da poker e true, dove:
le prime 13 sono di cuori, in ordine dall'asso (A) al re (K);
dalla 13 alla 25 sono di quadri; ecc.
Quindi a 10 corrisponde il fante di cuori (JC); a 13 l'asso di quadri (AQ).
Se invece l'argomento non e` nell'intervallo [0,52), restituisce false.

**Suggerimento**: sfruttare la soluzione adottata per associare numeri a nomi di mesi
nell'esercizio sull'estrazione della data.

- definite una funzione
estraiCarta() Carta
che genera un numero casuale in [0, 52) e restituisce la carta corrispondente. (Evitate di duplicare il codice della funzione carta).

- definite una funzione main che chiama estraiCarta e stampa la carta estratta.

- definite una funzione dai4Carte() che restituisce un array di 4 carte estratte con estraiCarta.

- Aggiungete al main le istruzioni per testare anche questa funzione.

**Nota**: Per generare numeri casuali, c'e` la funzione rand.Intn(n int) del pacchetto "math/rand".
Se avete una versione vecchia di Go, per generare sequenze sempre diverse, importare il pacchetto "time" e
usare l'istruzione rand.Seed(time.Now().Unix()) prima di iniziare a generare numeri casuali. Altrimenti non occorre.
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Carta struct {
	valore, seme string
}

const (
	SEMI   = 4
	VALORI = 13
	MAZZO  = 52
	DIM    = 4
)

func carta(n int) (laCarta Carta, flag bool) {

	//Se invece l'argomento non e` nell'intervallo [0,52), restituisce false.
	if n < 0 || n >= 52 {
		return
	}
	seme := n / VALORI // così trovo quale seme abbiamo

	switch seme {
	case 0:
		laCarta.seme = "\u2665" //cuori
	case 1:
		laCarta.seme = "\u2666" //quadri
	case 2:
		laCarta.seme = "\u2663" //fiori
	case 3:
		laCarta.seme = "\u2660" //picche
	}

	valore := n % VALORI
	switch valore {
	case 0:
		laCarta.valore = "A"
	case 2:
		laCarta.valore = "2"
	case 3:
		laCarta.valore = "3"
	case 4:
		laCarta.valore = "4"
	case 5:
		laCarta.valore = "5"
	case 6:
		laCarta.valore = "6"
	case 7:
		laCarta.valore = "7"
	case 8:
		laCarta.valore = "8"
	case 9:
		laCarta.valore = "9"
	case 10:
		laCarta.valore = "10"
	case 11:
		laCarta.valore = "J"
	case 12:
		laCarta.valore = "Q"
	case 13:
		laCarta.valore = "K"
	}
	flag = true
	return
	/* questa funzione può essere fatta in maniera mooolto + semplice:
		carta.valore = valori[n%nvalori] //nvalori ecc sono le costanti che in questa versione abbiamo scritto in maniera differente
	  	carta.seme = semi[n/nvalori]
	  	return*/
}

func estraiCarta() Carta {
	//che genera un numero casuale in [0, 52) e restituisce la carta corrispondente.
	rand.Seed(time.Now().Unix())
	random := rand.Intn(MAZZO)
	cartaRandom, _ := carta(random)
	return cartaRandom
}

func dai4Carte() (Caardo [DIM]Carta) {
	for i := 0; i < DIM; i++ {
		Caardo[i] = estraiCarta()
	}
	return Caardo
	//non ci sono controlli sulla ripetizione della carta, quindi ok
}

func main() {
	var numero int
	fmt.Print("Inserire carta da voler trovare nell'intervallo [0,52): ")
	fmt.Scan(&numero)
	poker, _ := carta(numero)
	fmt.Print("\necco la carta da lei inserita: ", poker)
	fmt.Print("\nadesso le estraggo 4 carte casuali: ", dai4Carte())
}
