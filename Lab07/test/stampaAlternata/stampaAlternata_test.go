/*
	vedi note sul repo
*/

package main

import ( //"log"
	//"os/exec"
	//"strings"
	"testing"
	//"bytes"
	//"io"
	//"os"
)

var prog = "./stampaAlternata"
var diffwidth = 120

func TestMinimaleMain(t *testing.T) {
	LanciaGenerica(t,
		prog,
		"d\np\nd\np\nd\np\n",
		"p\np\np\nd\nd\nd\n",
		"NIENTE")
}
