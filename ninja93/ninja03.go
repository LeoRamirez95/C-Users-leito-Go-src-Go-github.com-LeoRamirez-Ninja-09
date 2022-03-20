package main

import (
	"fmt"
	"sync"
)

func main() {

	var mu sync.Mutex
	var wg sync.WaitGroup
	incremento := 0

	gs := 100

	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := incremento
			v++
			incremento = v
			fmt.Println(incremento)
			mu.Unlock()
			wg.Done()
		}()

	}
	wg.Wait()
	fmt.Println("cuenta final:", incremento)
}
