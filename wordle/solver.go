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
	words := make([]string, len(s.game.dictionary))
	copy(words, s.game.dictionary)
	for {
		guess := s.getGuess(words)
		s.guessCounter++
		fmt.Printf(" > trying %s ", guess)
		result, err := s.game.CheckGuess(guess)
		if err != nil {
			log.Fatalln(err)
		}

		s.lastResult = result
		fmt.Printf("- result: %v", result)
		if s.isSolved(result) {
			fmt.Println("- SOLVED!")
			return guess, s.guessCounter, nil
		}

		if s.guessCounter >= 20 {
			return "", s.guessCounter, errors.New("too many failed attempts")
		}

		words = s.reduceWords(guess, result, words)
		fmt.Printf("- possible words: %d\n", len(words))
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
				// must have this rune BUT not in the same location
				if strings.ContainsRune(v, rune(guess[pos])) && v[pos] != guess[pos] {
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
