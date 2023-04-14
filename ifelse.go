package main

func main() {
	// name := "hesham"
	// if name == "hesham" {
	// 	println("hello", name)
	// } else if name == "mohamed" {
	// 	println("hello", name)
	// } else {
	// 	println("hello", name)
	// }
	// age := 26
	// if age > 24 {
	// 	println("you are less than 24 years old")
	// } else if age == 24 {
	// 	fmt.Println("you are", age, "years old")
	// } else {
	// 	fmt.Println("you are older than 24 years old")
	// }

	// age := 17
	// name := "hesham"
	// if name == "hesham" && age >= 18 {
	// 	println("you can drive the car mr.", name)
	// } else {
	// 	println("you can not drive the car mr.", name)
	// }

	// 	age := 24
	// 	name := "heshaam"

	// 	if age == 24 {
	// 		if name == "hesham" {
	// 			println("you can drive this car mr", name)
	// 		} else {
	// 			println("you can not drive this car mr", name, "because of you name is not hesham")
	// 		}
	// 	} else {
	// 		println("you can not drive this car mr", name, "because of you age not = 24")
	// 	}

	// age := 17
	// name := "hesham"
	// if name == "hesham" || age >= 18 {
	// 	println("you can drive the car mr.", name)
	// } else {
	// 	println("you can not drive the car mr.", name)
	// }

	// age := 0
	// switch age {
	// case 0:
	// 	println("can not be")
	// case 1:
	// 	println("he is", age, "years old")
	// case 2:
	// 	println("he is", age, "years old")
	// default:
	// 	println("he is", age, "years old")
	// }

	age := 55
	switch {
	case age > 18:
		println("he is older than 18 years old")
		fallthrough
	case age > 51:
		println("he is older than 51 years old")
	case age == 55:
		println("he is 55 years old")
	case age != 66:
		println("he is 66 years old")
	default:
		println("idk")
	}

}
