/* Scrivere un programma es0.go che fa le seguenti cose:
legge un byte
lo stampa
stampa il byte precedente, il byte stesso, e quello successivo in ordine lessicografico (ASCII). Ad esempio,
con 'd' come input, deve stampare: c d e

stabilisce se è una lettera tra A e L, o altro (stampa "A-L" o "altro")
poi legge una stringa e la stampa in verticale, una runa per riga.

Ad esempio "città" diventa:

c
i
t
t
à*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	var carattere byte

	//legge un byte
	fmt.Print("Inserire byte: ")
	fmt.Scanf("%c", &carattere)

	//lo stampa
	fmt.Printf("stampa: %c", carattere)

	//stampa carattere precedente e successivo
	fmt.Printf("\n%c, %c, %c\n", carattere-1, carattere, carattere+1)

	//si trova tra A-L o altro. tutto minuscolo per evitare problemi con maiuscole
	if strings.ToLower(string(carattere)) >= "a" && strings.ToLower(string(carattere)) <= "l" {
		fmt.Println("a-l")
	} else {
		fmt.Println("altro")
	}

	//legge una stringa e la stampa in verticale, una runa per volta
	var str string
	fmt.Print("Inserire una stringa: ")
	fmt.Scan(&str)

	for _, char := range str {
		fmt.Println(string(char))
	}
}
