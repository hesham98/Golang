package main

import "fmt"

func main() {

	age := 25
	pAge := &age
	fmt.Println(*pAge)

	name := "Main"
	fmt.Println("the main value", name)
	updateCopy(name)
	fmt.Println("after update Copy", name)
	updateRef(&name)
	fmt.Println("after update Refernce", name)
}

func updateCopy(s string) {
	s = "Copy"
}

func updateRef(s *string) {
	fmt.Println("address is ", s)
	*s = "Ref"
}
