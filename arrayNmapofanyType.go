package main

import("fmt"

)

func main(){

	arr := []interface{}{}


	arr = append(arr, 8, 765)

	fmt.Println(arr, len(arr) )


	mMap := map[interface{}]interface{}{}

	mMap[1] = "hello"
	mMap["a"] = 2.55

	fmt.Println(mMap, len(mMap))


	

}