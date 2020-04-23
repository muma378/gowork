package main
import (
    "fmt"
)


func Swap(numbers []int, index int) {
    temp := numbers[index]
    numbers[index] = numbers[index+1]
    numbers[index+1] = temp
}

func BubbleSort(numbers []int){
    for i := 0; i < len(numbers); i++ {
        // The rightmost i numbers were sorted
        for j := 0; j < len(numbers)-i-1; j++ {
            if numbers[j] > numbers[j+1] {
                // Moves the bigger to the right slot
                Swap(numbers, j)
            }
        }
    }
}

func main(){
    // numbers := []int{11, 23, 56, 4, 12, 7, 83, 32, 1, 11}
    numbers := make([]int, 0, 10)
    var enteredNum int
    // Promotes to input up to 10 integers
    fmt.Println("Enters integers(up to 10):")
    for i := 0; i < cap(numbers); i++ {
        fmt.Scan(&enteredNum)
        numbers = append(numbers, enteredNum)
    }
    BubbleSort(numbers)
    fmt.Print(numbers)
}
