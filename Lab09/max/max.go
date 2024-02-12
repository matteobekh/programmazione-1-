/* Scrivere un programma max.go, dotato di:

una funzione ricorsiva recursiveMax(list []int) int che restituisca il massimo
tra i valori di list
una funzione main() che legga da standard input (ctrl D per terminare) una lista
di numeri interi (che posono essere positivi, negativi, nulli)
ed emetta nel flusso d'uscita il massimo tra i numeri letti.

Il massimo di una sequenza vuota non è definito, quindi assumiamo
(non è richiesto che vengano fatti controlli) che la sequenza abbia sempre
almeno un numero.
Può essere comodo definire una funzione di supporto

func greater(m, n int) int
che restituisce il maggiore tra m e n */

package main

import (
  "fmt"
  //"os"
  "io"
)

func recursiveMax(list []int) int {
  if len(list) == 1 {
    return list[0]
  }
  return greater(list[0], recursiveMax(list[1:]))
}

func greater(m, n int) (max int) {
    if m >= n {
       max = m
    } else {
       max = n
    }
    return
}

func main() {
  var lista []int
  var num int

  for {
    _, err := fmt.Scanf("%d",&num)
    if err == io.EOF {
      break
    }
    lista = append(lista, num)
  }
  fmt.Println(recursiveMax(lista),"è il numero più grande delle sequenza di numeri inseriti")
}
