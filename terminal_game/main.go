package main

import (
	"fmt"
	"math"
	"math/rand"
)

type Player struct {
	name  string
	guess int
}

type Game struct {
	player1 Player
	player2 Player
	random  int
}

// Create a Player
func (p *Player) createPlayer(name string) {
	p.name = name
}

// Create Game
func (g *Game) createGame() {
	g.player1.createPlayer("Player 1")
	g.player2.createPlayer("Player 2")
	g.random = rand.Intn(100) + 1
}

// Read Input
func (p *Player) readInput() {
	var input int
	fmt.Printf("%s, Enter your guess: ", p.name)
	fmt.Scanf("%d", &input)
	p.guess = input
}

// Create Game Loop
func (g *Game) gameLoop() {
	// loop forever
	for {
		// Read player inputs
		g.player1.readInput()
		g.player2.readInput()

		// check who is closer
		p1Diff := math.Abs(float64(g.player1.guess - g.random))
		p2Diff := math.Abs(float64(g.player2.guess - g.random))

		if p1Diff < p2Diff {
			fmt.Printf("%s wins!\n", g.player1.name)
		} else if p1Diff > p2Diff {
			fmt.Printf("%s wins!\n", g.player2.name)
		} else {
			fmt.Println("It's a Tie!")
		}

		// check if player wants to play again
		var playAgain string
		fmt.Printf("%s, would you like to play again? (y/n): ", g.player1.name)
		fmt.Scanf("%s", &playAgain)

		if playAgain == "n" {
			break
		}
	}

	fmt.Println("Thanks for playing!")
}

func main() {
	game := Game{}
	game.createGame()
	game.gameLoop()
}
