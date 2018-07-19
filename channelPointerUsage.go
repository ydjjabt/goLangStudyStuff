package main

import(
		"sync"
		"fmt"
)



var wg sync.WaitGroup

func main(){

var ch = make( chan int,3)
fmt.Println("main goroutine waitGroup before add(1) is ", wg)
wg.Add(1)
fmt.Println("main goroutine waitGroup after add(1) is ", wg)

	go func(ch *chan int){
		defer wg.Done()
		defer close(*ch)
		fmt.Println("new goroutine", " channel length is ", len(*ch),"waitGroup is ", wg )

		*ch<-76
		*ch<-123
		*ch<-874

				fmt.Println("new goroutine", " channel length after inputting into channel is ", len(*ch), "waitGroup is ", wg )
			
		}(&ch)
fmt.Println("main goroutine <<beforewaitGroup Wait>> is call channel length is ", len(ch))
wg.Wait()
fmt.Println("main goroutine **after waitGroup Wait** is call channel length is ", len(ch))

	v,e:=<-ch

	fmt.Println("v is ", v, ". e is ", e)
}