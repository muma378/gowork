package main
import (
	"fmt"
	"time"
)

var deposit int = 100

func draw(amount int) {
	balance := deposit - amount
	fmt.Printf("Draws %d from deposit.\n", amount)
	deposit = balance
}

func save(amount int) {
	balance := deposit + amount
	fmt.Printf("Saves %d into deposit.\n", amount)
	deposit = balance
}

func main()  {
	fmt.Printf("Deposit: %d\n",  deposit)
	// Usually, the value of deposit should be update after doing save or draw
	go save(50)
	go draw(50)
 	// but, when the order of execution was interlaved
	// the value of deposit will be no more correct
	time.Sleep(time.Second * 1)
	fmt.Printf("Deposit: %d\n",  deposit)
}
