/*
	Scrivere un programma capoluoghi.go che legge da standard input (usate la ridirezione dal file capoluoghi) una serie di informazioni sui

capoluoghi di provincia organizzate su righe, nel seguente formato:

la prima riga dell'input è l'intestazione
le successive contengono i dati, separati da virgole, un riga per capoluogo, in questo ordine:

	Nome,Sigla,Regione,Popolazione,Superficie,Densità,Altitudine

Nome, Sigla, Regione sono stringhe; Popolazione,Superficie, Densità, Altitudine sono int.

Il programma deve poter memorizzare i dati relativi a Nome, Sigla, Regione, Superficie in modo sia possibile ottenere la stampa di tali
dati relativamente ai capoluoghi di provincia, le cui sigle sono state passate da linea di comando.
Il programma deve essere dotato di una struct Capoluogo con campi Nome, Sigla, Regione, Superficie, in quest'ordine.
NB: nella struct, il campo Sigla è "ridondante in apparenza"

# Esempio di esecuzione

$ ./capoluoghi MI < capoluoghi.csv
{MILANO MI LOM 181}
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Capoluogo struct {
	Nome, Sigla, Regione                         string
	Popolazione, Superficie, Densità, Altitudine int
}

func main() {
	sigla := os.Args[1]

	mappaCapoluogo := map[string]Capoluogo{}

	/* file, err := os.Open("capoluoghi.csv")
	if err != nil {
		fmt.Println("errore!")
		os.Exit(1)
	}

	defer file.Close()*/
	var capoluogo Capoluogo
	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner.Scan()
	for scanner.Scan() {
		riga := scanner.Text()

		elementi := strings.Split(riga, ",")
		/* Il programma deve poter memorizzare i dati relativi a Nome, Sigla, Regione, Superficie in modo sia possibile ottenere la stampa di
		tali dati relativamente ai capoluoghi di provincia, le cui sigle sono state passate da linea di comando.*/

		//popolazione, _ := strconv.Atoi(elementi[3])
		superficie, _ := strconv.Atoi(elementi[4])
		//densità, _ := strconv.Atoi(elementi[5])
		//altitudine, _ := strconv.Atoi(elementi[6])

		//creo un capoluogo che inserirò nella mappa se le sigle combaciano

		capoluogo = Capoluogo{
			Nome:    elementi[0],
			Sigla:   elementi[1],
			Regione: elementi[2],
			//Popolazione: popolazione,
			Superficie: superficie,
			//Densità:     densità,
			//Altitudine:  altitudine,
		}
		//controllo sulle sigle
		if capoluogo.Sigla == sigla {
			mappaCapoluogo[capoluogo.Sigla] = capoluogo
			break
		}

	}
	fmt.Printf("{%s %s %s %d}\n", capoluogo.Nome, capoluogo.Sigla, capoluogo.Regione, capoluogo.Superficie)
}

/* il programma va bene, tuttavia non credo sia questo il modo in cui lo voglia salvare. ma ce lo facciamo andar bene così */
