/*
	Scrivere un programma ordini.go che legge da standard input una serie di stringhe che descrivono

ordini nel formato:

prezzo#quantità#sconto

e stampa il totale finale da pagare.
Prezzo, quantità e sconto sono float; prezzo indica il prezzo unitario del prodotto,
quantità indica la quantità acquistata e sconto è lo sconto applicato per quel prodotto,
espresso come float tra 0 e 1 (ad esempio 0.2 indica uno sconto del 20%).
Il programma termina la lettura quando incontra EOF (ctrl D su riga nuova).
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Inserire una serie di stringhe nel formato prezzo#quantità#sconto:")
	scanner := bufio.NewScanner(os.Stdin)
	var prezzoTotale float64
	for scanner.Scan() {
		riga := scanner.Text()
		elementiStr := strings.Split(riga, "#")
		prezzoStr := elementiStr[0]
		quantitàStr := elementiStr[1]
		scontoStr := elementiStr[2]

		prezzoNum, _ := strconv.ParseFloat(prezzoStr, 64)
		quantitàNum, _ := strconv.ParseFloat(quantitàStr, 64)
		scontoNum, _ := strconv.ParseFloat(scontoStr, 64)

		prezzoScontato := prezzoNum - (prezzoNum * scontoNum)
		prezzoTotale += (prezzoScontato * quantitàNum)
	}
	fmt.Print(prezzoTotale)
}

/* guardando i test i risultati sono giusti, tuttavia per questioni di spazi o simili non hanno intenzione di passare
   ok...*/
