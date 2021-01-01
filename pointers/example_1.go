package main

import (
	"fmt"
	"strings"
)

type Person struct {
	Name string
	Age int32 
}

func (person *Person) SanitizeName()  {
	person.Name = strings.TrimSpace(strings.ToUpper(person.Name))
}

func main()  {
	p := Person{
		Name: "random dude",
		Age: 19,
	}
	fmt.Println(p)
	p.SanitizeName()
	fmt.Println(p)

	pList := []*Person {
		{
			Name: "random dude",
			Age: 32,
		},
		{
			Name: "another random dude",
			Age: 20,
		},
	}

	fmt.Println(pList[0])
	for _, person := range pList {
		person.SanitizeName()
	}
	fmt.Println(pList[0])
}
