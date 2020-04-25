package main

import (
    "os"
    "fmt"
    "sort"
    "math"
    "bufio"
    "strconv"
    "strings"
    )

var LOOP int = 4

func sliceSort(inCh chan []int, outCh chan []int)  {
    nums := <- inCh
    fmt.Println(nums)
    sort.Ints(nums)
    outCh <- nums
}

func merge(subarrays [4][]int) []int  {
    var heads [4]int    // pointer to subarrays' elements
    result := make([]int, 0)
    var iMinElem, minElem int
    var notInitFlag bool
    for {
        notInitFlag = true
        for i, j := range heads {   // find the smallest in all arrays' heads
            if j >= len(subarrays[i]) {
                continue
            }
            pointerTo := subarrays[i][j]
            if notInitFlag || minElem  > pointerTo{
                notInitFlag = false
                iMinElem = i
                minElem = pointerTo
            }
        }
        if notInitFlag {
            // iteratives all elements in head, but no one was available
            break
        }else{
            heads[iMinElem] ++ // moves forward in the ith array
            result = append(result, minElem)
        }
    }
    return result
}

func main()  {
    inCh := make(chan []int, 4)
    outCh := make(chan []int, 4)

    fmt.Println("Input a series of integers (space separated): ")
    reader := bufio.NewReader(os.Stdin)
    text, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println(err)
        return
    }
    scanner := bufio.NewScanner(strings.NewReader(text))
    scanner.Split(bufio.ScanWords)
    nums := make([]int, 0)
    for scanner.Scan() {
        val := scanner.Text()
        num, err := strconv.Atoi(val)
        if err != nil {
            fmt.Printf("Unable to parse the input value: %s", val)
            return
        }
        nums = append(nums, num)
    }
    nNums := len(nums)
    sliceSize := int(math.Ceil(float64(nNums) / float64(LOOP)))
    for i := 0;i < LOOP;i++ {
        go sliceSort(inCh, outCh)
        if (i+1)*sliceSize > nNums {
            inCh <- nums[i*sliceSize:nNums]
        }else{
            inCh <- nums[i*sliceSize:(i+1)*sliceSize]
        }
    }

    var subarrays [4][]int
    for i := 0; i < LOOP; i++ {
        subarrays[i] = <- outCh
        // fmt.Println(subarrays[i])
    }

    result := merge(subarrays)
    fmt.Println(result)
}
