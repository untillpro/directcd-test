package main

import (
	"fmt"
	"time"
)

func main() {
	cnt := 0
	for {
		cnt++
		fmt.Println(cnt, time.Now())
		select {
		case <-time.After(time.Second * 10):
		}
	}
}
