package main

import (
    "fmt"
    "sync"
)

var Wait sync.WaitGroup
var Counter int = 0

// var lock sync.Mutex = sync.Mutex{}   .... bottom is same as this line
var lock sync.Mutex //lock is type sync.Mutex with default initialization since it didnt get constructed with specific properties
func main() {

    for routine := 1; routine <= 10; routine++ {

        Wait.Add(1)
        go Routine(routine)
    }

    Wait.Wait()
    fmt.Printf("Final Counter: %d\n", Counter)
}

func Routine(id int) {
    defer Wait.Done()
    // Wait.Done()
    for count := 0; count < 2; count++ {
        fmt.Println("routine id:", id, "before <<lock. lock is ", lock)
                lock.Lock()
        fmt.Println("routine id:", id, "after<< lock. lock is ", lock)

        value := Counter
        fmt.Println("routine id:", id, "before incrementing value is ", value, ". counter is ", Counter, ".index is ", count)
        value++
                fmt.Println("routine id:", id," after incrementing value is ", value, ". counter is ", Counter, ".index is ", count)

        Counter = value
                	lock.Unlock()
        fmt.Println("routine id:", id, "after >> unlock. lock is ", lock)


    }
    var x int
    fmt.Println("uninitialized x is ", x)
}