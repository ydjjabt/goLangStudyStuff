package main

import(
		"fmt"
		"reflect"
)
type employee struct{
		name string
		age int
	}
	var	ed = &employee{}
	var edTwin = *ed

func main(){
	

	 tim  := employee{}

	 timTwin := &tim

	 joeRat := tim
	fmt.Println("dhasjkd",tim, timTwin, joeRat,)
	timTwin.name = "snake"
	joeRat.name = "geronimo"
		fmt.Println("yigijkl",tim, timTwin, joeRat)

	fmt.Println("dsjfjasdsdfa",reflect.TypeOf(add(5	)).Kind() )
		citiexs := []string{}
	fmt.Println("kdfasdfase", reflect.TypeOf(citiexs).Elem())
	fmt.Println("hehioherhe",reflect.TypeOf(tim).Name(),"",reflect.TypeOf(tim).Kind() )
	fmt.Println( add(5)(7))


	
}

// this is function closure
func add(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}