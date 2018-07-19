package main

import("fmt"
"sync")
var wg sync.WaitGroup

func partition(arr []int,start int, end int) int{
	partitionIndx := start - 1
	ch :=make(chan int, end - start + 2)
	for i:=start;i<end;i++{
		ch <-i
		wg.Add(1)
		go func(ch chan int){
			defer wg.Done()
			i:= <-ch
			if arr[end] > arr[i]{
				partitionIndx+=1
				swap := arr[partitionIndx]
				arr[partitionIndx] = arr[i]
				arr[i] = swap
			}
		}(ch)
	}
	wg.Wait()

	partitionIndx+=1
	swap := arr[partitionIndx]
	arr[partitionIndx] = arr[end]
	arr[end] = swap
	return partitionIndx
}

func quickSort(arr []int, start int, end int){
	if( start < end){
		partitionIndx := partition(arr,start,end)
		 quickSort(arr,start, partitionIndx - 1)
		 quickSort(arr, partitionIndx+1, end)
	}

}
func main(){
arr := []int{8,78,212,19,177,-6,0,123,1,3,234,23,13,3456,3239,332,342,5234,5453,
	163,45,77,83,456,-7,39,53,45,12,33,324,-7,85,2,898,7363,882,9923,736,981,342,234,211,342,534,765,
	2,4,16,5,65,32,2,898,22,4,56,5,65,32,4,5,1,414,322,141,14,384,273,9948,4747,9381,872,124,352,
	23423,145,45423,6,3,254,10,15,11,20,72,9,7,8,9,23,71,2,35,42,489,52,378,44,5,432,3,4,5,366,432,645,83,938,
	234,542,23,45,29,43,4,-7,89,2,898,12,4,6,5,64,32,3423,22312,243,532,5235,63,-9}

	quickSort(arr,0,len(arr) - 1)

	fmt.Println(arr)

}