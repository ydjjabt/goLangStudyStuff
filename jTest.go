package main

import( 
"fmt"
)

func main(){


var john organism = child{name:"Child name", parent: parent{name:"cool aid", walk:"cat walk"} }

var tom organism = parent{name:"PParent", walk:" The Run"} 
fmt.Println( john.getName() )

fmt.Println( john.getWalk() )



john = tom
fmt.Println( john.getName() )

fmt.Println( john.getWalk() )

}

type alien struct{
	name string
}

type parent struct{
	name string
	walk string
}


type child struct{
	name string
	parent 

}

type organism interface{
	getName()string
	getWalk()string
}


/////////parent struct
func(p parent) getName()string{
	return "inherited from parent " + p.name
}

func(p parent) getWalk()string{
	return "inherited from parent " + p.walk
}



//////////////child struct

func(c child) getName()string{
	return "override parent " + c.name
}

func(c child) getWalk()string{
	return "override parent " + c.parent.walk
}

/*
inherited class should always override parent function. polymorphism or not
if type is interface, u cannot change property. can change when it not interface using dot operator

*/

