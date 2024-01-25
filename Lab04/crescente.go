/*Scrivere un programma crescente.go che legge da standard input una stringa di sole lettere (ASCII) minuscole
e la stampa inserendo un '-' ogni volta che una lettera viene prima in ordine alfabetico della lettera precedente.
Per esempio, data in input la parola ambire, il programma stampa
am-bir-e
nomefile: crescente.go */

package main

import "fmt"

func main() {
	var str string

	fmt.Println("Inserire una stringa di sole lettere ASCII minuscole:")
	fmt.Scan(&str)

	lunghezza := len(str)
	fmt.Print(string(str[0]))

	for i := 1; i < lunghezza; i++ {
		if str[i] < str[i-1] { //confronto tra precedente e attuale, un po' come volevo fare io
			fmt.Print("-")
		}
		fmt.Print(string(str[i]))
	}
	fmt.Println()
}
