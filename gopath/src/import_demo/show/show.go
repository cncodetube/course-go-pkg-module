package show

import "fmt"

var IsShow = true

const Name = "SHOW NAME"

func Show() string {
	return Name
}

func init() {
	fmt.Println("show::init:: 0")
}