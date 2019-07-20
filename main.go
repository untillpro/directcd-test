package main

import (
	"fmt"
	"os"
	"time"

	print "github.com/untillpro/directcd-test-print"
)

func main() {

	fmt.Println("directcd_test", 4, "args=", os.Args)

	cnt := 0
	for {
		cnt++
		print.Print(cnt)
		select {
		case <-time.After(time.Second * 10):
		}
	}
}
