package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 50; i++ {
		x := func(i int) {
			fmt.Println(i)
		}
		go x(i)
	}

	time.Sleep(4)
}
