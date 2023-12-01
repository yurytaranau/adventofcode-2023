package days

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalibrationValue(t *testing.T) {
	input := []string{"1", "a", "b", "c", "2"}
	res := calibrationValue(input)
	assert.Equal(t, 12, res, input)
}

func TestConvertDigits(t *testing.T) {
	input := "two1nine"
	res := convertDigits(input)
	assert.Equal(t, "t2o1n9e", res, input)
}

func TestDay1_Part1(t *testing.T) {
	d := Days{}
	input := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
`
	res := d.Day1p1(strings.NewReader(input))
	assert.Equal(t, 142, res, input)
}

func TestDay1_Part2(t *testing.T) {
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
