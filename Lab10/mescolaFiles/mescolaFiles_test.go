/*
Scrivere un programma `mescolaFiles.go` che legge due file di input (i cui nomi sono passati su linea di comando) e produce in output una "mescola" che prende una parola dal primo file e una dal secondo (scrivendole seguite da un newline ciascuna) fino ad esaurire entrambi i file.
Se i file contengono numeri di parole diversi il programma deve comunque arrivare ad esaurire tutto l'input (usando le parole rimanenti del file che ne contiene di più).

Se mancano parametri a linea di comando il programma termina col messaggio "Inserire DUE nomi di file".

Esempio, dati due file che contengono rispettivamente:
1) uno due tre quattro cinque
2) alpha beta gamma

L'output che viene generato è:
uno
alpha
due
beta
tre
gamma
quattro
cinque

*****************************************************************
>>> NOTA BENE

I test qui sotto funzionano correttamente su Linux e con
l'oracolo fornito, ma dato che l'oracolo è compilato per Linux
tutti i test basati su di esso non funzioneranno (falliranno)
su altre piattaforme.
Ergo sono stati forniti anche dei test basati sui
file ".expected" per chi usa Windows o Mac a casa.
*****************************************************************



*/

package main

import (
	"testing"
)

var prog = "./mescolaFiles"
var diffwidth = 120

func TestMancaParametri(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"NIENTE",
		"Inserire DUE nomi di file\n")
}

func TestMinimo(t *testing.T) {
	LanciaGenericaConFileOutAtteso(
		t,
		prog,
		"NIENTE",
		"mescolaFiles.expected",
		"uno.input", "due.input")
}

func TestMinimoOracolo(t *testing.T) {
	ConfrontaConOracolo(
		t,
		prog,
		"",
		"uno.input", "due.input")
}

func TestUgualiOracolo(t *testing.T) {
	ConfrontaConOracolo(
		t,
		prog,
		"",
		"uno.input", "uno.input")
}

func TestLungoOracolo(t *testing.T) {
	ConfrontaConOracolo(
		t,
		prog,
		"",
		"uno.input", "mescolaFiles_test.go")
}
