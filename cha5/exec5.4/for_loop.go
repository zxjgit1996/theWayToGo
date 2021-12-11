package main

import "fmt"

func main() {
	fmt.Println("for")
	forRange()
	fmt.Println("goto")
	gotoRange()
}

func forRange() {
	for i := 0; i < 15; i++ {
		fmt.Printf("%d\n", i)
	}
}

func gotoRange() {
	i := 0
LOOP:
	if i < 15 {
		fmt.Printf("%d\n", i)
		i++
		goto LOOP
	}
}
