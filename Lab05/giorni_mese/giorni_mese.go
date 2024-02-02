/*
In un file giorni_mese.go definire una funzione

	giorniInMese(mese int) int

che, dato come parametro il numero corrispondente a un mese, restituisce il numero di giorni di quel mese
(28 per febbraio; 30 per aprile, giugno, settembre, novembre; 31 per G,M,M,L,A,O,D).

Assumiamo che il numero passato come parametro sia in [1,12], quindi non facciamo controlli.

Usiamo uno switch. Quanti casi conterrà (incluso il caso default, se lo si usa)? 3 casi, 12 casi o quanti?

Scrivere inoltre una funzione
	main()
per invocare e testare la funzione giorniInMese.

La funzione main deve leggere da standard input una stringa nel formato gg-mm-aaaa
(vedi funzione Atoi del pacchetto strconv) e stampare "il mese <x> ha <tot> giorni", dove x e tot sono numeri,
usando la funzione giorniInMese per determinare tot. Chiamare il programma giorni_mese.go e caricarlo su upload
(dopo aver lanciato i test)

Esempio di esecuzione
---------------------
Se il programma legge

09-11-2022

deve stampare

il mese 11 ha 30 giorni
*/

package main

import (
	"fmt"
	"strconv"
)

func giorniInMese(mese int) int {
	//non facciamo controlli, come richiesto da consegna (valore in [1,12])
	switch mese { //(28 per febbraio; 30 per aprile, giugno, settembre, novembre; 31 per G,M,M,L,A,O,D).
	case 2:
		return 28
	case 4, 6, 9, 11:
		return 30
	default:
		return 31
	}
	/* var numGiorni int
	switch mese {
	case 1, 3, 5, 7, 8, 10, 12:
		numGiorni = 31
	case 4, 6, 9, 11:
		numGiorni = 30
	case 2:
		numGiorni = 28
	}
	return numGiorni
	metodo dei docenti
	*/
}

func main() {
	var str string

	fmt.Print("Inserire data nel formato gg-mm-aaaa: ")
	fmt.Scan(&str)

	/* La funzione main deve leggere da standard input una stringa nel formato gg-mm-aaaa
	(vedi funzione Atoi del pacchetto strconv) e stampare "il mese <x> ha <tot> giorni", dove x e tot sono numeri,
	usando la funzione giorniInMese per determinare tot.*/
	meseInt, _ := strconv.Atoi(str[3:5]) //3 incluso, 5 escluso
	fmt.Print("Il mese ", meseInt, " ha ", giorniInMese(meseInt), " giorni")
}
