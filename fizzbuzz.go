package fizbuzz

import "fmt"

func fizzbuzz() {
	for i := 0; i < 100; i++ {
		if i%3 == 0 {
			fmt.Printf("Fizz")
		}
		if i%5 == 0 {
			fmt.Printf("Buzz")
		}
		if i%3 == 0 || i%5 == 0 {
			fmt.Printf("\n")
		}
	}
}
