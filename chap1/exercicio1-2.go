package main

import (
	"fmt"
	"os"
    "strconv"
    "time"
)

func main() {
    start := time.Now() 
	s, sep := "", ""

	for idx, arg := range os.Args[1:] {
        s += sep + strconv.Itoa(idx) + "-" + arg
		sep = " "
	}
	fmt.Println(s)
    fmt.Printf("%.4f elapsed\n", time.Since(start).Seconds())
}
