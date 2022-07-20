package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/lawjohnny/stupid-dumb-cal/simple"
)

func main() {
	uuid := uuid.New()
	fmt.Println(uuid)
	fmt.Printf("The sum is : %d\n", simple.Plus(1, 2))
}
