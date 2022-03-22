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
		_, attempts, err := solver.Solve()
		if err != nil {
			t.Fatal("error trying to solve", err)
		}

		data = append(data, float64(attempts))
	}

	fmt.Printf("Done!\n Total: %d\n", TEST_RUNS)
	mean, _ := stats.Mean(data)
	max, _ := stats.Max(data)
	min, _ := stats.Min(data)
	fmt.Printf(" Mean: %.2f, Max: %.0f, Min: %.0f\n", mean, max, min)
}
