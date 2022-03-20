package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Printf("Numero de CPUs al inicio: %v\n", runtime.NumCPU())
	fmt.Printf("Numero de Gorutinas al inicio: %v\n", runtime.NumGoroutine())

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("Hola desde la primera Gourutina.")
		wg.Done()
	}()

	go func() {
		fmt.Println("Hola desde la segunda Gourutina.")
		wg.Done()
	}()
	fmt.Printf("Numero de CPUs en el medio: %v\n", runtime.NumCPU())
	fmt.Printf("Numero de CPUs en el medio: %v\n", runtime.NumCPU())

	wg.Wait()

	fmt.Println("A punto de finalizar main.")
	fmt.Printf("Numero de CPUs al final: %v\n", runtime.NumCPU())
	fmt.Printf("Numero de CPUs al final: %v\n", runtime.NumCPU())

}
