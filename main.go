package main

import "fmt"

func main() {
	grades := []int{3, 4, 2, 5, 5, 3}
	fmt.Println(grades[1:4])
	fmt.Println(grades[:3])
	fmt.Println(grades[2:])
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
