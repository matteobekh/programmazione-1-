/*
Scrivere un programma (il cui file deve chiamarsi 'vini.go') dotato di:

  - una struttura BottigliaVino con i seguenti campi (dichiarati nell'ordine):
    nome string
    anno int
    grado float32
    quantita int

  - una funzione
    CreaBottiglia(nome string, anno int, grado float32, quantita int) (BottigliaVino, bool)
    che, se i parametri sono **validi** (campi string diversi da vuoto, campi int e float maggiori di zero) crea una bottiglia corrispondente a
    questi dati e lo restituisce insieme al valore 'true',  altrimenti restituisce una bottiglia "zero-value" e 'false'.

  - una funzione main() che crea una slice di bottiglie leggendo da un file (il cui nome è passato da linea di comando) delle righe che
    contengono ognuna i dati (separati da virgole) relativi ad una bottiglia, ad es:

    Rasol,2018,14,750
    Camunnorum,2015,15,750
    Dom Perignon,2019,12.5,1500
    Balon,2013,15,750
    Verdicchio,2020,11,375

    e stampa su stdout l'elenco delle bottiglie (esattamente nello stesso formato rappresentato qui sopra).
    Attenzione alle righe vuote (vedere vini.input), il programma deve essere "robusto" e ignorarle.

  - una funzione
    CreaBottigliaDaRiga(riga string) (BottigliaVino, bool)
    che crea una bottiglia a partire dalla sua rappresentazione sotto forma di riga di testo, come da formato specificato sopra; se i parametri ci sono tutti e sono validi (vedi sopra), crea una bottiglia corrispondente a questi dati e lo restituisce insieme al valore 'true',  altrimenti restituisce una bottiglia "zero-value" e 'false'.
    Non sono richiesti controlli sui tipi dei dati: si può assumere che i dati, se ci sono, siano nel formato corretto (ma i valori vanno controllati).

  - una funzione
    AggiungiBottiglia(bot BottigliaVino, cantina *[]BottigliaVino)
    che aggiunge una bottiglia alla lista

  - una funzione
    AggiungiBottigliaDaRiga(bot string, cantina *[]BottigliaVino)
    che effettua la stessa operazione di AggiungiBottiglia ma partendo da una riga di testo (attenzione che è string, nel formato specificato sopra, una riga)

  - una funzione
    ContaPerTipo(nome string, cantina []BottigliaVino) int
    che conta quante bottiglie dello stesso tipo (nome) di vino sono presenti nella cantina

  - un **metodo** (da applicare a BottigliaVino)
    String() string
    che restituisce una riga di descrizione della bottiglia nel seguente formato:  nome,anno,grado,quantità
    (cioè ad es. restituisca "Rasol,2018,14,750" per la prima riga dell'esempio sopra
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type BottigliaVino struct {
	nome     string
	anno     int
	grado    float32
	quantita int
}

/*
	 CreaBottiglia(nome string, anno int, grado float32, quantita int) (BottigliaVino, bool)
	    che, se i parametri sono **validi** (campi string diversi da vuoto, campi int e float maggiori di zero) crea una bottiglia corrispondente a
		questi dati e lo restituisce insieme al valore 'true',  altrimenti restituisce una bottiglia "zero-value" e 'false'.
*/
func CreaBottiglia(nome string, anno int, grado float32, quantita int) (BottigliaVino, bool) {
	if nome == "" || anno <= 0 || grado <= 0 || quantita <= 0 {
		return BottigliaVino{nome: "", anno: 0, quantita: 0, grado: 0}, false
	}
	return BottigliaVino{nome: nome, anno: anno, grado: grado, quantita: quantita}, true
}

/*
	 CreaBottigliaDaRiga(riga string) (BottigliaVino, bool)
	    che crea una bottiglia a partire dalla sua rappresentazione sotto forma di riga di testo, come da formato specificato sopra; se i parametri ci
		sono tutti e sono validi (vedi sopra), crea una bottiglia corrispondente a questi dati e lo restituisce insieme al valore 'true',
		altrimenti restituisce una bottiglia "zero-value" e 'false'.
	    Non sono richiesti controlli sui tipi dei dati: si può assumere che i dati, se ci sono, siano nel formato corretto
		(ma i valori vanno controllati).
*/
func CreaBottigliaDaRiga(riga string) (BottigliaVino, bool) {
	elementi := strings.Fields(riga)

	if len(elementi) != 4 || riga == "" { //se non ho esattamente 4 elementi o se la stringa è vuota
		return BottigliaVino{nome: "", anno: 0, grado: 0, quantita: 0}, false
	}

	nom := elementi[0]
	ann, _ := strconv.Atoi(elementi[1])
	grad, _ := strconv.ParseFloat(elementi[2], 32)
	quantit, _ := strconv.Atoi(elementi[3])
	return CreaBottiglia(nom, ann, float32(grad), quantit)
}

/*
AggiungiBottiglia(bot BottigliaVino, cantina *[]BottigliaVino)

	che aggiunge una bottiglia alla lista
*/
func AggiungiBottiglia(bot BottigliaVino, cantina *[]BottigliaVino) {
	*cantina = append(*cantina, bot)
}

/*
AggiungiBottigliaDaRiga(bot string, cantina *[]BottigliaVino)

	che effettua la stessa operazione di AggiungiBottiglia ma partendo da una riga di testo
	(attenzione che è string, nel formato specificato sopra, una riga)
*/
func AggiungiBottigliaDaRiga(bot string, cantina *[]BottigliaVino) {
	/* prendo bot, la divido in fields, inserisco tutto in una variabile di tipo Bottigliavino e faccio append*/
	bottiglia := strings.Fields(bot)

	nom := bottiglia[0]
	ann, _ := strconv.Atoi(bottiglia[1])
	grad, _ := strconv.ParseFloat(bottiglia[2], 32)
	quantit, _ := strconv.Atoi(bottiglia[3])

	botti, ok := CreaBottiglia(nom, ann, float32(grad), quantit)
	if ok {
		*cantina = append(*cantina, botti)
	}
}

/*
ContaPerTipo(nome string, cantina []BottigliaVino) int

	che conta quante bottiglie dello stesso tipo (nome) di vino sono presenti nella cantina
*/
func ContaPerTipo(nome string, cantina []BottigliaVino) int {
	counter := 1
	for i := 0; i < len(cantina); i++ {
		elemento := cantina[i]
		if elemento.nome == nome {
			counter++
		}
	}
	return counter
}

func (b *BottigliaVino) String() string {
	grado := strconv.FormatFloat(float64(b.grado), 'f', -1, 32)
	return fmt.Sprintf("%s,%d,%s,%d", b.nome, b.anno, grado, b.quantita)
}

/*
una funzione main() che crea una slice di bottiglie leggendo da un file (il cui nome è passato da linea di comando) delle righe che

	contengono ognuna i dati (separati da virgole) relativi ad una bottiglia, ad es:

	Rasol,2018,14,750
	Camunnorum,2015,15,750
	Dom Perignon,2019,12.5,1500
	Balon,2013,15,750
	Verdicchio,2020,11,375

	e stampa su stdout l'elenco delle bottiglie (esattamente nello stesso formato rappresentato qui sopra).
	Attenzione alle righe vuote (vedere vini.input), il programma deve essere "robusto" e ignorarle.
*/
func main() {
	bottiglie := []BottigliaVino{}

	if len(os.Args) != 2 {
		fmt.Println("MANCA NOME FILE")
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println("errore nell'apertura")
		os.Exit(2)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		AggiungiBottigliaDaRiga(scanner.Text(), &bottiglie)
	}

	for _, b := range bottiglie {
		fmt.Println(b.String())
	}
}
