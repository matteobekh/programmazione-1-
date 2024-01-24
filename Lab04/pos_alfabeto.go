/* Scrivere un programma pos_alfabeto.go che

legge una sequenza di caratteri ASCII (byte) terminati da '.' (es: abc!54LMN.),
per ciascun carattere letto (tranne '.'),

stabilisce se è una lettera minuscola (tra 'a' e 'z'), una cifra o altro
se è una lettera minuscola, stabilisce anche la sua posizione nell'alfabeto
se è una cifra, stabilisce il suo valore numerico
stampa un messaggio costituito da:

il carattere letto
se minuscola, la sua posizione nell'alfabeto (es "f è la 6^a");
se cifra, il suo valore
altrimenti " - altro" (ad esempio "N - altro")

quando termina, stampa la somma delle cifre e  "bye"

Esempio di esecuzione

una sequenza di caratteri terminata da '.': abc!54LMN.
a è la 1^a
b è la 2^a
c è la 3^a
! - altro
5 - 5
4 - 4
L - altro
M - altro
N - altro
somma: 9
bye

*/

package main

import "fmt"

func main() {
	var char byte

	fmt.Print("Inserire una sequenza di caratteri terminati con '.': ")

	somma := 0
	for {
		fmt.Scanf("%c", &char)

		if char == '.' {
			break
		}

		fmt.Print(string(char))
		if char >= 'a' && char <= 'z' {
			fmt.Print(" è la ", char-'a'+1, "'^a\n") //char - a + 1 fa si che si riesce a trovare posizione della lettera nell'alfabeto
		} else if char >= '0' && char <= '9' {
			num := int(char - '0')
			fmt.Print(" - ", num, "\n")
			somma += num
		} else {
			fmt.Println(" - altro")
		}
	}
	fmt.Println("somma:", somma)
	fmt.Println("bye")

}
