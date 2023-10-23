/*
Program polja prikazuje kako definiramo, inicializiramo in dostopamo do polj v programskem jeziku go
*/
package main

import "fmt"

func main() {
	// Definicija in izpis polja z dvema elementoma
	// Polje ima vedno definirano velikost
	var a [2]string
	a[0] = "Porazdeljeni"
	a[1] = "Sistemi"
	fmt.Printf("%T, %s, %s\n", a, a[0], a[1])
	fmt.Println(a)

	// Uporabimo lahko tudi kratko notacijo
	fibonacci := [6]int{1, 1, 2, 3, 5, 8}
	fmt.Println(fibonacci)

	// Večdimenzionalna polja
	magic := [3][3]int{{2, 7, 6}, {9, 5, 1}, {4, 3, 8}}
	fmt.Println(magic)
}
