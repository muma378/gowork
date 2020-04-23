package main
import (
	"fmt"
)

type Animal struct{
	food string
	locomotion string
	noise string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {
	var name, req string
	farm := map[string] Animal {
		"cow": Animal{"grass", "walk", "moo"},
		"bird": Animal{"worms", "fly", "peep"},
		"snake": Animal{"mice", "slither", "hsss"},
	}
	for {
		fmt.Printf("> ")
		_, e := fmt.Scan(&name, &req)
		if e != nil {
			fmt.Println("Invalid input, re-enter please.")
			continue
		}

		switch req {
		case "eat":
			farm[name].Eat()
		case "move":
			farm[name].Move()
		case "speak":
			farm[name].Speak()
		default:
			fmt.Println("Invalid input, re-enter please.")
			continue
		}
	}

}
