/*
	Scrivere un programma 'draw.go' dotato di:
	 - una funzione
	   DrawPoint(c byte, k int) string

che restituisce una stringa formata da k spazi bianchi
seguiti dal carattere c

  - una funzione
    DrawSegment(c byte, k, l int) string

che restituisce una stringa formata da k spazi bianchi
seguiti da l caratteri c

  - una funzione
    main()

che legge come parametri da linea di comando (in quest'ordine)
due numeri interi l e n e un carattere (byte) c,
e, se l e n sono > 0, produce su standard output una scala di n gradini
di lunghezza e altezza l disegnati usando il carattere c (vedi esempi sotto),
altrimenti non fa niente.

Si può assumere che il programma venga lanciato con tre parametri
dei tipi attesi (non occorre cioè fare controlli).

Esempi
------

$ ./draw 3 1 x
xxx

	x
	x

$ ./draw 3 2 +
+++

	  +
	  +
	  +++
		+
		+
*/
package main

import (
	"fmt"
	"os"
	"strconv"
)

// che restituisce una stringa formata da k spazi bianchi seguiti dal carattere c
func DrawPoint(c byte, k int) string {
	output := ""
	for i := 0; i < k; i++ {
		output += " "
	}
	output += string(c)
	return output
}

// che restituisce una stringa formata da k spazi bianchi seguiti da l caratteri c
func DrawSegment(c byte, k, l int) string {
	output := ""

	for i := 0; i < k; i++ { //stringa formata da k spazi bianchi
		output += " "
	}

	for i := 0; i < l; i++ {
		output += string(c)
	}
	return output
}

func main() {
	// non van fatti controlli sul numero di input e sulla tipologia, si da per scontato che siano giusti!
	lString := os.Args[1] //lunghezza e altezza gradini
	nString := os.Args[2] // numero gradini
	cString := os.Args[3] // carattere da usare per il disegno (deve essere byte)

	cSlice := []byte(cString)

	c := cSlice[0] //convertito la stringa c in byte (non so se è necessario, ma dalla consegna sembra di si )
	l, _ := strconv.Atoi(lString)
	n, _ := strconv.Atoi(nString)

	/* se l e n sono > 0, produce su standard output una scala di n gradini
	di lunghezza e altezza l disegnati usando il carattere c (vedi esempi sotto),
	altrimenti non fa niente.*/
	if l <= 0 && n <= 0 {
		return
	}
	spazioPunto := l - 1
	for i := 0; i < n; i++ {
		fmt.Print(DrawSegment(c, i*(l-1), l))
		for j := 0; j < l; j++ {
			fmt.Println(DrawPoint(c, spazioPunto))
		}
		spazioPunto += l - 1
	}
}
