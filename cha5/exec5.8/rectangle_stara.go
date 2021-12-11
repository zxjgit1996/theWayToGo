package main

import "fmt"

func main() {
	printfRectangle()
}

func printfRectangle() {
	w, h := 20, 10
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
