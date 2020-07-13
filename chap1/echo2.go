package main

import (
	"fmt"
	"os"
    "time"
)

func main() {
    start := time.Now()
	s, sep := "", ""

	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
    fmt.Printf("%.4f elapsed\n", time.Since(start).Seconds())
}
