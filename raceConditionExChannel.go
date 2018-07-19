package main

import (
    "fmt"
    "sync"
)



var Wait sync.WaitGroup
var Counter int = 0

func main() {
    ch := make(chan int,10)
    ch<-Counter
    for routine := 1; routine <= 10; routine++ {

        Wait.Add(1)
        go Routine(routine, ch)
    }

    Wait.Wait()
    fmt.Printf("Final Counter: %d\n", <-ch)
}

func Routine(id int, ch chan int) {
    defer Wait.Done()
    for count := 0; count < 2; count++ {
      
        fmt.Println("goroutine id:", id, "lenght of channel:",len(ch))
        value := <- ch
        
        value++
              

        ch <- value


    }
}