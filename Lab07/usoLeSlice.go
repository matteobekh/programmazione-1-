/* Scrivere un programma usoLeSlice.go che

legge da linea di comando una lista di almeno 3 parole;
legge da standard input un'altra lista di parole (anche almeno 3).

Quindi fa le seguenti operazioni e man mano stampa il risultato ottenuto:

1. delle due liste fa un'unica lista (slice), da elaborare e a cui, se serve, ci riferiremo con "mySlice"
2. mette in ordine alfabetico (lessicografico) la slice ottenuta (usare una funzione di libreria) e poi la stampa
3. cancella l'ultima parola dalla slice
4. rimuove dalla slice gli elementi di indice dal 2 (incluso) al 4 (escluso)
5. crea una nuova slice con "aa", "bb" e "cc"
6. la inserisce in posizione 1 della vecchia slice (mySlice)
7. legge una nuova parola e la inserisce in posizione 2 della slice mySlice
8. legge una nuova parola e la inserisce in fondo alla slice mySlice
9. estende la slice mySlice con una nuova slice di lunghezza 2 (di stringhe vuote)
10. inserisce alla posizione 3 di mySlice una nuova slice di lunghezza 3 (di stringhe vuote)
11. copia in una nuova slice la slice mySlice ottenuta fin qui
12. dalla copia rimuove l'ultimo elemento e stampa sia la slice mySlice che la copia.

Esempio di esecuzione:

$ ./usoLeSlice.go uno due tre
scrivi almeno 3 parole (seguite da ctrl D)
4 5 6 7
 1. [uno due tre 4 5 6 7]
 2. [4 5 6 7 due tre uno]
 3. [4 5 6 7 due tre]
 4. [4 5 due tre]
 5. [aa bb cc]
 6. [4 aa bb cc 5 due tre]
scrivi una parola: pos2
 7. [4 aa pos2 bb cc 5 due tre]
scrivi una parola: lastPos
 8. [4 aa pos2 bb cc 5 due tre lastPos]
 9. [4 aa pos2 bb cc 5 due tre lastPos  ]
10. [4 aa pos2    bb cc 5 due tre lastPos  ]
11. [4 aa pos2    bb cc 5 due tre lastPos  ]
12. [4 aa pos2    bb cc 5 due tre lastPos  ]
    [4 aa pos2    bb cc 5 due tre lastPos ]
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	terminale := os.Args[1:]
	//var parole string
	paroleStdin := make([]string, 0)
	fmt.Println("scrivi almeno 3 parole seguite da Ctrl D")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		riga := scanner.Text()

		//unire alla slice delle cose inserite da terminale le parole sarebbe comodo, ma il programma vuole un'altra lista
		paroleStdin = append(paroleStdin, riga)
	}
	// 1. delle due liste fa un'unica lista (slice), da elaborare e a cui, se serve, ci riferiremo con "mySlice"
	mySlice := append(terminale, paroleStdin...) //aggiungendo i 3 punti, unisci 2 slice
	fmt.Println("\n1.", mySlice)

	//2. mette in ordine alfabetico (lessicografico) la slice ottenuta (usare una funzione di libreria) e poi la stampa
	sort.Strings(mySlice)
	fmt.Println("2.", mySlice)

	//3. cancella l'ultima parola dalla slice
	mySlice = mySlice[:len(mySlice)-1]
	fmt.Println("3.", mySlice)

	//4. rimuove dalla slice gli elementi di indice dal 2 (incluso) al 4 (escluso)
	mySlice = append(mySlice[:2], mySlice[4:]...)
	fmt.Println("4.", mySlice)

	//5. crea una nuova slice con "aa", "bb" e "cc"
	newSlice := []string{"aa", "bb", "cc"}
	fmt.Println("5.", newSlice)

	//6. la inserisce in posizione 1 della vecchia slice (mySlice)
	mySlice = append(mySlice[:1], append(newSlice, mySlice[1:]...)...)
	fmt.Println("6.", mySlice)

	//7. legge una nuova parola e la inserisce in posizione 2 della slice mySlice
	var parola string
	fmt.Scan(&parola)
	mySlice = append(mySlice, append([]string{parola}, mySlice[2:]...)...)
	fmt.Println("7.", mySlice)

	//8. legge una nuova parola e la inserisce in fondo alla slice mySlice
	fmt.Scan(&parola) //riuso la variabile altrimenti non finisco + di creare variabili da usare solamente una volta
	mySlice = append(mySlice, parola)
	fmt.Println("8.", mySlice)

	//9. estende la slice mySlice con una nuova slice di lunghezza 2 (di stringhe vuote)
	vuota := make([]string, 2)
	mySlice = append(mySlice, vuota...)
	fmt.Println("9.", mySlice)

	//10. inserisce alla posizione 3 di mySlice una nuova slice di lunghezza 3 (di stringhe vuote)
	vuota = append(vuota, "") //aggiungo un altro spazio vuoto nella slice con 2 spazi vuoti della richiesta precedente
	mySlice = append(mySlice[:3], append(vuota, mySlice[3:]...)...)
	fmt.Println("10.", mySlice)

	//11. copia in una nuova slice la slice mySlice ottenuta fin qui
	copia := mySlice
	fmt.Println("11.", copia)

	//12. dalla copia rimuove l'ultimo elemento e stampa sia la slice mySlice che la copia.
	copia = copia[:len(copia)-1]
	fmt.Println("12.", mySlice, copia)

}
