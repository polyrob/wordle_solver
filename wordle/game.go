package wordle

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
)

const (
	NO_MATCH        int = 0
	OUT_OF_POSITION     = 1
	CORRECT             = 2
)

type Game struct {
	dictionary []string
	wordSize   int
	targetWord string
}

func NewGame(dictionaryPath string, wordSize int) Game {
	dictionary := fetchWords(dictionaryPath, wordSize)

	rand.Seed(time.Now().Unix())
	targetWord := dictionary[rand.Intn(len(dictionary))]
	fmt.Printf("Target word is %s\n", targetWord)

	return Game{
		dictionary,
		wordSize,
		targetWord,
	}
}

func fetchWords(filepath string, wordSize int) []string {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	toReturn := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if isValidWord(line, wordSize) {
			toReturn = append(toReturn, line)
		}

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return toReturn
}

func isValidWord(word string, wordSize int) bool {
	if len(word) != wordSize {
		return false
	}

	if unicode.IsUpper(rune(word[0])) {
		return false
	}
	return true
}

func (g *Game) CheckGuess(guess string) ([]int, error) {
	if len(guess) != g.wordSize {
		return nil, errors.New("guess doesn't match target word length")
	}

	out := []int{}
	for pos, char := range guess {
		if char == rune(g.targetWord[pos]) {
			out = append(out, CORRECT)
		} else if strings.ContainsRune(g.targetWord, char) {
			out = append(out, OUT_OF_POSITION)
		} else {
			out = append(out, NO_MATCH)
		}
	}
	return out, nil
}
