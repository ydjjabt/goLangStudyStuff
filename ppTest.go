package main

import( "fmt"
	"reflect"
)
type Animal interface {
    Speak() string
}
// type Dog struct {
//     Age interface{}
// }
type Dog struct {
}

func (d Dog) Speak() string {
    return "Woof!"
}

type Cat struct {
}

func (c Cat) Speak() string {
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

	city := map[string]interface{}{
		"New York":    8336697,
		"Los Angeles": "hi",
		"Chicago":     2714.856,
		"snake": true,
	}

fmt.Println(city)
//     dog := Dog{}
//     dog.Age = "3"
//     fmt.Printf("%#v %T\n", dog.Age, dog.Age)

//     dog.Age = 3
//     fmt.Printf("%#v %T\n", dog.Age, dog.Age)

//     dog.Age = "not really an age"
//     fmt.Printf("%#v %T", dog.Age, dog.Age)

// city := map[string]int{
// 		"New York":    8336697,
// 		"Los Angeles": 3857799,
// 		"Chicago":     2714856,
// 	}
//     dog.Age = city
//     fmt.Printf("%#v %T", dog.Age, dog.Age)

var tred interface{}

tred = 9
fmt.Println( "dsayuafa", reflect.TypeOf(tred).Name())
tred = "hi dasfasd"
fmt.Println( "sre2e", reflect.TypeOf(tred).Name())
tred = 53.2
fmt.Println( "doubeww", reflect.TypeOf(tred).Name())



animals := []Animal{Dog{}, Cat{}, Llama{}, JavaProgrammer{}}
    for whatThis, animal := range animals {
        fmt.Println(animal.Speak())
        fmt.Println("uiuiwwer", reflect.TypeOf(animal).Name())
                fmt.Println("qwe", reflect.TypeOf(animal).Kind())

        fmt.Println("dafsdasd", whatThis)
          fmt.Println("werw", reflect.TypeOf(whatThis).Name())
                fmt.Println("34223", reflect.TypeOf(whatThis).Kind())
                fmt.Println()

    }
}