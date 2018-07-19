package main

import(
"fmt"
)
func divide(arr []int) []int{

	// id:= randomString(15)
	if len(arr) > 1{
		lCh,rCh := keepDivding(arr)
		lout :=<-lCh
		rout :=<-rCh
		left := divide(lout)
		right := divide(rout)
		mergeNsortedArr := mergeSort(left,right)
		return mergeNsortedArr
	}else{
		return arr
	}
}

func keepDivding(arr []int)(chan []int, chan []int){

		left := func()chan []int{
			lCh := make(chan []int)
			go func(){
				lCh <- arr[0:len(arr)/2]
				return
			}()
			return lCh
		}()

		right := func()chan []int{
			rCh := make(chan []int)
			go func(){
				rCh <- arr[len(arr)/2:]
				return
			}()
			return rCh
		}()

		return left,right
			
}

func mergeSort(l []int, r []int)[]int{
	mergeNsortedArr := []int{}
	
	for; len(l) >= 1 || len(r) >= 1;{
		if( len(l) >= 1 && len(r) >= 1){
			notDone := true
			for;notDone;{
				if l[0] < r[0]{
					mergeNsortedArr = append(mergeNsortedArr, l[0])
					l = append(l[1:])
					if(len(l) >0){ notDone = true}else{notDone = false}
				}else{
					mergeNsortedArr = append(mergeNsortedArr, r[0])
					r = append(r[1:])
					if( len(r) > 0){notDone = true}else{notDone = false}
				}
			}
		}else if len(l) >= 1{			
				notDone := true
				for;notDone;{
					mergeNsortedArr = append(mergeNsortedArr, l[0])
					l = append(l[1:])
					if(len(l) >0){ notDone = true}else{notDone = false}
				}
		}else if len(r) >= 1{
				notDone := true
				for;notDone;{
					mergeNsortedArr = append(mergeNsortedArr, r[0])
					r = append(r[1:])
					if(len(r) > 0){ notDone = true}else{notDone = false}
				}
		}
	}

	return mergeNsortedArr
	
}

func main(){
	arr := []int{8,78,212,19,177,-6,0,123,1,3,234,23,13,3456,3239,332,342,5234,5453,
	163,45,77,83,456,-7,39,53,45,12,33,324,-7,85,2,898,7363,882,9923,736,981,342,234,211,342,534,765,
	2,4,16,5,65,32,2,898,22,4,56,5,65,32,4,5,1,414,322,141,14,384,273,9948,4747,9381,872,124,352,
	23423,145,45423,6,3,254,5,0,72,9,7,8,9,23,71,2,35,42,489,52,378,44,5,432,3,4,5,366,432,645,83,938,
	234,542,23,45,29,43,4,-7,89,2,898,12,4,6,5,64,32,3423,22312,243,532,5235,63,77}

	fmt.Println( divide(arr) )


}