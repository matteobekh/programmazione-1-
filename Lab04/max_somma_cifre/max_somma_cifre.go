/* Si vuole scrivere un programma max_somma_cifre.go che

legge da standard input una serie di numeri >= 0, terminata da 999,
trova il numero (escludendo 999) la cui somma delle cifre è la maggiore
e stampa tale somma.

Scrivere il programma max_somma_cifre.go completando il seguente codice. Ricordarsi di compilarlo e testarlo
prima di caricalo su upload.
Si noti che il programma contiene due for annidati, uno servirà per determinare un massimo e uno per calcolare
una somma.*/

package main

import (
	"fmt"
)

func main() {
	var n int

	maxSum := 0
	somma := 0
	for {
		fmt.Scan(&n)
		if n == 999 {
			break
		}

		for n > 0 { //for in cui scompongo il numero e ci faccio la somma, credo
			cifra := n % 10
			somma += cifra
			n /= 10
		}
		if maxSum < somma {
			maxSum = somma
		}
	}
	fmt.Println("la somma + alta tra numero è:", maxSum)
}
