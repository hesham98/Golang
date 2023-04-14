package main

func main() {
	// string
	// fmt.Println("hello", "Hesham", "\n24 years old")

	//int
	// fmt.Println(1)
	// fmt.Println(1, 2)
	// fmt.Println(4 / 2)
	// fmt.Println(3 / 4)

	//string with int
	// fmt.Println("Hesham", 24, "years old")
	// println("hesham", 22)

	// float
	// fmt.Println("float numbers", 3.0/4.0)
	// println(7.0 / 2)

	//bool
	// fmt.Println(true)
	// fmt.Println(false)

	//var declation
	// var name string = "hesham"
	// fmt.Println("hello", name)

	// age := 25
	// println(age)

	// const game = "XO"
	// fmt.Println(game)

	//struct
	// type person struct {
	// 	name string
	// 	age  int
	// }

	// p1 := person{name: "hesham", age: 24}
	// println(p1.name, p1.age)

	//channels
	messages := make(chan string)
	go func() { messages <- "Hi" }()
	msg := <-messages
	println(msg)
}
