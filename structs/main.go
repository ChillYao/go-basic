package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

func (p *Person) changeName(name string) {
	p.Name = name
}

func main() {
	myPerson := NewPerson("YAO", 25)
	myPerson.changeName("sb")
	// fmt.Printf("%+v", myPerson)
	// a := 9
	// b := &a
	// fmt.Println(a, b, *b, &a, &b)

	mySlice := []int{1, 2, 3}

	for index, _ := range mySlice {
		mySlice[index]++
	}

	fmt.Println(mySlice)

}
