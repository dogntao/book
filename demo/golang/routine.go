package golang

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func solutine(j int, ch chan int) {
	// time.Sleep(time.Second * 1)
	defer func() {
		fmt.Println(<-ch)
		wg.Done()
	}()
	// fmt.Println(j)
}

func Routine() {
	i := 1680000
	j := 1
	ch := make(chan int, 20)
	for i >= 0 {
		wg.Add(1)
		ch <- i
		go solutine(j, ch)
		j++
		i -= 1000
	}
	wg.Wait()
}
