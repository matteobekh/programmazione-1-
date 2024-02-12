/* Scrivere un programma `repInPlace.go` che definisca una funzione:
	func repInPlace(stringa []rune, old rune, new rune)
che modifichi la `stringa` passata sostituendo ogni occorrenza della runa `old`
con la runa `new`

Aggiungere inoltre un main che accetti tre parametri a linea di comando,
rispettivamente:
- stringa da modificare
- carattere sostituendo
- carattere di rimpiazzo */

package main

import (
  "fmt"
  "os"
)

func repInPlace(stringa *[]rune, old rune, new rune) {
  for i, char := range (*stringa) {
    if char == old {
      (*stringa)[i] = new
    }
  }
  return
}

func main() {
  str := []rune(os.Args[1])
  vecchioCarattere := []rune(os.Args[2])[0] // 0 per prendere il primo elemento che hai inserito 
  nuovoCarattere := []rune(os.Args[3])[0]
  //faccio una stampa di ciò che andrò ad utilizzare solo per sicurezza
  fmt.Println(string(str), string(vecchioCarattere), string(nuovoCarattere))
  repInPlace(&str, vecchioCarattere, nuovoCarattere)// sta ad indicare il puntatore

  fmt.Println("stringa modificata: ",string(str))
}
