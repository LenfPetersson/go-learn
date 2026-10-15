package main

import "fmt"

func grade(score int) string {
	if score >= 90 {
		return "отлично"
	} else if score >= 75 && score <= 89 {
		return "хорошо"
	} else if score >= 50 && score <= 74 {
		return "удовлетворительно"
	} else {
		return "неудовлетворительно"
	}

}
func divide(a, b int) (int, bool) {
	if b != 0 {
		return a / b, true
	} else {
		return 0, false
	}
}
func positives(nums []int) []int {
	result := []int{}
	for _, v := range nums {
		if v > 0 {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	fmt.Println(positives([]int{3, -1, 0, 7, -5, 2}))

	result, ok := divide(10, 2)
	if ok {
		fmt.Println(result)
	} else {
		fmt.Println("на 0 делить нельзя")
	}
	result1, ok1 := divide(10, 0)
	if ok1 {
		fmt.Println(result1)
	} else {
		fmt.Println("на 0 делить нельзя")
	}

	srez := []int{4, 7, 2, 9, 10, 3}

	for _, v := range srez {
		if v%2 == 0 {
			fmt.Println(v)
		}
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
