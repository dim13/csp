package main

import (
	"flag"
	"fmt"

	"github.com/dim13/csp"
)

func main() {
	var i []int
	var s []string

	flag.Var(csp.Int{&i}, "i", "comma separated integers")
	flag.Var(csp.String{&s}, "s", "comma separated strings")
	flag.Parse()

	fmt.Println("string", s)
	fmt.Println("int", i)
}
