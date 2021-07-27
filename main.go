package main

import "fmt"

func fizzBuzz(n int) []string {
	res := make([]string, n)
	for i := 1; i < n+1; i++ {
		if (i%3 == 0) && (i%5 == 0) {
			res[i-1] = "FizzBuzz"
			continue
		}

		if i%3 == 0 {
			res[i-1] = "Fizz"
			continue
		}

		if i%5 == 0 {
			res[i-1] = "Buzz"
			continue
		}

		res[i-1] = fmt.Sprintf("%v", i)
	}
	return res
}
