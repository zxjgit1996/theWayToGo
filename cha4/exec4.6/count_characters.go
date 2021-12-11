package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	string1 := "asSASA ddd dsjkdsjs dk"
	fmt.Printf("the number of byte in string1 is %d\n", len(string1))
	fmt.Printf("the number of characters in string1 is %d\n", utf8.RuneCountInString(string1))
	string2 := "asSASA ddd dsjkdsjsこん dk"
	fmt.Printf("the number of byte in string2 is %d\n", len(string2))
	fmt.Printf("the number of characters in string2 is %d\n", utf8.RuneCountInString(string2))

}
