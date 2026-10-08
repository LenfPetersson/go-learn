package main

import "fmt"

func main() {
	name1 := "Yana"
	name2 := "Sasha"
	greet(name1)
	greet(name2)

	result := square(5)
	fmt.Println(result)
	fmt.Println(square(12))

	sum := 0

	for i := 1; i <= 10; i++ {
		if isEven(i) {
			sum += i
		}
	}
	fmt.Println(sum)
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
