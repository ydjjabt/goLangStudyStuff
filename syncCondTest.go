package main



import(
	"fmt"
	"sync"
	// "time"
)	



func main(){
	
	 c := sync.NewCond( &sync.Mutex{})

	 var x = 9

	 go func(){
	 	fmt.Println("in the new goroutine... x is ", x)
	 	c.L.Lock()
	 		x = 100
	 	fmt.Println(" in the new goroutine......Signal 1 ")
	 	c.Signal()
	 	c.L.Unlock()

	 	// c.L.Lock()
	 	// 	x = 555
	 	// c.Signal()
	 	// c.L.Unlock()

	 }()
	 	fmt.Println("1 in the main goroutine ... x is",x)

// c.L.Lock() before c.Wait() follow c.L.Unlock() to make the goroutine wait for the Signal() from the other goroutine that this goroutine can continue running 	
	c.L.Lock()
	// cannot access read while it is locked
		// fmt.Println("2 in the main goroutine ... x is",x) 
	c.Wait()
	//cannot access read while it is locked
		// fmt.Println("3 in the main goroutine ... x is",x)
	c.L.Unlock()
	fmt.Println("wait 1 in the main goroutine ... x is",x)

	// c.L.Lock()
	// c.Wait()
	// c.L.Unlock()
	// 	fmt.Println("wait 1 in the main goroutine ... x is",x)


}