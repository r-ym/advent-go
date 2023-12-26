package main

import (
	"advent-go/utils"
	"fmt"
	"strconv"
	"regexp"
)

func main() {
	var input, _ = util.ReadURL("https://adventofcode.com/2023/day/1/input")

	// part1(input)
	part2(input)
}

func part1(input []string) {
	running_total := 0
	for _, line := range input {
		flag := false
		var last int
		for _, char := range line {
			val, err := strconv.Atoi(string(char))
			if err == nil {
				if !flag {
					running_total += val * 10
					last = val
					flag = true
				} else {
					last = val
				}
			}
		}
		running_total += last
	}
	fmt.Println(running_total)
}

func part2(input []string) {
	// Define the regular expression pattern
	re := regexp.MustCompile(`[1-9]|(one|two|three|four|five|six|seven|eight|nine)`)

	// define map for string to int conversion
	str_to_int := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	running_total := 0
	for _, line := range input {

		// Find all matches of the pattern in the input string
		matches := re.FindAllString(line, -1)

		first := matches[0]
		val1, err1 := strconv.Atoi(first)
		if err1 == nil {
			running_total += val1 * 10
		} else {
			running_total += str_to_int[first] * 10
		}

		last := matches[len(matches)-1]

		val2, err2 := strconv.Atoi(last)
		if err2 == nil {
			running_total += val2
		} else {
			running_total += str_to_int[last]
		}

	}
	fmt.Println(running_total)
}
