/* Scrivere un programma euclide.go che legge da standard input due interi a e b, con a >= b,
calcola il MCD tra i due numeri usando l'algoritmo di Euclide e lo stampa
(seguito da newline, senza stampare altro).
Algoritmo di Euclide:
Dati due numeri naturali a e b:

controlla se b è zero.
se lo è, a è il MCD.
se non lo è, assegna ad r il resto della divisione a / b,
poni a = b e b = r e ripeti da 1.

*/

package main

import "fmt"

func Euclide(n1, n2 int) int {
	var resto int
	if n2 == 0 { // b
		return n1
	}

	for n2 != 0 {
		resto = n1 % n2
		n1, n2 = n2, resto
	}
	return n1
}

func main() {
	var a, b int

	fmt.Print("Inserire 2 numeri a e b dove a >= b: ")
	fmt.Scan(&a, &b)

	fmt.Println("Il tuo MCD tra", a, "e", b, "è:", Euclide(a, b))
}
