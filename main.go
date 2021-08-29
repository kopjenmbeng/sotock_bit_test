package main

import (
	"runtime"
	// "fmt"
	// "time"
	"github.com/kopjenmbeng/sotock_bit_test/cmd"
)

func main() {
	// fmt.Println(time.Now().Unix())
	runtime.GOMAXPROCS(runtime.NumCPU())
	cmd.Run()
}
