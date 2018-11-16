package main

import( 
"fmt"
)

func main(){

	// var aParent organism = parent{name:"parent name", walk:"parent walk"}

var john organism = child{name:"child name", parent: parent{name:"cool aid", walk:"cat walk"} }


fmt.Println( john.getName() )
fmt.Println( john.getWalk() )



/*
for composition no polymorphism: cannot use setter method to change child  property or parent property. 
can use dot operator to change child property or parent property.
can  use getter method or dot operator to get child property or parent property.
child inherit parent property and functions.


for polymorphism composition, u inherit child property and function. cannot change parent or child property. cann use dot operator to get property.

u cannot embed polymorphism.

when polymorphism reassignment, you can use only the inherited property and function. cannot use override function

when polymorphis, u cannot change property.
*/


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


/////////monkey struct
func(p parent) getName()string{
	return "inherited from parent " + p.name
}

func(p parent) getWalk()string{
	return "inherited from parent " + p.walk
}



//////////////human struct



func( c child) getName()string{
	return "override over parent "  + c.name
}

func(c child) getWalk()string{
	return "override over parent " + c.parent.walk
}


