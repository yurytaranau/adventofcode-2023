package days

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type cubeLimit struct {
	Blue  int
	Red   int
	Green int
}

func possibleGameDetector(s string, limit cubeLimit) int {
	var id int
	var games string
	split := strings.Split(s, ":")
	if len(split) >= 2 {
		id, games = parseGameID(split[0]), split[1]
	} else {
		return 0
	}
	gamesSplit := strings.Split(games, ";")
	for _, game := range gamesSplit {
		valid := checkLimits(game, limit)
		if !valid {
			return 0
		}
	}
	return id
}

func checkLimits(game string, limit cubeLimit) bool {
	blue, red, green := parseCubes(game)
	if blue+red+green > limit.Blue+limit.Green+limit.Red {
		return false
	} else if blue > limit.Blue {
		return false
	} else if red > limit.Red {
		return false
	} else if green > limit.Green {
		return false
	}
	return true
}

func parseCubes(game string) (blue, red, green int) {
	cubeSplit := strings.Split(game, ", ")
	if len(cubeSplit) == 0 {
		return
	}
	for _, cube := range cubeSplit {
		split := strings.Split(strings.TrimSpace(cube), " ")
		if len(split) == 0 {
			return
		}
		switch split[1] {
		case "blue":
			blue = parseInt(split[0])
		case "red":
			red = parseInt(split[0])
		case "green":
			green = parseInt(split[0])
		}
	}
	return
}

func parseGameID(s string) int {
	split := strings.Split(s, " ")
	if len(split) >= 2 {
		v, err := strconv.Atoi(split[1])
		if err != nil {
			return 0
		}
		return v
	}
	return 0
}

func parseInt(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return v
}

// func calibrationValue(ss []string) int {
// 	digits := []int{}
// 	for _, char := range ss {
// 		v, err := strconv.Atoi(char)
// 		if err == nil {
// 			digits = append(digits, v)
// 		}
// 	}
// 	cv := ""
// 	if len(digits) == 1 {
// 		cv = fmt.Sprintf("%d%d", digits[0], digits[0])
// 	} else {
// 		cv = fmt.Sprintf("%d%d", digits[0], digits[len(digits)-1])
// 	}
// 	cvi, err := strconv.Atoi(cv)
// 	if err != nil {
// 		return 0
// 	}
// 	return cvi
// }

// func convertDigits(s string) string {
// 	r := strings.NewReplacer(
// 		"one", "o1e",
// 		"two", "t2o",
// 		"three", "t3e",
// 		"four", "f4r",
// 		"five", "f5e",
// 		"six", "s6x",
// 		"seven", "s7n",
// 		"eight", "e8t",
// 		"nine", "n9e",
// 	)
// 	res := s
// 	for i := 1; i < 9; i++ {
// 		res = r.Replace(res)
// 	}
// 	return res
// }

func (d Days) Day2p1(source io.Reader) (cvs int) {
	fs := bufio.NewScanner(source)
	sum := 0
	for fs.Scan() {
		sum += possibleGameDetector(fs.Text(), cubeLimit{Blue: 14, Red: 12, Green: 13})
	}
	fmt.Println("(Part 1) sum of possible game IDs\n", sum)
	return
}

func (d Days) Day2p2(source io.Reader) (cvs int) {
	fs := bufio.NewScanner(source)
	for fs.Scan() {
		line := fs.Text()
		convertedLine := convertDigits(line)
		cvs += calibrationValue(strings.Split(convertedLine, ""))
	}
	fmt.Println("(Part 2) sum of calibration values\n", cvs)
	return
}
