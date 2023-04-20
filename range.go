package main

import "fmt"

func main() {
	arr := [4]string{"hello", "dear", "hesham", "ali"}

	for key, value := range arr {
		fmt.Printf("key: %v\n", key)
		fmt.Printf("value: %v\n", value)
	}
	slc := []int{11, 22, 33, 44, 55, 66}
	for k, v := range slc {
		fmt.Printf("k: %v\n", k)
		fmt.Printf("v: %v\n", v)
	}

	mp := map[string]string{
		"fname": "hesham",
		"lname": "mohamed",
	}
	for k, v := range mp {
		fmt.Printf("k: %v\n", k)
		fmt.Printf("v: %v\n", v)
	}
}
