package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

//O(n^2)
func main() {
	people := []Person{
		{"Bob", 32},
		{"John", 25},
		{"Michael", 54},
		{"Jenny", 43},
	}
	var sorted []Person

	for range people {
		greater := searchGreater(people)
		sorted = append(sorted, people[greater])
		people = append(people[:greater], people[greater+1:]...) //A elipse (...) é usada para desempacotar os valores de um slice, já que a função "append" espera argumentos e não um slice
	}

	fmt.Println(sorted)
}

func searchGreater(people []Person) int {
	greater := 0
	for k, v := range people {
		if v.Age > people[greater].Age {
			greater = k
		}
	}
	return greater
}
