/* Scrivere un programma (il cui file deve chiamarsi 'vini.go') dotato di:

  - una struttura BottigliaVino con i seguenti campi (dichiarati nell'ordine):
    nome string
    anno int
    gradi float32
    cl int

  - una funzione
    CreaBottiglia(nome string, anno int, gradi float32, cl int) (BottigliaVino, bool)
    che, se i parametri sono **validi** (campi string diversi da vuoto, campi int e float maggiori di zero) ù
	crea una bottiglia corrispondente a questi dati e la restituisce insieme al valore 'true',  altrimenti restituisce una bottiglia "zero-value" e 'false'.

  - una funzione
    CreaBottigliaDaRiga(riga string) (BottigliaVino, bool)
    che crea una bottiglia a partire dalla sua rappresentazione sotto forma di riga di testo CSV (vedi esempio nelle specifiche del main);
    se i parametri ci sono tutti e sono validi (vedi sopra), crea una bottiglia corrispondente a questi dati e la restituisce insieme al valore 'true',
	altrimenti restituisce una bottiglia "zero-value" e 'false'.
    Non sono richiesti controlli sui tipi dei dati: si può assumere che i dati, se ci sono, siano nel formato corretto (ma i valori vanno controllati).

  - un **metodo** per BottigliaVino
    String() string
    che restituisce una riga di descrizione della bottiglia nel seguente formato:  <nome>, <anno>, <gradi>°, <cl>cl
    (cioè ad es. "Rasol, 2018, 14°, 750cl" per la prima riga dell'esempio sopra).
    Suggerimento: i "format verb" %g e %v formattano i float omettendo il punto decimale quando non ci sono cifre dopo la virgola

- una funzione main() che legge da un file (il cui nome è passato da linea di comando) delle righe che contengono ognuna i dati relativi ad una bottiglia di una cantina,
separati da virgole, nel formato del seguente esempio (nome,anno,gradi,cl):

Rasol,2018,14,750
Camunnorum,2015,15,750
Dom Perignon,2019,12.5,1500
Balon,2013,15,750
Verdicchio,2020,11,375

e stampa su stdout:

  - l'elenco delle bottiglie della cantina (esattamente nello stesso formato rappresentato qui sopra).
    Attenzione alle righe vuote (vedere vini.input), il programma deve essere "robusto" e ignorarle.

  - il numero di bottiglie nella cantina

  - la bottiglia con gradazione massima

  - la bottiglia più vecchia

  - i cl totali di vino della cantina

Esempio di esecuzione con input vini.input
---------------------

Rasol, 2018, 14°, 750cl
Camunnorum, 2015, 15°, 750cl
Dom Perignon, 2019, 12.5°, 1500cl
Balon, 2013, 15°, 750cl
Verdicchio, 2020, 11°, 375cl
Rasol, 2018, 14°, 1000cl
Verdicchio, 2020, 11°, 375cl
n. bottiglie: 7
bottiglia con grado max: Camunnorum, 2015, 15°, 750cl
bottiglia più vecchia: Balon, 2013, 15°, 750cl
tot vino: 5500 cl*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type BottigliaVino struct {
	nome  string
	anno  int
	gradi float32
	cl    int
}

func CreaBottiglia(nome string, anno int, gradi float32, cl int) (BottigliaVino, bool) {
	var bott BottigliaVino
	var ok bool

	if len(nome) == 0 {
		return bott, ok
	}
	if anno < 0 {
		return bott, ok
	}
	if gradi < 0 {
		return bott, ok
	}
	if cl < 0 {
		return bott, ok
	}
	ok = true

	bott.nome = nome
	bott.anno = anno
	bott.gradi = gradi
	bott.cl = cl

	return bott, ok
}

func CreaBottigliaDaRiga(riga string) (BottigliaVino, bool) {
	var bott BottigliaVino
	var ok bool

	if len(riga) == 0 {
		return bott, ok
	}

	sliceBotSingola := strings.Split(riga, ",")

	anno, err := strconv.Atoi(sliceBotSingola[1])
	if err != nil {
		return bott, ok
	}
	gradi, err := strconv.ParseFloat(sliceBotSingola[2], 32)
	if err != nil {
		return bott, ok
	}
	cl, err := strconv.Atoi(sliceBotSingola[3])
	if err != nil {
		return bott, ok
	}

	bott, ok = CreaBottiglia(sliceBotSingola[0], anno, float32(gradi), cl)

	return bott, ok
}

/*
un **metodo** per BottigliaVino

	String() string
	che restituisce una riga di descrizione della bottiglia nel seguente formato:  <nome>, <anno>, <gradi>°, <cl>cl
	(cioè ad es. "Rasol, 2018, 14°, 750cl" per la prima riga dell'esempio sopra).
	Suggerimento: i "format verb" %g e %v formattano i float omettendo il punto decimale quando non ci sono cifre dopo la virgola
*/
func (bottiglia *BottigliaVino) String() string { //forse non restituisce stringa...
	return fmt.Sprintf("%s, %d, %v°, %dcl", bottiglia.nome, bottiglia.anno, bottiglia.gradi, bottiglia.cl)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("manca nome file")
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println("errore in apertura file")
		os.Exit(2)
	}

	scanner := bufio.NewScanner(file)
	var totVino int
	var numBott int
	var bottGradoMax, bottVecch BottigliaVino

	//var sliceBottiglie []BottigliaVino
	for scanner.Scan() {
		riga := scanner.Text()
		bottiglia, ok := CreaBottigliaDaRiga(riga)
		if ok {
			fmt.Println(bottiglia.String())
			numBott++
			totVino += bottiglia.cl
			if bottiglia.gradi > bottGradoMax.gradi {
				bottGradoMax = bottiglia
			}
			if bottiglia.anno < bottVecch.anno && bottVecch.anno == 0 {
				bottVecch = bottiglia
			}
		}
	}
	defer file.Close()

	fmt.Println("n bottiglie:", numBott)
	fmt.Println("bottiglia con grado max:", bottGradoMax.String())
	fmt.Println("bottiglia più vecchia:", bottVecch.String())
	fmt.Println("tot vino:", totVino)
}

/* per bottiglia max e vecchia è molto + comoda una slice di bottiglie che ti crei
, metti nella bottiglia + vecchia il primo elemento e poi fai un for range della slice alla
ricerca di ciò che ti serve
inoltre per returnare il zero value pupi anche fare, in questo caso:
return BottigliaVino{}, false*/
