package puppy

import (
	"fmt"

	"github.com/RiyaChaubey/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}

func version11() {
	fmt.Println("This is version v1.1.0")
}

func version12() {
	fmt.Println("This is version v1.2.0")
}

func version13() {
	fmt.Println("This is version v1.3.0")
}
