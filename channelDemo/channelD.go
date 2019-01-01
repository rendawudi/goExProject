package main

import (
	"fmt"
	"sync"
)

type work struct {
	in chan int
	doFunc func()
}

func doWork(w work) {
	for {
		fmt.Println(<-w.in)
		w.doFunc()
	}
}

func createWorks(i int, wg *sync.WaitGroup) work{
	w := work{
		in: make(chan int),
		doFunc: func() {
			wg.Done()
		},
	}
	go doWork(w)
	return w
}

func chanDemo() {
	var works [10]work
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		works[i] = createWorks(i, &wg)
	}

	for i := 0; i < 10; i++ {
			works[i].in <- 'a' + i
			wg.Add(1)
}
	wg.Wait()
}

func main() {
	chanDemo()
}

