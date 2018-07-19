package main
import(
	"sync"
	"fmt"
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
var wg sync.WaitGroup
func main(){
	ch := make(chan string,3)
	for _,salutation := range []string{"hello", "greetings", "good day"}{
	ch <- salutation
	// fmt.Println("salutation is ", salutation)
	wg.Add(1)
	go func(ch chan string ){
		id:= randomString(15)
		fmt.Println("goroutine id:",id, <-ch)
			defer wg.Done()
			return
	}(ch)
	}


	wg.Wait()
	

}