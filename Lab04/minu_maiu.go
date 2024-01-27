/* Scrivere un programma minu_maiu.go che legge da standard input una stringa (di rune) e stabilisce se la
stringa contiene solo minuscole, solo maiuscole o sia minuscole che maiuscole, quindi stampa un messaggio opportuno
(es. "solo minuscole", "solo minuscole", "sia minuscole che maiuscole").
Utilizzare le funzioni del pacchetto unicode.
*/

package main

import (
	"fmt"
	"unicode"
)

func main() {
	var str string

	fmt.Print("Inserire una stringa di rune: ")
	fmt.Scan(&str)
	var minuscole, maiuscole bool

	for _, char := range str {
		if unicode.IsLower(char) {
			minuscole = true
		} else if unicode.IsUpper(char) {
			maiuscole = true
		}
		if minuscole && maiuscole {
			break
		}
	}

	fmt.Print("minuscole: ", minuscole, " maiuscole: ", maiuscole)

}
