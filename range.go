package main

import "fmt"

func main() {
	// names := [4]string{"Hesham", "Mohamed", "Ali", "Naga"}
	// for key, value := range names {
	// 	fmt.Println("key =", key, "value =", value)
	// }

	// ages := []int{24, 25, 52, 23}
	// for k, v := range ages {
	// 	ages[2] = 5
	// 	fmt.Println("key =", k, "value =", v)
	// }

	mapOfSlice := map[string][]string{
		"name":    []string{"hesham", "mohamed"},
		"age":     []string{"24"},
		"address": []string{"Egypt", "Portsaid", "Zohor"},
		"phone":   []string{"+00101010101"},
	}
	for ke, va := range mapOfSlice {
		fmt.Println("key =", ke, "value =", va)
		for kee, vaa := range va {
			fmt.Println("key =", kee, "value =", vaa)
		}
	}

}
