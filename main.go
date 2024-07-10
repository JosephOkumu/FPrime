package main

import (
	"fmt"
	"os"
	"strconv"
)

func isPrime(nb int) bool {
	if nb <= 1 {
		return false
	}else if nb == 2 || nb == 3 {
		return true
	}
	for i := 5; i*i <= nb; i+= 6 {
		if i%2 == 0 || i%3 == 0 {
			return false
		}
	}
	return true
}

func fPrime(nb int) string {
	slice := []int{}
	for i := 1; i <= nb; i++ {
		if isPrime(i) {
			for nb%i == 0 {
				slice = append(slice, i)
				nb/=i
			}
		}
	}
	
	result := ""
	for _, v := range slice {
		if result != "" {
			result += "*"
		}
		result += strconv.Itoa(v)
	}
	return result
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println(0)
		return
	}
	args := os.Args[1]

	num, err := strconv.Atoi(args)
	if err != nil {
		return
	}
	output := fPrime(num)
	fmt.Println(output)
}
