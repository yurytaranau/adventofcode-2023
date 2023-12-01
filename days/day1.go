package days

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func calibrationValue(ss []string) int {
	digits := []int{}
	for _, char := range ss {
		v, err := strconv.Atoi(char)
		if err == nil {
			digits = append(digits, v)
		}
	}
	cv := ""
	if len(digits) == 1 {
		cv = fmt.Sprintf("%d%d", digits[0], digits[0])
	} else {
		cv = fmt.Sprintf("%d%d", digits[0], digits[len(digits)-1])
	}
	cvi, err := strconv.Atoi(cv)
	if err != nil {
		return 0
	}
	return cvi
}

func convertDigits(s string) string {
	r := strings.NewReplacer(
		"one", "o1e",
		"two", "t2o",
		"three", "t3e",
		"four", "f4r",
		"five", "f5e",
		"six", "s6x",
		"seven", "s7n",
		"eight", "e8t",
		"nine", "n9e",
	)
	res := s
	for i := 1; i < 9; i++ {
		res = r.Replace(res)
	}
	return res
}

func (d Days) Day1p1(source io.Reader) (cvs int) {
	fs := bufio.NewScanner(source)
	for fs.Scan() {
		cvs += calibrationValue(strings.Split(fs.Text(), ""))
	}
	fmt.Println("(Part 1) sum of calibration values\n", cvs)
	return
}

func (d Days) Day1p2(source io.Reader) (cvs int) {
	fs := bufio.NewScanner(source)
	for fs.Scan() {
		line := fs.Text()
		convertedLine := convertDigits(line)
		cvs += calibrationValue(strings.Split(convertedLine, ""))
	}
	fmt.Println("(Part 2) sum of calibration values\n", cvs)
	return
}
