package wordle

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strings"
)

var STARTERS = []string{"raise", "store", "crane", "notes", "acrid"}

type Solver struct {
	game         *Game
	guessCounter int
	lastResult   []int
}

func NewSolver(game *Game) Solver {
	return Solver{game, 0, nil}
}

func (s *Solver) Solve() (string, int, error) {
	words := s.game.dictionary
	for {
		guess := s.getGuess(words)
		s.guessCounter++
		fmt.Printf(" > trying %s ", guess)
		result, err := s.game.CheckGuess(guess)
		if err != nil {
			log.Fatalln(err)
		}

		s.lastResult = result
		fmt.Printf("- result from guess is %v\n", result)
		if s.isSolved(result) {
			return guess, s.guessCounter, nil
		}

		if s.guessCounter >= 20 {
			return "", s.guessCounter, errors.New("\nquitting after too many failed attempts\n")
		}

		words = s.reduceWords(guess, result, words)
	}

}

func (s *Solver) getGuess(words []string) string {
	if s.lastResult == nil {
		return STARTERS[rand.Intn(len(STARTERS))]
	}

	// TODO: something better
	index := rand.Intn(len(words))
	return words[index]
}

func (s *Solver) reduceWords(guess string, result []int, words []string) []string {
	for pos, code := range result {
		k := 0
		if code == NO_MATCH {
			for _, v := range words {
				if !strings.ContainsRune(v, rune(guess[pos])) {
					words[k] = v
					k++
				}
			}
			words = words[:k]
		} else if code == CORRECT {
			for _, v := range words {
				if v[pos] == guess[pos] {
					words[k] = v
					k++
				}
			}
			words = words[:k]
		} else if code == OUT_OF_POSITION {
			for _, v := range words {
				if strings.ContainsRune(v, rune(guess[pos])) {
					words[k] = v
					k++
				}
			}
			words = words[:k]
		}

	}
	return words
}

func (s *Solver) isSolved(result []int) bool {
	for _, val := range result {
		if val != CORRECT {
			return false
		}
	}
	return true
}
