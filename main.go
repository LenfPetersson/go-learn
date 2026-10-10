package main

import "fmt"

func main() {
	ages := map[string]int{
		"Kris":   19,
		"Max":    21,
		"Kirill": 20,
	}
	fmt.Println(len(ages))

	delete(ages, "Max")
	fmt.Println(len(ages))
	fmt.Println(ages)

	delete(ages, "Nobody")
	fmt.Println(len(ages))

	for name, age := range ages {
		fmt.Println(name, age)
	}
}

func greet(name string) {
	fmt.Println("Привет,", name)
}

func square(n int) int {
	return n * n
}

func isEven(n int) bool {
	return n%2 == 0
}
