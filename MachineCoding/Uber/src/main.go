package main

import (
	"fmt"
	"time"
)

func main2(i int, a chan int, done chan bool) {
	fmt.Println("Received i=", i)
	if i%7 == 0 {
		fmt.Println("Populating channels for i=", i)
		a <- 7
		done <- true
		return
	}
	fmt.Println("Sleeping for i=", i)
	time.Sleep(7 * time.Second)
	fmt.Println("Done Sleeping for i=", i)
}
func main1(a chan int) {
	count := 0
	done := make(chan bool, 5)
	for i := 1; i < 15; i++ {
		go main2(i, a, done)
		count++
		fmt.Println(count)
		if count%5 == 0 {
		forloop1:
			for {
				select {
				case d := <-done:
					fmt.Printf("DONE RECEIVED with value %+v", d)
					if d {
						fmt.Println("DONE RECEIVED")
						return
					}
				case <-time.After(5 * time.Second):
					fmt.Println("Time for next batch")
					break forloop1
				}
			}
			fmt.Println("------------------------------------")
		}
	}

	fmt.Println("DONE Exceuting main1")
}

func main() {
	a := make(chan int, 5)
	go main1(a)
	// forloop:
	//for {
	select {
	case val := <-a:
		fmt.Println("RECEIVED ", val)
		break //forloop
	case <-time.After(1 * time.Minute):
		fmt.Println("main default!!")
		time.Sleep(1 * time.Second)
	}
	// }

	fmt.Println("HERE!!!!!")
	time.Sleep(20 * time.Second)
}
