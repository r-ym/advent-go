package main

import (
	"fmt"
	"advent-go/utils"
	"strconv"
)

func main(){
	var input, _ = util.ReadURL("https://adventofcode.com/2023/day/1/input")
	
	part1(input)
	part2(input)
}

func part1(input []string) {
	running_total := 0
	max := 0
	for _, line := range input {
		if line != "" {
			i, _ := strconv.Atoi(line)
			running_total += i
		} else {
			if running_total > max {
				max = running_total
			}
			running_total = 0
		}
	}
	fmt.Println(max)
}

func part2(input []string) {
	running_total := 0
	top_three_maxes := [3]int{0, 0, 0}
	for _, line := range input {
		if line != "" {
			i, _ := strconv.Atoi(line)
			running_total += i
		} else {
			if running_total > top_three_maxes[0] {
				top_three_maxes[2] = top_three_maxes[1]
				top_three_maxes[1] = top_three_maxes[0]
				top_three_maxes[0] = running_total
			} else if running_total > top_three_maxes[1] {
				top_three_maxes[2] = top_three_maxes[1]
				top_three_maxes[1] = running_total
			} else if running_total > top_three_maxes[2] {
				top_three_maxes[2] = running_total
			}
			running_total = 0
		}
	}
	max_sum := 0
	for _, max := range top_three_maxes{
		max_sum += max
	}
	fmt.Println(max_sum)
}