package main

import (
	"fmt"
	"os"

	"github.com/adrianprayoga/gophercises/ex4_link"
)

func main() {
	fmt.Println("starting")
	r, _ := os.Open("../ex4_link/htmls/ex1.html")
	fmt.Println(ex4_link.GetLinks(r))
}
