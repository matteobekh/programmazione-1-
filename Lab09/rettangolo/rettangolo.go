/*
	Scrivere un programma rettangolo.go, dotato di:

una struttura Rettangolo, con campi base e altezza, in quest'ordine, di tipo int
un metodo String() per Rettangolo che restituisce il disegno del rettangolo, pieno e disegnato con '.', terminato da un new-line ('\n').
Ad esempio, per un Rettangolo di base 5 e altezza 3, il metodo deve restituire la seguente stringa:

.....
.....
.....

Se però la base o l'altezza del rettangolo sono uguali a 0, il metodo String deve restituire il messaggio "rettangolo degenere".

una funzione main() che, dati due numeri naturali come
argomenti sulla linea di comando, emetta nel flusso d'uscita il disegno del rettangolo che ha quei valori per la base e l'altezza (utilizzando il metodo String).

NB: vedi tar per i test
*/
package main

import (
	"fmt"
	"os"
	"strconv"
)

type Rettangolo struct {
	base, altezza int
}

func (rett Rettangolo) String() string { //puntatore non necessario
	disegno := ""
	if rett.base == 0 || rett.altezza == 0 {
		messaggio := "rettangolo degenere"
		return messaggio

	}
	for i := 0; i < rett.altezza; i++ {
		for j := 0; j < rett.base; j++ {
			disegno += "."
		}
		disegno += "\n"
	}
	disegno += "\n"
	return disegno
}

func main() {
	var rettangolo Rettangolo

	rettangolo.base, _ = strconv.Atoi(os.Args[1])
	rettangolo.altezza, _ = strconv.Atoi(os.Args[2])
	fmt.Println()
	fmt.Println(rettangolo.String())

}

/* non funge testmainminimale*/
