package main
import(
"sync"
"time"
"math/rand"
"fmt"
)
///////////////////////
/* randomInt and randomString are two function 
	to generate random string
	use randomString(25) ... 25 is the number of string character to produce
*/
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
type value struct{
	lock sync.Mutex
	aChan chan int
	bChan chan int

}
func printSum( a *value, b *value){
	defer wg.Done()
	aTemp:=<-a.aChan
	aTemp+=1
	a.aChan<-aTemp
	time.Sleep( time.Second )
	bTemp:=<-b.bChan
	bTemp+=1
	b.bChan<-bTemp

	aI:=<-a.aChan
	bI:=<-b.bChan
	fmt.Println( "sum is ", aI + bI)
}

var wg sync.WaitGroup
func main(){
	var a, c value
	wg.Add(2)
	go printSum(&a,&c)
	go printSum(&c,&a)
	fmt.Println("from main routine a lock is ",a.lock,". c lock is ",c.lock )
	wg.Wait()
	
}