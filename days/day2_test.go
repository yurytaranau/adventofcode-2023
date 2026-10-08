package days

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestCalibrationValue(t *testing.T) {
// 	input := []string{"1", "a", "b", "c", "2"}
// 	res := calibrationValue(input)
// 	assert.Equal(t, 12, res, input)
// }

// func TestConvertDigits(t *testing.T) {
// 	input := "two1nine"
// 	res := convertDigits(input)
// 	assert.Equal(t, "t2o1n9e", res, input)
// }

func TestDay2_Part1(t *testing.T) {
	d := Days{}
	input := `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
	Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
	Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
	Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
`
	res := d.Day1p1(strings.NewReader(input))
	assert.Equal(t, 8, res, input)
}

func TestDay2_Part2(t *testing.T) {
	d := Days{}
	input := `two1nine
	eightwothree
	abcone2threexyz
	xtwone3four
	4nineeightseven2
	zoneight234
	7pqrstsixteen
`
	res := d.Day1p2(strings.NewReader(input))
	assert.Equal(t, 281, res, input)
}
