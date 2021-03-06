package main

import (
	"fmt"
	"os"
    "strings"
    "time"
)

func main() {
    start := time.Now()
    fmt.Println(strings.Join(os.Args[1:], " "))
    fmt.Printf("%.4f elapsed", time.Since(start).Seconds())
}
