package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_lineToints(t *testing.T) {
	lines, rules := getInput(testInput)

	require.True(t, isLineOK(lines[0], rules))
	require.True(t, isLineOK(lines[1], rules))
	require.True(t, isLineOK(lines[2], rules))
	require.False(t, isLineOK(lines[3], rules))
	require.False(t, isLineOK(lines[4], rules))
	require.False(t, isLineOK(lines[5], rules))

}

func Test_fixLine(t *testing.T) {
	lines, rules := getInput(testInput)
	fmt.Printf("lines3: %+v\n", lines[3])
	fixLine(lines[3], rules)
	fmt.Printf("lines3: %+v\n", lines[3])
	require.True(t, fixLine(lines[3], rules))
}

func Test_getTotalScoreofCorrectLines(t *testing.T) {
	require.Equal(t, 143, getTotalScoreofCorrectLines(getInput(testInput)))
}

func Test_getTotalScoreANDCorrectLines(t *testing.T) {
	require.Equal(t, 123, getTotalScoreANDCorrectLines(getInput(testInput)))
}

var testInput = `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

func Test_getLineScore(t *testing.T) {
	require.Equal(t, 3, getLineScore([]int{1, 2, 3, 4, 5}))
	require.Equal(t, 1, getLineScore([]int{0, 1, 9}))
}
