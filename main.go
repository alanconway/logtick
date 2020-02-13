package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	for {
		fmt.Fprintf(os.Stderr, "log tick at %v\n", time.Now())
		time.Sleep(3 * time.Second)
	}
}
