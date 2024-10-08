package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Result struct {
	Value bool `json:"resultValue"`
}

func validate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	var result Result
	credCardNum := r.URL.Query().Get("number")
	num, err := strconv.Atoi(credCardNum)
	if err != nil {
		log.Fatal(err)
	}
	result.Value = luhnAlogorithm(num)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(result)
}

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
	mux := http.NewServeMux()
	mux.HandleFunc("/validator", validate)

	err := http.ListenAndServe(":42069", mux)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("Couldn't open server")
	} else if err != nil {
		log.Fatal(err)
	}
}
