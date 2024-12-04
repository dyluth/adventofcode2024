package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var test = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

func Test_isLineOK(t *testing.T) {

	part2 = false
	require.Equal(t, true, isLineOK(GetInput("7 6 4 2 1")[0]))
	require.Equal(t, false, isLineOK(GetInput("1 2 7 8 9")[0]))
	require.Equal(t, false, isLineOK(GetInput("9 7 6 2 1")[0]))
	require.Equal(t, false, isLineOK(GetInput("1 3 2 4 5")[0]))
	require.Equal(t, false, isLineOK(GetInput("8 6 4 4 1")[0]))
	require.Equal(t, true, isLineOK(GetInput("1 3 6 7 9")[0]))

	part2 = true
	require.Equal(t, true, isLineOK(GetInput("7 6 4 2 1")[0]))
	require.Equal(t, false, isLineOK(GetInput("1 2 7 8 9")[0]))
	require.Equal(t, false, isLineOK(GetInput("9 7 6 2 1")[0]))
	require.Equal(t, true, isLineOK(GetInput("1 3 2 4 5")[0]))
	require.Equal(t, true, isLineOK(GetInput("8 6 4 4 1")[0]))
	require.Equal(t, true, isLineOK(GetInput("1 3 6 7 9")[0]))

}

func Test_countSafeLines(t *testing.T) {
	part2 = false
	require.Equal(t, 2, countSafeLines(GetInput(test)))
	part2 = true
	require.Equal(t, 4, countSafeLines(GetInput(test)))

}
