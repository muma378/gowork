package main

import (
    "fmt"
    "sort"
    "strconv"
)

func main()  {
    var enterChar string
    var index int
    digits := make([]int, 3, 10)
    for {
        fmt.Println("Enter an integer:")
        fmt.Scan(&enterChar)
        if enterChar == "X" {
            fmt.Println("Exited.")
            return
        }
        if enterNum, err := strconv.Atoi(enterChar); err != nil {
            fmt.Printf("'%s' is not an interger, re-enter\n", enterChar)
            continue
        } else {
            if index < len(digits) {
                digits[index] = enterNum
            }else{
                digits = append(digits, enterNum)
                print(len(digits))
            }
            index += 1
        }
        copyDigits := make([]int, len(digits))
        copy(copyDigits, digits)
        sort.Slice(copyDigits, func (i, j int) bool {
            return copyDigits[i] < copyDigits[j]
        })
        fmt.Println(copyDigits)
    }
}
