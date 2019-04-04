package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("directcd_test", 1)

	cnt := 0
	for {
		cnt++
		fmt.Println(cnt, time.Now())
		select {
		case <-time.After(time.Second * 10):
		}
	}
}
