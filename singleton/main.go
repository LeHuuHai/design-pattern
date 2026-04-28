package main

import (
	"design-pattern/singleton/counter"
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			counter := counter.GetInstance()
			for j := 0; j < 100; j++ {
				counter.Increase()
			}
		}()
	}
	wg.Wait()
	counter := counter.GetInstance()
	fmt.Println(counter.GetValue())
}
