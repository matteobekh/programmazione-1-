/* Scrivere un programma contaLettere.go che legge un testo (da stdin) e stampa quante volte (anche 0) appare nel testo
ciascuna lettera minuscola dell'alfabeto ('a'-'z').

Il programma è dotato di una costante
const LEN_ALFABETO = 26
e di una funzione
func aggiornaConteggi(s string, contaMinu *[LEN_ALFABETO]int)
che, data una stringa s, aggiorna i conteggi delle lettere minuscole di s.
Si consiglia di usare la ridirezione dell'input per provare il programma.
Ad esempio (attenzione che il risultato potrebbe cambiare in funzione del contenuto del file, cioè i numeri sono esemplificativi):

$ ./contaLettere < questeSpecifiche
a 35
b 1
c 15
d 11
e 38
f 3
g 18
h 4
i 27
j 1
k 1
l 15
m 10
n 29
o 24
p 11
q 2
r 23
s 16
t 28
u 14
v 4
w 1
x 1
y 1
z 3
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

const LEN_ALFABETO = 26

func aggiornaConteggi(s string, contaMinu *[LEN_ALFABETO]int) {

	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			lettera := char - 'a'
			(*contaMinu)[lettera]++
		}
	}
}

func main() {
	contaMinu := [LEN_ALFABETO]int{}

	scanner := bufio.NewScanner(os.Stdin)
	//fmt.Println("testo:")
	for scanner.Scan() {
		riga := scanner.Text()
		aggiornaConteggi(riga, &contaMinu)
	}

	for i := 0; i < LEN_ALFABETO; i++ {
		char := 'a' + i
		count := contaMinu[i]
		fmt.Printf("%c: %d\n", char, count)
	}
}

/* mi sono scordato di farla con la redirezione. Con quello probabilmente si risolve il problema con l'ultimo test */
