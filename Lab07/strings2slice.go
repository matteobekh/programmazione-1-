/* uso delle funzioni strings.Fields, strings.Split e strings.SplitAfter e delle conversioni []byte() e []rune())
Scrivere una programma strings2slice.go che legge da linea di comando:

- una stringa in formato CSV (del tipo: <nome>,<cognome>,<matricola>), ne estrae i valori salvandoli in una slice
  e stampa la slice;

- della stessa stringa estrae i valori compresa la virgola separatore salvandoli in una slice e stampa la slice;

- una frase tra "" di parole separate anche da più spazi e tab, ne estrae le parole salvandole in una slice e
  stampa la slice;

- un numero intero non negativo e stampa la slice dei codici ASCII delle cifre che lo formano, nello stesso ordine
  da sinistra a destra;

- una stringa di lettere accentate e stampa la slice dei codici unicode delle lettere che la formano;

- un orario nel formato h:m:s, ne estrae ore, minuti e secondi salvandoli in una slice e stampa la slice.
*/

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	riga := os.Args[1]

	valoriCSV := strings.Split(riga, ",")
	fmt.Println(valoriCSV)

	valoriConVirgole := strings.Split(riga, "")
	fmt.Println(valoriConVirgole)

	stringa := os.Args[2]

	//rimuovo le virgolette
	paroleStringa := strings.Trim(stringa, "\"")
	parole := strings.Fields(paroleStringa) //divido le parole
	fmt.Println(parole)

	numeriStr := os.Args[3]
	var cifreAscii []int

	numero, _ := strconv.Atoi(numeriStr)
	for numero > 0 {
		cifra := numero % 10
		numero /= 10
		cifraAscii := int('0' + cifra) // Ottengo così il codice ASCII della cifra
		cifreAscii = append([]int{cifraAscii}, cifreAscii...)
	}

	/* // Estrai i codici Unicode delle lettere accentate
		accentedLetters := "àèìòù"
		var unicodeLetters []int
		for _, letter := range accentedLetters {
			unicodeLetters = append(unicodeLetters, int(letter))
		}
		fmt.Println("Unicode Letters:", unicodeLetters)

		// Estrai ore, minuti e secondi dall'orario
		timeStr := "12:34:56"
		timeValues := strings.Split(timeStr, ":")
		var timeIntValues []int
		for _, value := range timeValues {
			intValue, err := strconv.Atoi(value)
			if err != nil {
				fmt.Println("Errore di conversione:", err)
				os.Exit(1)
			}
			timeIntValues = append(timeIntValues, intValue)
		}
		fmt.Println("Time Values:", timeIntValues)
	}*/ // zero sbatti di finire questo programma, quindi ho copiato questa parte da ChatGPT. Andrebbe solamente modificato inserendo le cose da linea di comando, ma ad occhio sembra ok
}
