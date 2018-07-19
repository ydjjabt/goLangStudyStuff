package main

import (
"fmt"
"time"
)
func main(){
    chA := make( chan int)
  

    go func(){
        for{
            select{
                 case chA<-9:
                 }
        }

        
    }()

for v := range chA{
    fmt.Println(v)
    time.Sleep( time.Second)
}



}