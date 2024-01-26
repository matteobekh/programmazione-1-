package main

import (
	//"strings"
	//"log"
	"testing"
)

var prog = "./disegna_v"
var diffwidth = 120

func Test1(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"1",
		"*\n")
}

func Test2(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"2",
		"* *\n *\n")
}

func Test3(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"3",
		"*   *\n * *\n  *\n")
}

func Test4(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"4",
		"*     *\n *   *\n  * *\n   *\n")
}

func Test10(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"10",
		"*                 *\n *               *\n  *             *\n   *           *\n    *         *\n     *       *\n      *     *\n       *   *\n        * *\n         *\n")
}
