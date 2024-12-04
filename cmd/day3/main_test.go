package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var input2 = `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`

var input3 = `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`

func TestGetInput(t *testing.T) {
	part2 = false
	total := getAllMuls(GetInput(input2))
	require.Equal(t, 161, total)

	do = true
	part2 = true
	total = getAllMuls(GetInput(input3))
	require.Equal(t, 48, total)

}
