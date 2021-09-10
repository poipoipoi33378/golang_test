package main

import (
	"fmt"
	"zoo/animals"
)

const Name = "Zoo Application"

func main() {
	fmt.Println(AppName())

	fmt.Println(animals.ElephantFeed())
	fmt.Println(animals.Elepahant)
	fmt.Println(animals.RabbitFeed())
	fmt.Println(animals.MonkeyFeed())
}
