package main

import "github.com/01-edu/z01"

/*
Write a program that displays the alphabet,
with even letters in uppercase, and odd letters in lowercase,
followed by a newline ('\n').
*/
func main() {
	for i := 'a'; i < 'z'; i++ {
		if i%2 == 0 {
			z01.PrintRune(i - 32)
		} else {
			z01.PrintRune(i)
		}
	}
	z01.PrintRune(10)
}
