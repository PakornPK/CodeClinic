package main

import (
	"fmt"
	"strconv"
)

// func isCreditCard(number string) bool {
// 	sumEven := 0
// 	for i := 14; i >= 0; i -= 2 {
// 		result := ((int)(number[i]) - 48) * 2
// 		if len(fmt.Sprintf("%d", result)) == 2 {
// 			resultString := fmt.Sprintf("%d", result)
// 			result = ((int)(resultString[0]) - 48) + ((int)(resultString[1]) - 48)
// 		}
// 		sumEven += result
// 	}
// 	sumOdd := 0
// 	for i := 15; i >= 0; i -= 2 {
// 		sumOdd += ((int)(number[i]) - 48)
// 	}
// 	indicator := sumOdd + sumEven
// 	return indicator%10 == 0
// }

func isCreditCard(cn int) bool {
	cnStr := strconv.Itoa(cn)
	Odd, Even := 0, 0
	for i := len(cnStr) - 1; i > 0; i-- {
		if i%2 == 0 {
			ic, err := strconv.Atoi(string(cnStr[i]))
			if err != nil {
				return false
			}
			Even += (ic * 2)
		} else {
			ic, err := strconv.Atoi(string(cnStr[i]))
			if err != nil {
				return false
			}
			Odd += ic
		}
	}
	return (Odd+Even)%10 == 0
}

func main() {
	// fmt.Println(isCreditCard("4388576018410707"))
	fmt.Println(isCreditCard(4388576018410707))
	fmt.Printf("%s\n", "*")
}
