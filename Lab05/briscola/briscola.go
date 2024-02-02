/* In un file briscola.go si scriva una funzione

	punti(s string) int

che accetta una stringa di 3 caratteri che rappresenta una mano di tre carte e restituisce il punteggio
complessivo relativo per il gioco della briscola.
Ad esempio per la mano "53J" restituisce 12 (10 della carta 3 + 2 del fante).
Se la stringa contiene un simbolo che non corrisponde a nessuna carta, la funzione restituisce -1.

Si scriva un main per invocare e testare la funzione. Il programma legge da standard input una stringa e
controlla che sia di tre caratteri. Poi stampa
mano <mano>: <tot> punti

Punti a briscola:
Asso (A): 11
3: 10
Re (K): 4
Donna (Q): 3
Fante (J): 2
7, 6, 5, 4, 2: 0
*/

package main

import "fmt"

func punti(s string) (punteggio int) {

	for _, char := range s {

		switch char {
		case 'A':
			punteggio += 11
		case '3':
			punteggio += 10
		case 'K':
			punteggio += 4
		case 'Q':
			punteggio += 3
		case 'J':
			punteggio += 2
		case '7', '6', '5', '4', '2':
			continue
		default:
			return -1
		}
	}
	return
}

func main() {
	var str string

	fmt.Scan(&str)
	if len(str) != 3 {
		fmt.Println("Inserire stringa da 3 caratteri...")
		return
	}
	fmt.Print("mano ", str, ": ", punti(str), " punti")
}
