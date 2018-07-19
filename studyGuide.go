package main

import (
	"fmt"
	"strconv"
	"reflect"
	"strings"
	"math/rand"
)
///////////////////////
/* randomInt and randomString are two function to generate random string
	use randomString(25) ... 25 is the number of string character to produce
	must import "math/rand"
*/
func randomInt(min, max int) int {
    return min + rand.Intn(max-min)
}

// Generate a random string of A-Z chars with len = l
func randomString(len int) string {
    bytes := make([]byte, len)
    for i := 0; i < len; i++ {
        bytes[i] = byte(randomInt(65, 90))
    }
    return string(bytes)
}
///////////////////////

// this is function closure
func add(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}
/////////////////////////
func main() {
tyu:="giJoe"
fmt.Println( "dhfsd",tyu[0:2])

for i,v := range tyu{
	fmt.Println(i,string(v))
}
	// hashmap
	city := map[string]int{
		"New York":    8336697,
		"Los Angeles": 3857799,
		"Chicago":     2714856,
	}
	for key, value := range city {
		fmt.Println( key, value)
	}
	// put new key and value into map 
	city["Tokyo"]=6969 
	fmt.Println("$$$$$$$$$$$$$$$$$$", city)
	delete( city, "Chicago") // delete key and value from map
		fmt.Println("$$$$$$$$$$$$$$$$$$", city)

	// dynamic array
	cities := []string{}
	cities = append(cities, "San Diego")
	cities = append( cities, "NYC")
	cities = append(cities, "SF")
	cities = append(cities, "33")
		cities = append(cities, "iowa")

	for i:=0; i< len(cities); i++{
		fmt.Println( i,cities[i])
	}
	for i,v:= range cities{
		fmt.Println("range style", i, string(v))
	}
	fmt.Println( "popdfd",cities[0:2])
	fmt.Println ( "foo bar",cities )
	// convert integer to string
	i := 123
	t := strconv.Itoa(i)
	fmt.Println(reflect.TypeOf(t).Name() =="string", " String type????")

	
	// get type of a variable
	var name = "papa john"
	var red = "977";
	fmt.Println( reflect.TypeOf(name).Name() == "string", " name type reflect is ")
	println( reflect.TypeOf(name).Name() )
	fmt.Println( reflect.TypeOf(red).Name() == "int", " red type reflect is ")
	
	// conditiions
	if 5 == 5{
		fmt.Println("5 equal 5")
	}
	if 5 <= 7{
		fmt.Println(" 5 less than equal to 7")
	}
	if 5>=2{
		fmt.Println("5 greatr than or equal to 2")
	}
	


	// iterate over a string
	g := "hello"
	// var result string;
	result:="";
	for i:=0; i <len(g);i++ {
		fmt.Println("index is ", i, " .... character is ", string(g[i]))
			result += string(g[i]) + "$"
        }

fmt.Println("result is ", result)

// string to lowercase
 fmt.Println("GopHer to lowercase is ", strings.ToLower("GopHer"))


 ghet := "d"
 ghet += "8" //string cannot add number. only add string
 fmt.Println( ghet)


 i, e := strconv.Atoi("69")

if( e != nil){
	fmt.Println("there an error", reflect.TypeOf(e), e == nil)
} else{
	fmt.Println("number is ", i, reflect.TypeOf(i).Name() , e, e == nil)
}


	// assign function to a variable and pass function as argument
f := func(a, b int) int {
		return a + b
	}
	simple(f)

	

	 var yu I = T{name:"tim",age:99}
	 fmt.Println( yu.mm())
fmt.Println( reflect.TypeOf(yu).Name() )
	 
	 su := S{name:"jane", age: 11}
	 su.hobby = "enjoy life"
	fmt.Println( "hobbyyy ",su.showHobby()) 
	// su:=I{name:"jane", age: 11} cannot instantiate an Interface
	 fmt.Println( su.mm() )
	 fmt.Println (reflect.TypeOf(su).Name())
} // /////////////////////////////////////////////////////////////////////end of main program


func simple(a func(a, b int) int) {
	fmt.Println(a(60, 7))
}
////////////////////////////////////////////////// OOP programming
type I interface {
    mm() (string,int)
}
type T struct {
    name string;
    age int
}

type S struct{
	name string;
	age int;
	hobby string
}
func (x *S) showHobby() string{
	return x.hobby
}
func (t T) mm() (string,int) {
    return t.name,t.age
}

func (s *S) mm() (string,int){
	return s.name,s.age
}
/////////////////////////////////