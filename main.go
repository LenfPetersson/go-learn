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

func main() {
	p := Person{Name: "Maxim", Age: 21, City: "Novosibirsk"}
	fmt.Println(p.Name, p.Age)
	p.Age = 22
	fmt.Println(p)

	d := Person{Name: "Kirill"}
	fmt.Println(d.Age, d.City)

	people := []Person{
		{Name: "Kris", Age: 19, City: "Krasnoyarsk"},
		{Name: "Artem", Age: 22, City: "Taldikurgan"},
		{Name: "Danill", Age: 21, City: "Kemerovo"},
		{Name: "Petr", Age: 16, City: "Abakan"},
	}
	maxAgePeople := oldest(people)
	fmt.Println(maxAgePeople)
	fmt.Println(maxAgePeople.Name)

	for _, person := range people {
		fmt.Println(person.Name, "из", person.City)
	}

	for _, person := range people {
		if isAdult(person) {
			fmt.Println(person.Name)
		}
	}
}

func isAdult(p Person) bool {
	if p.Age >= 18 {
		return true
	} else {
		return false
	}
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

type Person struct {
	Name string
	Age  int
	City string
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
