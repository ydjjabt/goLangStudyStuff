package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan interface{})
	d := make(chan interface{})
	go func() {
		// defer close(d)
		// 	defer close(c)
		for i:=1;i<=2;i++{
			
			c<-i
			d<-i*10
		}
		c<-2323
	}()
/*
when channel is not close, default case in select block will execute. 
when channel is close, close channel case in select block will execute.
if one channel is open and one is close. close channel case  will execute
and default case will not execute
*/
		/*
		defer close(d)
		c<-666
		d<-11
		========
		defer close(d)
		defer close(c)
		c<-666
		d<-11
		=======
			defer close(d)
			c<-666
			d<-11
			when select only(no for encapsulate select)for all cases above only 666 send to channel c. 
			11 doesnt get send to channel d

		*/

			/*
		defer close(d)
		defer close(c)
		c<-666
		d<-11
		Blocking on read...
from channel C 666 true
from channel D 11 true
from channel C <nil> false
from channel D <nil> false
from channel D <nil> false
from channel C <nil> false




			*/

	
	runAgain:=true
	for runAgain{
		fmt.Println("still running in while loop")
		select {
			case cV,cS:=<-c: // <2>
			fmt.Println("c channel BLOCK")
				time.Sleep(1 * time.Second)
				if cS{
				 fmt.Println("channel C",cV,cS)
				 runAgain = true
				}else{ 
					fmt.Println("channel C",cV,cS)
					runAgain = false}
				// fmt.Println("channel C",cV,cS)
			case dV,dS:=<-d:
				fmt.Println("d channel BLOCK")
				time.Sleep( time.Second)
				if dS{ 
					fmt.Println("channel D",dV,dS)
					runAgain = true
				}else{ 
					fmt.Println("channel D",dV,dS)
					runAgain = false}
				// fmt.Println("channel D",dV,dS)

			default:
				time.Sleep( time.Second)
				fmt.Println("default case happening")
				runAgain = false
		}
	}

	
	
}
