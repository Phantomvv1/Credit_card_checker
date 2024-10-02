package main

import (
	"fmt"
)

func sumDigits(number int) int {
	if number <= 9 {
		return number
	} else {
		return number/10 + number%10
	}
}

func luhnAlogorithm(number int) bool {
	checkDigit := number % 10
	number = number / 10
	var arrToSum []int
	for i := 0; number > 0; i++ {
		if i%2 == 1 {
			arrToSum = append(arrToSum, number%10)
		} else {
			arrToSum = append(arrToSum, (number%10)*2)
		}
		number = number / 10
	}
	for i, number := range arrToSum {
		arrToSum[i] = sumDigits(number)
	}
	var result int
	for _, num := range arrToSum {
		result = result + num
	}
	result = 10 - result%10
	return result == checkDigit
}

func main() {
	fmt.Println(luhnAlogorithm(17893729974))
	fmt.Println("In order to del the master branch")
}
