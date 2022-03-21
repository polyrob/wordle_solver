package main

import (
	"fmt"
	"testing"

	"github.com/montanaflynn/stats"
	"github.com/polyrob/gosandbox/wordle"
)

const TEST_RUNS = 1000

func TestSolve(t *testing.T) {

	game := wordle.NewGame(DICT_PATH, 5)
	data := []float64{}

	for i := 0; i < TEST_RUNS; i++ {
		game.NewTarget()
		solver := wordle.NewSolver(&game)
		_, attemps, err := solver.Solve()
		if err != nil {
			t.Fatal("error trying to solve", err)
		}
		data = append(data, float64(attemps))
	}

	fmt.Printf("Done!\n Total: %d\n", TEST_RUNS)
	mean, _ := stats.Mean(data)
	fmt.Printf(" Mean: %.2f\n", mean)
	max, _ := stats.Max(data)
	fmt.Printf(" Max: %.0f\n", max)
	min, _ := stats.Min(data)
	fmt.Printf(" Min: %.0f\n", min)
}
