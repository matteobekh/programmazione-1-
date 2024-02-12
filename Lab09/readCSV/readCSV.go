/*
Scrivere un programma `readCSV.go` che legga da stdin una sequenza di dati in formato CSV
(Comma Separated Values) così organizzati:

	timestamp,temperatura,umidità

Dove:
- `timestamp` è una stringa nel formato YYYY-MM-DD/HH:mm (anno mese giorno ora minuto)
- `temperatura` è un float
- `umidità` è un float
- il separatore è una virgola ","

Il programma deve usare Scanf per leggere ogni riga dell'input.

Il programma deve calcolare il massimo e il minimo dei valori di temperatura e
umidità e stamparli con il timestamp (in formato diverso dall'originale,
deve essere nella forma "X(°/%) misura delle ore HH, minuti mm del giorno DD del mese MM dell'anno YYYY")
in cui sono stati misurati.

Ad esempio, il file fornito come input genera questo risultato:

minTemp: 1.0° misura delle ore 4, minuti 22 del giorno 11 del mese 12 dell'anno 2022
maxTemp: 15.0° misura delle ore 9, minuti 22 del giorno 11 del mese 12 dell'anno 2022
minHumid: 49.0% misura delle ore 9, minuti 24 del giorno 11 del mese 12 dell'anno 2022
maxHumid: 91.0% misura delle ore 9, minuti 31 del giorno 6 del mese 12 dell'anno 2022
*/

package main

import (
  "fmt"
  "io"
)

type Rileva struct {
  aa, mm, gg, hh, min int
  temp, umid float64
}

func main() {
  var rilev, tempMin, tempMax, umidMin, umidMax Rileva

  _, err := fmt.Scanf("%d-%d-%d/%d:%d,%f,%f",&rilev.aa,&rilev.mm,&rilev.gg,&rilev.hh,&rilev.min,&rilev.temp,&rilev.umid)

  if err != io.EOF {
    tempMin = rilev
    tempMax = rilev
    umidMin = rilev
    umidMax = rilev
  }
  for {
    _, err := fmt.Scanf("%d-%d-%d/%d:%d,%f,%f",&rilev.aa,&rilev.mm,&rilev.gg,&rilev.hh,&rilev.min,&rilev.temp,&rilev.umid)

    if err == io.EOF {
      break
    }
    if rilev.temp > tempMax.temp {
      tempMax = rilev
    } else if rilev.temp < tempMin.temp {
      tempMin = rilev
    }
    if rilev.umid > umidMax.umid {
      umidMax = rilev
    } else if rilev.umid < umidMin.umid {
      umidMin = rilev
    }
  }
  fmt.Printf("minTemp: %.1f° misura delle ore %d, minuti %d del giorno %d del mese %d dell'anno %d\n", tempMin.temp, tempMin.hh, tempMin.min, tempMin.gg, tempMin.mm, tempMin.aa)

	fmt.Printf("maxTemp: %.1f° misura delle ore %d, minuti %d del giorno %d del mese %d dell'anno %d\n", tempMax.temp, tempMax.hh, tempMax.min, tempMax.gg, tempMax.mm, tempMax.aa)

	fmt.Printf("minHumid: %.1f%% misura delle ore %d, minuti %d del giorno %d del mese %d dell'anno %d\n", umidMin.umid, umidMin.hh, umidMin.min, umidMin.gg, umidMin.mm, umidMin.aa)

	fmt.Printf("maxHumid: %.1f%% misura delle ore %d, minuti %d del giorno %d del mese %d dell'anno %d\n", umidMax.umid, umidMax.hh, umidMax.min, umidMax.gg, umidMax.mm, umidMax.aa)
}
/* questo programma sfrutta lo stesso concetto di quando devi trovare massimi, minimi eccetera di diversi valori. Anche in questo caso
parti mettendo i valori iniziali come massimi, minimi eccetera. Poi fai i confronti: se trovi un nuovo massimo lo sostituisci
stessa cosa per i minimi. Unica cosa qui è che vengono modificate intere strutture e non solo alcuni valori. Questo funziona anche perché
alcuni valori rimanevano invariati quindi modificando tutta la struttura non avrebbe fatto nulla. diverso
sarebbe successo se ogni parte dovesse essere modificata (anche se dovrebbe funzionare in ogni caso) */ 
