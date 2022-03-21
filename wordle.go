package main

import (
	"fmt"
	"log"

	"github.com/polyrob/gosandbox/wordle"
)

const DICT_PATH = "/usr/share/dict/words"

func main() {
	game := wordle.NewGame(DICT_PATH, 5)

	solver := wordle.NewSolver(&game)
	result, attemps, err := solver.Solve()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Solved for %s in %d tries.\n", result, attemps)
}
