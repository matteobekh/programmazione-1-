/* Scrivere un programma somma_cifre.go che stampa la somma delle cifre di un numero intero positivo fornito
da standard input.*/

package main

import "fmt"

func main() {
	var num int

	fmt.Print("Inserire un numero: ")
	fmt.Scan(&num)

	cifra := 0
	somma := 0
	for num > 0 {
		cifra = num % 10
		somma += cifra
		num /= 10
	}
	fmt.Println("La somma delle cifre del numero è", somma)
}

/* il programma si può migliorare dicendo che la somma delle cifre di NUM è ...
   per farlo, basta salvare la variabile di num in un'altra variabile che non cambia nel corso dello svolgimento del
   programma! ma funge comunque, quindi chill  */
