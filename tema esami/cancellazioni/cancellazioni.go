package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
che, per ogni numero n >0 (intero) presente in lista,
cancella il numero stesso e gli n elementi successivi della lista
(o quelli che ci sono per arrivare alla fine della lista)
e restituisce la nuova lista così prodotta;
*/
func cancella(lista []string) []string {
	var risultato []string

	salto := 0

	for i := 0; i < len(lista); i++ {
		if salto > 0 {
			salto--  // Decrementa il contatore degli elementi da saltare
			continue // Salta l'elemento corrente
		}
		num, err := strconv.Atoi(lista[i])
		if err != nil { //se c'è un errore, allora non è un numero, quindi append
			risultato = append(risultato, lista[i])
		} else {
			salto = num
		}
	}
	return risultato
}

/*
una funzione main() che legge da file (il cui nome viene passato
come parametro su linea di comando) un testo di parole (sequenze di caratteri separate da whitespace),
tra cui anche numeri interi non negativi, stampa la lista delle parole lette e poi
la nuova lista ottenuta cancellando, per ogni numero n presente nella lista originale,
il numero stesso e gli n elementi successivi (vedi sopra).
*/
func main() {
	if len(os.Args) != 2 {
		fmt.Println("Fornire 1 nome di file")
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("File non accessibile")
		os.Exit(2)
	}
	defer file.Close()

	var elenco []string

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		elemento := scanner.Text()
		elenco = append(elenco, elemento) //riempo una slice con tutti gli elementi
	}
	fmt.Println(elenco)
	fmt.Println(cancella(elenco))
}
