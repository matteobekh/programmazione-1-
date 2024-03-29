/*
In un file briscola.go si scriva una funzione

	punti(s string) int

che accetta una stringa di 3 caratteri che rappresenta una mano di tre carte e restituisce il punteggio complessivo relativo per il gioco della briscola.
Ad esempio per la mano "53J" restituisce 12 (10 della carta 3 + 2 del fante).
Se la stringa contiene un simbolo che non corrisponde a nessuna carta, la funzione restituisce -1.

Si scriva un main per invocare e testare la funzione. Il programma legge da standard input una stringa e controlla che sia di tre caratteri. Poi stampa
mano <mano>: <tot> punti

Punti a briscola:
Asso (A): 11
3: 10
Re (K): 4
Donna (Q): 3
Fante (J): 2
7, 6, 5, 4, 2: 0

nomefile: briscola.go
*/

package main

import (
	//"fmt"
	//"log"
	//"os/exec"

	//"strings"
	"testing"
	//"bytes"
	//"io"
	//"os"
)

var prog = "./briscola"
var diffwidth = 120

func TestEsempio(t *testing.T) {
	if punti("53J") != 12 {
		t.Fail()
	}
}
func TestMainEsempio(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"53J\n",
		"mano: punti = 12\n")
}

func TestUna(t *testing.T) {
	if punti("K") != 4 {
		t.Fail()
	}
}
func TestMainUna(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"K\n",
		"mano: errore nell'input, le carte devono essere 3\npunti = 4\n")
}

func TestFigure(t *testing.T) {
	if punti("QQJ") != 8 {
		t.Fail()
	}
}

func TestMainFigure(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"QQJ\n",
		"mano: punti = 8\n")
}
