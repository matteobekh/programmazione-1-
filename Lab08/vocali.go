/* Scrivere un programma vocali.go che analizza un testo e conta le occorrenze delle vocali (sia minuscole che maiuscole, ma non le accentate)
nel testo e stampa, ma solo per le vocali presenti nel testo, il numero di volte che le vocali stesse sono presenti nel testo.
In particolare il programma è dotato di:


una funzione

func contaVocali(s string, vocali map[rune]int)


per contare le occorrenze delle diverse vocali (sia minuscole che maiuscole - vedi es sotto) in tutte le stringhe che le vengono passate.
La funzione, data una stringa s e una mappa vocali, aggiorna opportunamente la mappa vocali aggiungendo eventuali vocali
non ancora presenti / incrementandone i valori.
Nota: per individuare le vocali e aggiornare la mappa vocali usate uno switch con un solo case o un if, sempre con un solo caso,
che ovviamente sarà parametrico nelle vocali.

una funzione

func main()

che legge una riga di testo da standard input e produce una mappa tra vocali presenti nel testo e il numero delle loro occorrenze nel testo,
e la stampa.


Esempio
Se l'input è:

jdhkas c'è dkasjhkdjashkdh askdh ksah @@@ €€€ ### Ħ  wi) Ø qwqwe qwyewquteuqwte q 312312 2312wweqe €łłŧŧŧŧŧ sdasdas AA JKJLKLJLKJ LIIIIII  u ù aeiou 	AEIO


L'output è (salvo l'ordine di stampa):

o : 1
E : 1
u : 4
e : 7
A : 3
I : 7
O : 1
a : 8
i : 2

Come si potrebbero stampare le vocali con le loro occorrenze in ordine alfabetico?
Scrivere una seconda versione vocali_bis.go che produce l'output in ordine:

A : 3
E : 1
I : 7
O : 1
a : 8
e : 7
i : 2
o : 1
u : 4
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func contaVocali(s string, vocali map[rune]int) {
	for _, char := range s {
		switch char {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			vocali[char]++
		}
	}
}

func main() {

	vocali := make(map[rune]int)
	scanner := bufio.NewReader(os.Stdin)

	riga, _ := scanner.ReadString('\n')
	contaVocali(riga, vocali)

	/* for k, v := range vocali {
		fmt.Println(string(k), ":", v)
	} */ // con questo semplicemente stampa le vocali in ordine casuale
	chiavi := make([]rune, len(vocali))
	i := 0
	for k := range vocali {
		chiavi[i] = k
		i++
	}

	for _, k := range chiavi {
		fmt.Print("\n", string(k), " ", vocali[k])
	}
}

// preso dalle note del lab 08
/* // stampa dell'intera mappa
	fmt.Println(myMap)

  // stampa elemento per elemento, non importa in che ordine
	for key, value := range myMap {
		fmt.Println(key, value)
	}

  // stampa elemento per elemento, in ordine crescente di chiave, chiavi non presenti (con valore zero) incluse
	for key := min; key <= max; key++ {
		fmt.Println(key, myMap[key])
	}

  // stampa elemento per elemento, in ordine crescente di chiave, chiavi non presenti escluse

	for key := min; key <= max; key++ {
		if value, ok := myMap[key]; ok {
			fmt.Println(key, value)
		}
	}
*/
