package main

type player struct {
	fname string
	lname string
	score int
}

type game struct {
	players map[string]*player
}

func (g *game) getWinner() *player {
	maxScore := 0
	var winner *player

	for _, i := range g.players {
		if i.score > maxScore {
			maxScore = i.score
			winner = i
		}
	}
	return winner
}

func (g game) printOK() {
	println("ok")
}

type tryInher struct {
	game
}

func main() {
	// g := game{
	// 	players: make(map[string]*player),
	// }
	// fmt.Println(g)

	// p1 := player{}
	// fmt.Printf("p1: %v\n", p1)
	// g.players["p1"] = &p1

	// var p2 player
	// p2.fname = "hesham"
	// p2.lname = "ali"
	// p2.score = 45
	// g.players["p2"] = &p2
	// fmt.Printf("p2: %v\n", p2)

	// p3 := player{
	// 	fname: "amr",
	// 	lname: "soliman",
	// 	score: 55,
	// }
	// g.players["p3"] = &p3
	// fmt.Printf("p3: %v\n", p3)
	// fmt.Printf("g: %v\n", g)

	// fmt.Printf("g.players[\"p2\"]: %v\n", *g.players["p2"])

	// winner := g.getWinner()
	// fmt.Printf("winner: %v\n", winner)

	//try inhertence
	test := tryInher{}
	test.printOK()
}
