package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

var test = `3   4
4   3
2   5
1   3
3   9
3   3`

func TestSortInput(t *testing.T) {
	got, got1 := SortInput(test)
	require.Len(t, got, 6)
	require.Len(t, got1, 6)
	require.Equal(t, 1, got[0])
	require.Equal(t, 3, got1[0])
	require.Equal(t, 3, got1[1])

	fmt.Printf("%+v\n%+v\n", got, got1)
}

func Test_getDiffs(t *testing.T) {

	got := getDiffs(SortInput(test))
	fmt.Printf("%+v\n", got)

	require.Equal(t, 11, total(got))
}

func Test_multiplyInstances(t *testing.T) {
	got := multiplyInstances(SortInput(test))
	fmt.Printf("%+v\n", got)

	require.Equal(t, 9, got[0])
	require.True(t, false)

}
