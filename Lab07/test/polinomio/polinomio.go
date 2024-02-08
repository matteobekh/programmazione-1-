/* Scrivere un programma polinomio.go che legge da standard input una riga che contiene dei numeri a, b, ....
Il programma calcola il valore in x del polinomio
a + bx + cx^2 + dx^3 ....
corrispondente alla sequenza dei numeri letti tranne l'ultimo; l'ultimo valore della sequenza è il valore di x.
Ad esempio,
3 2 0 7 5
corrisponde al polinomio 3 + 2x + 7x^3 da valutare per x = 5
*/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	//fmt.Println("inserire numeri del polinomio:")

	scanner := bufio.NewReader(os.Stdin)
	input, _ := scanner.ReadString('\n')
	//fmt.Println("numeri inseriti:", input)
	input = strings.TrimSpace(input)
	numeri := strings.Fields(input)
	//fmt.Println("slice di numeri:", numeri)

	costanteNum, _ := strconv.ParseFloat(numeri[0], 64)
	valoreX, _ := strconv.ParseFloat(numeri[len(numeri)-1], 64)
	//fmt.Println("valore di X", valoreX)
	totale := 0.0
	totale += costanteNum
	//fmt.Println("somma di totale con la costante", totale)
	for i := 1; i <= len(numeri)-2; i++ {
		//fmt.Println("ultimo elemento prima della x", numeri[len(numeri)-2])
		numero, _ := strconv.ParseFloat(numeri[i], 64)
		totale += numero * math.Pow(valoreX, float64(i))
	}

	fmt.Println(int(totale))
	//fmt.Println()

}

/* anche in questo caso non funge il test nonostante i risultati siano corretti. Non capisco bene se il problema sia quindi del pc o se ci sia
qualcosa che non va. In ogni caso, i risultati perlomeno sono corretti*/
