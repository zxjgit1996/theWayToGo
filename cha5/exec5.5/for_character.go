package main

import "fmt"

func main() {
	printCharacter2()
}

func printCharacter1() {
	for i := 1; i < 26; i++ {
		for j := 0; j < i; j++ {
			fmt.Printf("G")
		}
		fmt.Println()
	}
}

func printCharacter2() {
	s := "G"
	for i := 1; i <= 25; i++ {
		fmt.Println(s)
		s += "G"
	}
}
