package main

import "fmt"

func main() {
	var month int
	fmt.Scanln(&month)
	season := Season(month)
	fmt.Println(season)
}

func Season(month int) string {
	switch month {
	case 3, 4, 5:
		return "spring"
	case 6, 7, 8:
		return "summer"
	case 9, 10, 11:
		return "autumn"
	case 12, 1, 2:
		return "winter"
	default:
		return "unknown"
	}
}
