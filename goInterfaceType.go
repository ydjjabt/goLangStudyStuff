package main

import (
	"fmt"
)

type Animal interface {
	Speak() string
}

type Dog struct {
}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {
	name string
}

func (c *Cat) Speak() string {
	return "Meow!"
}

type Llama struct {
}

func (l Llama) Speak() string {
	return "?????"
}

type JavaProgrammer struct {
}

func (j JavaProgrammer) Speak() string {
	return "Design patterns!"
}
func main() {
	animals := []Animal{Dog{}, &Cat{name:"joe"}, Llama{}, JavaProgrammer{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}


	arr := []int{}
	fmt.Println("empty array->", arr)
	arr = append(arr,77)
	arr = append(arr,22)
	arr = append(arr,123)
	fmt.Println(arr)

	for k,v := range arr{
		fmt.Println(k,v)
	}

	m := map[string]int{}
	fmt.Println("empty map->", m)
	m["hi"] = 998
	m["go"] = 111
	m["where"] = 779

	fmt.Println(m)
	delete(m, "go")
		fmt.Println(m)

	for k,v:=range m{
		fmt.Println(k,v)
	}

	    a := []string{"Hello1", "Hello2", "Hello3", "snake", "geronimo"}
	    fmt.Println(a)
	    a = append(a[:2], a[3:]...)
	    fmt.Println(a)



}
