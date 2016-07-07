package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

const (
	amount = 10000000
)

type todo struct {
	msg string
}

func main() {
	var info []*todo
	var mutex sync.Mutex
	var wg sync.WaitGroup

	start := time.Now()

	for i := 0; i < amount; i++ {
		wg.Add(1)

		go func(j int) {
			mutex.Lock()
			info = append(info, &todo{msg: "abc" + strconv.Itoa(j)})
			mutex.Unlock()

			wg.Done()
		}(i)
	}

	wg.Wait()

	fmt.Println(len(info))
	fmt.Println(time.Since(start))
}
