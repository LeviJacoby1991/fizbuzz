package fizbuzz

import (
	"fmt"
	"hash/fnv"
)

func Fizzbuzz() {
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

func WhatIsTheirFirstProgram(name string) {
	fmt.Println("He makes hello world")
}

func Hash(s string) uint32 {
	return 0
}

func hash(s string) uint32 {
	h := fnv.New32()
	h.Write([]byte(s))
	return h.Sum32()
}
