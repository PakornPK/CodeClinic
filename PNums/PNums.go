package main

import "fmt"

func isPerfact(Num int) []int {
	count := []int{}
	for i := 1; i < Num; i++ {
		if Num%i == 0 {
			count = append(count, i)
		}
	}
	return count
}

func main() {
	t1 := 6
	t2 := 28

	fmt.Printf("t1 = %v\n", isPerfact(t1))
	fmt.Printf("t2 = %v\n", isPerfact(t2))
	fmt.Printf("7 = %v\n", isPerfact(7))
	fmt.Printf("11 = %v\n", isPerfact(11))
	fmt.Printf("496 = %v\n", isPerfact(496))
	fmt.Printf("8128 = %v\n", isPerfact(8128))
}
