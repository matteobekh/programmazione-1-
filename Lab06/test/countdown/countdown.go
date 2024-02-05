/*
Scrivere un programma countdown.go dotato di:

- una struttura `Clock` con tre campi `hour`, `min`, `sec` di tipo `int`, dichiarati in quest'ordine

- una funzione `countdown` che fa scalare l'orario di un secondo, invocando opportunamente `downMin` (vedi sotto)
quando sono da modificare anche i minuti

- una funzione `downMin` che fa scalare l'orario di un minuto, invocando opportunamente `downHour` (vedi sotto)
quando sono da modificare anche le ore

- una funzione `downHour` che fa scalare l'orario di un'ora

Stabilite voi la segnatura adeguata per le funzioni qui sopra.

La funzione `main` chiede l'orario di partenza nel formato:
ore minuti secondi
e fa partire il countdown, stampando l'orario a ogni secondo, fino a raggiungere 0 0 0.

Nota. Il programma deve creare **un solo Clock** e aggiornarne via via i valori dei campi,
NON deve creare un nuovo Clock a ogni secondo.

**Suggerimento**: usare l'istruzione `time.Sleep(time.Duration(1) * time.Second)` (potete copiarla così come è)
per far passare (circa) un secondo prima di stampare ogni nuovo orario.

Esempio
$ go run countdown.go
time (hh mm ss): 1 0 3
{1 0 2}
{1 0 1}
{1 0 0}
{0 59 59}
{0 59 58}
{0 59 57}
{0 59 56}
{0 59 55}
...
{0 0 0}
*/
package main

import (
	"fmt"
	"time"
)

type Clock struct {
	hour int
	min  int
	sec  int
}

func countdown(orol *Clock) {
	if orol.sec > 0 {
		orol.sec--
	} else {
		// Se i secondi sono già 0, decrementa un minuto
		downMin(orol)
		orol.sec = 59
	}
}

func downMin(orol *Clock) {
	if orol.min > 0 {
		orol.min--
	} else {
		// Se i minuti sono già 0, decrementa un'ora
		downHour(orol)
		orol.min = 59
	}
}

func downHour(orol *Clock) {
	orol.hour--
}

func main() {
	var orologio Clock

	fmt.Println("time(hh mm ss):")
	fmt.Scan(&orologio.hour, &orologio.min, &orologio.sec)

	for orologio.hour >= 0 && orologio.min >= 0 && orologio.sec >= 0 {
		fmt.Println(orologio)
		countdown(&orologio)
		time.Sleep(time.Duration(1) * time.Second)
	}
	fmt.Println("NIENTE")
}
