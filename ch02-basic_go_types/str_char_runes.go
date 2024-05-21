package main

import "fmt"

func main() {
	s := "ASASDASF"
	s1 := []byte("ASASDASF")

	fmt.Printf("s: %s\ttype: %T\n", s, s)
	fmt.Printf("s1: %s\ttype: %T", s1, s1)
	fmt.Println()
	r := '€'
	fmt.Printf("This is a rune: %d and the value is: %c", r, r)
}
