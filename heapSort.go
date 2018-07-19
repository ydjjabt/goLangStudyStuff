package main

import("fmt"
	"sync"
	"math/rand"
)
//////////////////
func randomInt(min, max int) int {
    return min + rand.Intn(max-min)
}

// Generate a random string of A-Z chars with len = l
func randomString(len int) string {
    bytes := make([]byte, len)
    for i := 0; i < len; i++ {
        bytes[i] = byte(randomInt(65, 90))
    }
    return string(bytes)
}
///////////////////////
 func swap(root int, size int, arr []int) {
	largestRoot := root
	leftChild := 2*root + 1
	rightChild := 2*root + 2
	if leftChild < size && arr[leftChild] > arr[largestRoot]{
		largestRoot = leftChild
	}
	if rightChild < size && arr[rightChild] > arr[largestRoot]{
		largestRoot = rightChild
	}
	if largestRoot != root{
		temp := arr[root]
		arr[root] = arr[largestRoot]
		arr[largestRoot] = temp
		swap(largestRoot,size,arr)
	}

}

var wg sync.WaitGroup
 func traverseHeap(arr []int) []int{
	size := len(arr)
	middleInd := size / 2 - 1
	ch := make(chan int, middleInd + 2)
	for i:= middleInd; i >= 0;i--{
		ch<-i
		wg.Add(1)
		go func(ch chan int){
			defer wg.Done()
		 	j:=<-ch
		 	swap(j,size,arr)
		 	return
		 }(ch)
	}
	wg.Wait()


	chLast := make( chan int, size + 2)
	for i:= size - 1; i>=1;i--{
		chLast<-i
		wg.Add(1)
		go func(chLast chan int){
			defer wg.Done()
			j:=<-chLast
			temp := arr[j]
			arr[j] = arr[0]
			arr[0] = temp
			swap(0,j,arr)
			return
		}(chLast)
		
	}
	wg.Wait()
	return arr
 }
func main(){

arr := []int{8,78,212,19,177,-6,0,123,1,3,234,23,13,3456,3239,332,342,5234,5453,
	163,45,77,83,456,-7,39,53,45,12,33,324,-7,85,2,898,7363,882,9923,736,981,342,234,211,342,534,765,
	2,4,16,5,65,32,2,898,22,4,56,5,65,32,4,5,1,414,322,141,14,384,273,9948,4747,9381,872,124,352,
	23423,145,45423,6,3,254,10,15,11,20,72,9,7,8,9,23,71,2,35,42,489,52,378,44,5,432,3,4,5,366,432,645,83,938,
	234,542,23,45,29,43,4,-7,89,2,898,12,4,6,5,64,32,3423,22312,243,532,5235,63,77}

fmt.Println( traverseHeap(arr))

}