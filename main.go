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
func minMax(nums []int) (int, int) {
	resultMin := nums[0]
	resultMax := nums[0]

	for _, v := range nums {
		if v > resultMax {
			resultMax = v
		}
		if v < resultMin {
			resultMin = v
		}
	}
	return resultMin, resultMax
}

type Person struct {
	Name string
	Age  int
	City string
}

func (p Person) IsAdult() bool {
	if p.Age >= 18 {
		return true
	} else {
		return false
	}
}

func main() {
	people := []Person{
		{Name: "Kris", Age: 17, City: "Krasnoyarsk"},
		{Name: "Artem", Age: 25, City: "Taldikurgan"},
	}

	for _, v := range people {

		fmt.Println(v.Name, ":", v.IsAdult())

	}

	tovari := []Product{
		{Name: "Pizza", Price: 133.2, InSock: true},
		{Name: "Banana", Price: 99, InSock: true},
		{Name: "Apple", Price: 122.8, InSock: false},
		{Name: "Potato", Price: 14, InSock: false},
	}
	count := 0.0

	for _, v := range tovari {
		if v.InSock {
			count += v.Price
		}
	}
	fmt.Println(count)
}

type Product struct {
	Name   string
	Price  float64
	InSock bool
}

func oldest(people []Person) Person {
	best := people[0]
	for _, person := range people {
		if person.Age > best.Age {
			best = person
		}
	}
	return best
}

func find(book map[string]string, name string) (string, bool) {
	if v, ok := book[name]; ok {
		return v, true
	} else {
		return "", false
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
