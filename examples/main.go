package main

import (
	"fmt"
	"github.com/foruscommunity/gopincode"
)

func main() {
	c, s := gopincode.GenSix()

	fmt.Printf("int: %d\n", c)
	fmt.Printf("string: %s\n", s)
}
