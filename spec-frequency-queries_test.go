package main

import "testing"

// expected output [0, 1]
func TestFrequencyQueries1(t *testing.T) {
	queries := [][]int32{
		[]int32{1, 5},
		[]int32{1, 6},
		[]int32{3, 2},
		[]int32{1, 10},
		[]int32{1, 10},
		[]int32{1, 6},
		[]int32{2, 5},
		[]int32{3, 2},
	}

	freqQuery(queries)
}

// expected output [0, 1, 1]
func TestFrequencyQueries2(t *testing.T) {
	queries := [][]int32{
		[]int32{1, 3},
		[]int32{2, 3},
		[]int32{3, 2},
		[]int32{1, 4},
		[]int32{1, 5},
		[]int32{1, 5},
		[]int32{1, 4},
		[]int32{3, 2},
		[]int32{2, 4},
		[]int32{3, 2},
	}

	freqQuery(queries)
}

func TestFrequencyQueries3(t *testing.T) {
	queries := [][]int32{
		[]int32{2, 3},
		[]int32{2, 3},
		[]int32{2, 3},
		[]int32{1, 3},
		[]int32{2, 3},
		[]int32{2, 3},
		[]int32{2, 3},
		[]int32{1, 5},
		[]int32{1, 5},
		[]int32{1, 4},
		[]int32{3, 2},
		[]int32{2, 4},
		[]int32{3, 2},
	}

	freqQuery(queries)
}
