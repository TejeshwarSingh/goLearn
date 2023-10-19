package main

import (
	"fmt"
	"sync"
)

// structs with tagging for json and yaml for serialization and deserialization.
type person struct {
	firstName string `json:"firstName" yaml:"firstName"`
	lastName  string `json:"lastName" yaml:"lastName"`
	age       int    `json:"age" yaml:"age"`
}

//func helloHandler(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprint(w, "Hello, Go!")
//}

func main() {
	//http.HandleFunc("/", helloHandler)
	fmt.Println("main start 2")
	//http.ListenAndServe(":8095", nil)
	var x string = "Hello, World"
	var y string = "Hello, World"
	var i int = 10

	fmt.Println(x == y)
	fmt.Println("i=", i)
	// define array

	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println("a=", a)
	// define slice
	b := []int{1, 2, 3, 4, 5}

	b = append(b, 6, 7)
	fmt.Println("b=", b)
	fmt.Println("2 to 5 slice=", b[2:5])

	for k, v := range b {
		fmt.Println("k=", k, "v=", v)
	}

	// define map
	c := map[string]int{"a": 1, "b": 2}
	fmt.Println("c=", c)

	// 2d map
	d := map[string]map[string]int{"a": {"a1": 1, "a2": 2}, "b": {"b1": 1, "b2": 2}}
	fmt.Println("d=", d)
	for k, v := range d {
		fmt.Println("k=", k, "v=", v)
		for k1, v1 := range v {
			fmt.Println("k1=", k1, "v1=", v1)
		}
	}

	// structs
	// initialize struct person
	p1 := person{firstName: "John", lastName: "Doe", age: 25}
	fmt.Println("p=", p1)

	type animal struct {
		name        string
		description []string
	}

	// initialize struct animal
	a1 := animal{name: "dog", description: []string{"bark", "run"}}

	fmt.Println("a1=", a1)

	for _, v := range a1.description {
		fmt.Println("v=", v)
	}

	//  animal struct as embedded struct promotion of fields
	type herbevor struct {
		animal
		food string
		sync.Mutex
	}

	h1 := herbevor{animal: animal{name: "cow", description: []string{"moo", "walk"}}, food: "grass"}
	h2 := herbevor{animal: animal{name: "goat", description: []string{"meh", "walk"}}, food: "grass"}

	fmt.Println("h1=", h1)
	fmt.Println("h2.name=", h2.name)

	h1.Lock()

	// anonymous struct with
	// No name for this struct
	// But we need to declare the type here
	bio := struct {
		firstName string
		friends   map[string]int
		favDrinks []string
	}{
		firstName: "Steven",
		friends: map[string]int{
			"Tim":   12345678,
			"Abdul": 34789657,
			"Kally": 28993332,
		},
		favDrinks: []string{
			"Pepsi",
			"Cocacola",
		},
	}
	fmt.Println(bio.firstName)

	for k, v := range bio.friends {
		fmt.Println(k, v)
	}

	for k, v := range bio.favDrinks {
		fmt.Println(k, v)
	}
	// defer function calls it is executed in a stack hence reverse order

	defer LastHi()
	defer func() {
		fmt.Println("Hi, defer")
	}()
	// functions
	// function can used as variable
	hv := Hello

	hello := hv()
	fmt.Println(hello)

	v1, v2 := ReturnTwoValues()
	fmt.Println("v1=", v1, "v2=", v2)

}

func Hello() string {
	return "Hello, World"
}
func ReturnTwoValues() (string, string) {
	return "val 1 ", "val 2"
}

func LastHi() {
	fmt.Println("Last Hi")
}
