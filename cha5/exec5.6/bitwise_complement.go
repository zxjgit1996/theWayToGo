package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		bitwise(i)
	}

}

func bitwise(i int) {
	fmt.Printf("the complement of %d is %b\n", i, ^i)
}
