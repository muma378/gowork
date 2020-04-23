package main
import (
	"fmt"
	"errors"
)

type Animal interface{
	Eat()
	Move()
	Speak()
}

type AnimalType struct{
	name string
}

type Cow AnimalType
type Bird AnimalType
type Snake AnimalType

func (_ Cow) Eat() {
	fmt.Println("grass")
}

func (_ Cow) Move() {
	fmt.Println("walk")
}

func (_ Cow) Speak() {
	fmt.Println("moo")
}

func (_ Bird) Eat() {
	fmt.Println("worms")
}

func (_ Bird) Move() {
	fmt.Println("fly")
}

func (_ Bird) Speak() {
	fmt.Println("peep")
}

func (_ Snake) Eat() {
	fmt.Println("mice")
}

func (_ Snake) Move() {
	fmt.Println("slither")
}

func (_ Snake) Speak() {
	fmt.Println("hsss")
}

func createNewAnimal(arg string) (Animal, error) {
	var animal Animal
	switch arg {
	case "cow":
		animal = new(Cow)
	case "snake":
		animal = new(Snake)
	case "bird":
		animal = new(Bird)
	default:
		return nil, errors.New(fmt.Sprintf("No specified animal type: `%s` found", arg))
	}
	return animal, nil
}


func queryAnimalAction(arg string, animal Animal) (bool, error) {
	switch arg {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	default:
		return false, errors.New(fmt.Sprintf("No specified action: `%s` found", arg))
	}
	return true, nil
}

func main() {
	var cmd, name, arg string
	animalSet := make(map[string] Animal)
	for {
		fmt.Printf("> ")
		_, e := fmt.Scan(&cmd, &name, &arg)
		if e != nil {
			fmt.Println("Invalid input, re-enter please.")
			continue
		}
		switch cmd {
		case "newanimal":
			a, err := createNewAnimal(arg)
			if err != nil {
				fmt.Println(err)
				continue
			}
			animalSet[name] = a
		case "query":
			_, err := queryAnimalAction(arg, animalSet[name])
			if err != nil {
				fmt.Println(err)
				continue
			}
		default:
			fmt.Printf("No specified command `%s` found\n", cmd)
		}
	}

}
