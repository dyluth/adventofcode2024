package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

var test = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

func Test_getAllMatches(t *testing.T) {

	got := getAllMatches(getInput(test))
	fmt.Printf("matches: %v\n", got)
	require.Equal(t, 18, got)
}

func Test_getAllMatchesP2(t *testing.T) {

	got := getAllMatchesP2(getInput(test))
	fmt.Printf("matches: %v\n", got)
	require.Equal(t, 9, got)
}
