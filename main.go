package main

import "fmt"

func main() {
	age := 21
	if age >= 18 {
		fmt.Println("можно голосовать")
	} else {
		fmt.Println("ещё рано")
	}
	temperature := 25
	if temperature < 10 {
		fmt.Println("холодно")
	} else if temperature <= 25 {
		fmt.Println("нормально")
	} else {
		fmt.Println("жарко")
	}

}
