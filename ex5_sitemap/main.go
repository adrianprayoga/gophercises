package main

import (
	"fmt"
	"os"

	link "github.com/adrianprayoga/gophercises/link"
)

func main() {
	fmt.Println("starting")
	r, _ := os.Open("../ex4_link/htmls/ex1.html")
	fmt.Println(link.GetLinks(r))
}
