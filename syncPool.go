package main


import (
	"fmt"
	"sync"
	// "time"
)

// Pool for our struct A
var pool *sync.Pool

// A dummy struct with a member 
type A struct {
	Name string
}

// Func to init pool
func initPool() {
	pool = &sync.Pool {
		New: func()interface{} {
			fmt.Println("Returning new A")
			// return new(A)
			return &A{}
		},
	}
}

// Main func
func main() {
	// Initializing pool
	initPool()
	// Get hold of instance one


	fmt.Println( "dasdfasdfa",pool.Get())
	one := pool.Get().(*A)
	one.Name = "first"
	fmt.Printf("one.Name = %s\n", one.Name)
	// Submit back the instance after using
	pool.Put(one)
	// Now the same instance becomes usable by another routine without allocating it again
}