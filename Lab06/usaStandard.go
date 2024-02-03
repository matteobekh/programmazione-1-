/* Scrivere un programma usaStandard.go che dichiara e inizializza le seguenti costanti:

VOCALI = "aeiouAEIOU"
una stringa S1 a piacere lunga 1 solo carattere
un'altra stringa S2 a piacere, più lunga di 1.

Il programma deve leggere da standard input una parola e utilizzare le funzioni dei pacchetti strings e strconv per
svolgere i seguenti compiti:
(1) verificare se la stringa letta contiene S2 come sottostringa (nel caso stampare un messaggio)
(2) verificare se la stringa letta contiene almeno una vocale (nel caso stampare un messaggio)
(3) determinare quante occorrenze di S1 ha la stringa letta (e stampare un messaggio)
(4) determinare in che posizione si trova la prima vocale della stringa letta e in che posizione si trova
l'ultima vocale della stringa letta e stampare i risultati ottenuti
(5) stampare la stringa ottenuta concatenando S2 3 volte con se stessa.
(6) stampare la stringa ottenuta concatenando S1 5 volte con se stessa.
(7) estrarre le eventuali cifre contenute nella parola letta, concatenarle in una stringa s nello stesso ordine
in cui le incontra e stamparla
(8) convertire a int la stringa s e stampare l'int n ottenuto
(usare fmt.Printf("intero %d\n", n))
(9) ottenere il float64 f corrispondente a "0.s" e stamparlo
(usare fmt.Printf("float %f\n", f)).
Esempio di esecuzione con S1 = "c" e S2 = "sc", su input "1scirocco45"
./usaStandard
1scirocco45

(1) 1scirocco45 contiene sc
(2) 1scirocco45 ha vocali
(3) 1scirocco45 ha 3 c
(4) la prima vocale di 1scirocco45 è in posizione 3 , l'ultima in posizione 8
(5) scscsc
(6) ccccc
(7) stringa 145
(8) intero 145
(9) float 0.145000
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

const (
	VOCALI = "aeiouAEIOU"
	S1     = "c"
	S2     = "sc"
)

func main() {
	var str string

	fmt.Print("Inserire una stringa: ") // inserire 1scirocco45
	fmt.Scan(&str)

	//(1) verificare se la stringa letta contiene S2 come sottostringa (nel caso stampare un messaggio)
	fmt.Print("(1) ", str)
	if !strings.Contains(str, S2) {
		fmt.Print(" non ")
	}
	fmt.Print(" contiene ", S2)

	//(2) verificare se la stringa letta contiene almeno una vocale (nel caso stampare un messaggio)
	fmt.Print("\n(2) ", str)
	if !strings.ContainsAny(str, VOCALI) {
		fmt.Print(" non ")
	}
	fmt.Print(" contiene vocali")

	//(3) determinare quante occorrenze di S1 ha la stringa letta (e stampare un messaggio)
	fmt.Print("\n(3) ", str)
	fmt.Print(" ha ", strings.Count(str, S1), " ", S1)

	//(4) determinare in che posizione si trova la prima vocale della stringa letta e in che posizione si trova l'ultima vocale della stringa letta e stampare i risultati ottenuti
	fmt.Print("\n(4) la prima vocale in ", str, " è in posizione ", strings.IndexAny(str, VOCALI), ", l'ultima in posizione ", strings.LastIndexAny(str, VOCALI))

	//(5) stampare la stringa ottenuta concatenando S2 3 volte con se stessa.
	fmt.Print("\n(5) ", strings.Repeat(S2, 3))

	//(6) stampare la stringa ottenuta concatenando S1 5 volte con se stessa.
	fmt.Print("\n(6) ", strings.Repeat(S1, 5))

	//(7) estrarre le eventuali cifre contenute nella parola letta, concatenarle in una stringa s nello stesso ordine in cui le incontra e stamparla
	s := ""
	for _, char := range str {
		if unicode.IsDigit(char) {
			s += string(char)
		}
	}
	fmt.Print("\n(7) stringa ", s)

	//(8) convertire a int la stringa s e stampare l'int n ottenuto
	n, _ := strconv.Atoi(s)
	fmt.Printf("\n(8) intero %d\n", n)

	//(9) ottenere il float64 f corrispondente a "0.s" e stamparlo
	float, _ := strconv.ParseFloat("0."+s, 64)
	fmt.Printf("(9) float %f\n", float)

}
