package main

import (
	"fmt"
)

var x = 42
var y = "James Bond"
var z = true

type s string

func main() {

	s := fmt.Sprintf("%v %v  %v", x, y, z)
	println(s)

}
