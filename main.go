package main

import (
	"fmt"
	"time"

	print "github.com/untillpro/directcd-test-print"
)

func main() {

	fmt.Println("directcd_test", 1)

	cnt := 0
	for {
		cnt++
		print.Print(cnt)
		select {
		case <-time.After(time.Second * 10):
		}
	}
}
