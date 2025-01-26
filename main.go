package main

import (
	"fmt"

	"github.com/love-yy/puppy"
)

func main() {
	fmt.Println("Hello Gophers 😘😜👀😎✌")

	s1 := puppy.Bark()
	s2 := puppy.Barks()
	fmt.Println(s1)
	fmt.Println(s2)

	s3 := puppy.BigBark()
	fmt.Println(s3)
	s4 := puppy.BigBarks()
	fmt.Println(s4)
}
