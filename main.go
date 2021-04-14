package main

import (
	"fmt"
	exclude "github.com/cncodetube/course-go-pkg-module-exclude"
	_ "golang.org/x/text"
)

func main() {
	fmt.Println(exclude.GetUUID())
}
