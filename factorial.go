package main

import "fmt"

func factorial(nums int) int {
	if nums == 0 {
		return 1
	}
	return nums * factorial(nums-1)
}

func main() {
	for i := 0; i < 20; i++ {
		case1 := factorial(i)
		fmt.Printf("%d! = %v\n", i,case1)
	}
}
