package main


import(
	"fmt"
)

type human struct{
	name string
	age int
}

type employee struct{
	employeeID int
	human  
}

func (h *human) talk(){
	fmt.Println("human can talk")
}

func (e *employee) work(){
	fmt.Println("employee work to make money")
}
func main(){

	// asian := human{}
	// asian.talk()

	jack := &employee{  human : human{name:"jack",age :99}, employeeID: 34756}
	jack.talk()
	jack.work()
	
}