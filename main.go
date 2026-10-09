package main

import "fmt"

func main() {

	ages := map[string]int{
		"Kris":   19,
		"Max":    21,
		"Kirill": 20,
	}
	fmt.Println(ages["Kris"])
	fmt.Println(len(ages))

	ages["Tanya"] = 20
	fmt.Println(len(ages))

	ages["Max"] = 19

	fmt.Println(ages["Max"])

	fmt.Println(ages["nobody"])

	age, ok := ages["nobody"]
	if ok {
		fmt.Println("vozrast: ", age)
	} else {
		fmt.Println("net takogo")
	}

	words := []string{"go", "is", "fun", "go", "go", "is"}
	counts := map[string]int{}
	for _, word := range words {
		counts[word]++
	}
	fmt.Println(counts)

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
