/* In questo esercizio prendiamo in considerazione solo caratteri ASCII (byte).
In un file funzioni.go definire le seguenti funzioni:

- hasUpper(s string) bool
La funzione riceve una stringa come parametro e restituisce true se la stringa contiene almeno una lettera
latina maiuscola (tra 'A' e 'Z'), false altrimenti.

- firstUpper(s string) int
La funzione riceve una stringa come parametro e restituisce la posizione della prima lettera latina maiuscola
(tra 'A' e 'Z'), se la stringa ne contiene almeno una, -1 altrimenti.

- lastUpper(s string) int
La funzione riceve una stringa come parametro e restituisce la posizione dell'ultima lettera latina maiuscola
(tra 'A' e 'Z'), se la stringa ne contiene almeno una, -1 altrimenti.

- countDigitsLettersOthers(s string) int int int
La funzione riceve una stringa come parametro, conta quante cifre, quante lettere e quanti altri caratteri
(né cifre né lettere) contiene e restituisce i tre risultati in questo ordine.

- isPalindrome(s string) bool
La funzione riceve una stringa come parametro e restituisce true se la stringa è palindroma, e false altrimenti.
Una parola è palindroma se può essere letta  sia da sinistra a destra che da destra a sinistra. Ad esempio
"osso" e "ingegni" sono palindrome, ma anche "" (la stringa vuota) e "t" (le stringhe di un solo carattere).

Scrivete infine un main che legge una stringa da standard input, usa le funzioni qui sopra per determinare
se la stringa letta contiene maiuscole, in che posizione è la prima maiuscola, in che posizione è l'ultima maiuscola,
 quante cifre, lettere e altri caratteri contiene, se è palindroma, e stampa i risultati
 ("ha maiuscole" / "non ha maiuscole", "prima maiuscola in posizione ...", “palindroma" / "non palindroma", ecc.).
*/

package main

import (
	"fmt"
	"unicode"
)

func hasUpper(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			return true
		}
	}
	return false
}

func firstUpper(s string) int {

	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			return i
		}
	}
	return -1
}

func lastUpper(s string) int {

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] >= 'A' && s[i] <= 'Z' {
			return i
		}
	}
	return -1
}

func countDigitsLettersOthers(s string) (cifre, lettere, altri int) {
	/*for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			cifre++
		}
		if s[i] >= 'A' && s[i] <= 'Z' || s[i] >= 'a' && s[i] <= 'z' {
			lettere++
		} else {
			altriCaratteri++
		}
	}
	return cifre, lettere, altriCaratteri */
	l := len(s)
	for indice := 0; indice < l; indice++ {
		r := rune(s[indice])
		if unicode.IsDigit(r) {
			cifre++
		} else if unicode.IsLetter(r) {
			lettere++
		} else {
			altri++
		}
	}
	return
	/* altra soluzione in realtà molto comoda*/
}

func isPalindrome(s string) bool {
	if s == "" {
		return true
	}
	if len(s) == 1 {
		return true
	}

	for i := 0; i <= len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var str string

	//fmt.Print("Inserire una stringa da analizzare: ")
	fmt.Scan(&str)

	if !hasUpper(str) {
		fmt.Print("non ")
	}
	fmt.Println("ha maiuscole")
	fmt.Println("prima maiuscola in posizione", firstUpper(str))
	fmt.Println("ultima maiuscola in posizione", lastUpper(str))
	//fmt.Print("\ncifre, lettere, altro: ")
	fmt.Println(countDigitsLettersOthers(str))

	if !isPalindrome(str) {
		fmt.Print("non ")
	}
	fmt.Println("palindroma")

}
