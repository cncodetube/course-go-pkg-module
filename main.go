package main

import (
	"fmt"
	"import_demo/show"
	"os"
)

func main() {
	fmt.Println(os.Getenv("GOPATH"))
	fmt.Println(os.Getenv("GO111MODULE"))
	fmt.Println(show.Show())
}
