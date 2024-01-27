package main

import (
	//"strings"
	//"log"
	"testing"
)

var prog = "./max_somma_cifre"
var diffwidth = 120

func Test0(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"999",
		"\n")
}

func Test1(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"1 999",
		"1\n")
}

func Test2(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"1 3 5 7 21 58 101 999",
		"13\n")
}

func Test3(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"1000 2000 3000 499 500 400 300 999",
		"22\n")
}

func Test4(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"9999 99 9 9999999 1000 99999 999",
		"63\n")
}
