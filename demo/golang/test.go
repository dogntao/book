package golang

import (
	"fmt"
	"time"
)

var runStat bool

func Test() {
	if runStat {
		return
	} else {
		fmt.Println("starting")
	}
	runStat = true
	defer func() {
		runStat = false
	}()
	time.Sleep(time.Second * 100)
}
