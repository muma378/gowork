package main

import (
    "fmt"
    "strings"
    "strconv"
)

func main(){
    var x float64
    fmt.Printf("Please enter a floating point number:\n")
    fmt.Scan(&x)
    s64 := strconv.FormatFloat(x, 'f', -1, 64)
    // dotIndex := strings.Index(s64, '.')
    ss := strings.Split(s64, ".")
    fmt.Printf("Value after truncated: %s\n", ss[0])
}
