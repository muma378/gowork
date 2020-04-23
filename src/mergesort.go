package main

import (
    "fmt"
    "sort"
    "math"
    )

var LOOP int = 4

func mergeSort(inCh chan []int, outCh chan []int)  {
    nums := <- inCh
    fmt.Println(nums)
    sort.Ints(nums)
    outCh <- nums
}

func main()  {
    inCh := make(chan []int, 4)
    outCh := make(chan []int, 4)

    nums := []int{8, 7, 3, 5, 10, 22, 1, 3, 6, -1}
    nNums := len(nums)
    sliceSize := int(math.Ceil(float64(nNums) / float64(LOOP)))
    for i := 0;i < LOOP;i++ {
        go mergeSort(inCh, outCh)
        if (i+1)*sliceSize > nNums {
            inCh <- nums[i*sliceSize:nNums]
        }else{
            inCh <- nums[i*sliceSize:(i+1)*sliceSize]
        }
    }
    for i := 0; i < LOOP; i++ {
        sortedNums := <- outCh
        fmt.Println(sortedNums)
    }
}
