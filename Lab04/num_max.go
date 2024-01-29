/* Scrivere un programma num_max.go che legge una sequenza di 10 interi positivi
e stampa il massimo intero letto e quante volte tale massimo compare nella sequenza.
E` possibile risolvere il problema senza tenere in memoria la sequenza di interi.
Che tipo di composizione occorrerà per mettere insieme il calcolo del massimo e il conteggio
delle occorrenze di tale massimo?
*/

package main

import "fmt"

func main() {
	var num int

	fmt.Scan(&num)
	max := num
	count := 1
	for i := 1; i < 10; i++ {
		fmt.Scan(&num)

		if num == max {
			count++
		} else if num > max {
			max = num
			count = 1
		}

	}
	fmt.Println("massimo:", max)
	fmt.Println("count:", count)
}
