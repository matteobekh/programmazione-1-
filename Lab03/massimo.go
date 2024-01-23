/* Scrivere un programma massimo.go che genera 10 numeri interi casuali tra 10 e 30, li stampa su una riga,
separati da uno spazio, e stampa su una nuova riga il massimo valore generato.
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	const (
		MAX = 30
		MIN = 10
	)

	rand.Seed(time.Now().Unix())

	max := MIN //setto il massimo attuale al numero minimo disponibile
	for i := 0; i < 10; i++ {
		n := rand.Intn(MAX-MIN+1) + MIN
		fmt.Println(n)

		if n > max {
			max = n
		}
	}

	fmt.Println("Il massimo è:", max)
}

/* alcune note:
   rand.Intn(MAX - MIN + 1) genera un numero casuale in [0, 21), da 0 a 20 incluso, tuttavia dovesse uscire un numero
   fuori dall'intervallo sarebbe un problema.
   Quindi scelgo di far generare un numero tra 0 e 20 e sommo 10 così arrivo sempre alla soglia minima*/
