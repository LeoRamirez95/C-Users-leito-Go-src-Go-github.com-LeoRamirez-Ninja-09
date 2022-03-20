package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Os\t", runtime.GOOS)
	fmt.Println("arch\t", runtime.GOARCH)

}
