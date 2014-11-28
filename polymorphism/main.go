package main

import "fmt"

type intInterface struct {
}

type stringInterface struct {
}

func (self intInterface) Add(a, b int) int {
	return a + b
}

func (self stringInterface) Add(a, b string) string {
	return a + b
}

func main() {
	number := new(intInterface)
	fmt.Println(number.Add(1, 2))
	text := new(stringInterface)
	fmt.Println(text.Add("1", "2"))
}
