/*
draw
------

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
	//"strings"
	//"log"
	"testing"
)

func TestDrawPointNorm(t *testing.T) {
	if DrawPoint('b', 3) != "   b" {
		t.Fail()
	}
}

func TestDrawPointZero(t *testing.T) {
	if DrawPoint('g', 0) != "g" {
		t.Fail()
	}
}

func TestDrawPointLong(t *testing.T) {
	if DrawPoint('M', 30) != "                              M" {
		t.Fail()
	}
}

func TestDrawSegmentNorm(t *testing.T) {
	if DrawSegment('p', 5, 3) != "     ppp" {
		t.Fail()
	}
}

var prog = "./draw"
var diffwidth = 120

func TestMain43o(t *testing.T) {
	ConfrontaConOracolo(
		t,
		prog,
		"",
		"4", "3", "o")
}

func TestMain14_3_cancelletto(t *testing.T) {
	ConfrontaConOracolo(
		t,
		prog,
		"",
		"14", "3", "#")
}
