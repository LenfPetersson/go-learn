package main

import "fmt"

func main() {
	grades := []int{3, 4, 2, 5, 5}
	fmt.Println(grades[0])
	fmt.Println(grades[len(grades)-1])

	fmt.Println(len(grades))
	grades = append(grades, 3)
	fmt.Println(len(grades))
	fmt.Println(grades[len(grades)-1])

	sum := 0
	for _, value := range grades {
		sum += value
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
