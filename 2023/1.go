package main

import (
	"advent-go/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var input, _ = util.ReadURL("https://adventofcode.com/2023/day/1/input")

	part1(input)
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

var int_str = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
}

func part2(input []string) {

	running_total := 0
	for _, line := range input {

		val1 := get_first_digit(line)
		val2 := get_last_digit(line)

		running_total += val1*10 + val2

	}
	fmt.Println(running_total)
}

func get_first_digit(line string) int {
	min := len(line)
	var val int
	for k, _ := range int_str {
		index := strings.Index(line, k)
		if index != -1 {
			if index < min {
				min = index
				val = int_str[k]
			}
		}
	}
	return val
}

func get_last_digit(line string) int {
	max := -1
	var val int
	for k, _ := range int_str {
		index := strings.LastIndex(line, k)
		if index != -1 {
			if index > max {
				max = index
				val = int_str[k]
			}
		}
	}
	return val
}
