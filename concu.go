package main

import (
	"fmt"
	"time"
	"sync"
)

var wg sync.WaitGroup

func say(s string) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Microsecond * 100)
	}

}

func main() {
	wg.Add(1)
	go say("Hey")
	wg.Add(1)
	go say("There")
	//wg.Wait()

	wg.Add(1)
	go say("dude")
	wg.Wait()
	//time.Sleep(time.Second)
}
