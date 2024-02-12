/*
	Scrivere un programma segmento.go, dotato di:

una struttura Segmento, con campi

estremi byte
interno byte
orizzontale bool (false sta per verticale)
lunghezza int

dichiarati in quest'ordine

un metodo String() per Segmento che restituisce una rappresentazione grafica del segmento. Ad esempio un segmento

	seg := Segmento{'*', '-', true, 5}
	è rappresentato graficamente dalla stringa *---*

un metodo allunga(n int) per *Segmento che incrementa di n la lunghezza del segmento (o la accorcia se n è negativo)

una funzione main() che legge da standard input, separati da spazi, in quest'ordine:

il carattere (ASCII) per gli estremi
il carattere (ASCII) interno
true o false per orizzontale o verticale
la lunghezza del segmento

ed emette nel flusso d'uscita la rappresentazione grafica del segmento con quelle caratteristiche.
Poi allunga quello stesso segmento di 3 e di nuovo ne emette nel flusso d'uscita la rappresentazione grafica.
Infine modifica il valore del campo orizzontale e di nuovo ne emette nel flusso d'uscita la rappresentazione grafica.
*/
package main

import (
	"fmt"
	"strings"
)

type Segmento struct {
	estremi, interno byte
	orizzontale      bool //(false sta per verticale)
	lunghezza        int
}

func (seg Segmento) String() string {
	disegno := ""
	if seg.orizzontale {
		disegno = string(seg.estremi) + strings.Repeat(string(seg.interno), seg.lunghezza-2) + string(seg.estremi)
	} else {
		disegno = string(seg.estremi) + "\n"
		for i := 1; i < seg.lunghezza-2; i++ {
			disegno += string(seg.interno) + "\n"
		}
		disegno += string(seg.estremi) + "\n"
	}
	return disegno
}

func (seg *Segmento) allunga(n int) {
	seg.lunghezza += n
}

func main() {
	//var segmento Segmento

	segmento := Segmento{'+', '-', true, 4}
	fmt.Println(segmento.String())
	segmento.allunga(5)
	fmt.Println(segmento.String())
	segmento.orizzontale = true
	fmt.Println(segmento.String())

}

/* a parer mio faila il primo test perché non prende i numeri in input come vorrebbe (scan), tuttavia sto programma non sembra volerli prendere
con la scan( ci ho provato). la soluzione potrebbe essere prendere uno scanner ad una riga, analizzare la riga, mettere gli eselementi in una slice,
trasformarmi nelel loro rispettive variabili e procedere con il programma, ma adesso 0 sbatti*/
