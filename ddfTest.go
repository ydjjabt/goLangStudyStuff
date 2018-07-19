package main

import "fmt"
import "strings"
import "time"

func main() {
    x := "sna@ke@gijoe"

    i := strings.Index(x, "@g")
    fmt.Println("Index: ", i)
    if i > -1 {
        chars := x[:i]
        arefun := x[i+1:]
        fmt.Println(chars)
        fmt.Println(arefun)
    } else {
        fmt.Println("Index not found")
        fmt.Println(x)
    }

    yuop := strings.Split(x,"@")
    fmt.Println("dsafds", yuop[0])
    fmt.Println( yuop[1])
    fmt.Println( yuop[2])

    sA := "hi"
    sB := "bob"

    xs:= sA + " " + sB
    fmt.Println( xs)


    arr := []int{8,19,4,2,3,7}

    first := arr[2:4]
    second := arr[3:5]
    baf := append(first,second...)
    fmt.Println(first,second)
    fmt.Println(baf)

    modMap:=map[string]interface{}{"sender":"john", "data":989}
    fmt.Println(modMap, modMap["data"])

}