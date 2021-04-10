package main

import (
	"fmt"
	"github.com/cncodetube/faker"
	"github.com/google/uuid"
	_ "golang.org/x/text"
)

func main() {
	id := uuid.New().String()
	fmt.Println(id)

	faker := faker.New()
	fmt.Println(faker.Name())
}
