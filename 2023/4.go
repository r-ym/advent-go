package main

import (
	util "advent-go/utils"
	"fmt"
	"math"
	"strconv"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
)

func main() {
	var input, _ = util.ReadURL("https://adventofcode.com/2023/day/4/input")

	part1(input)
	part2(input)
}

func part1(input []string) {
	ans := 0
	for _, line := range input {
		count := parse_card_win_count(line)
		if count != 0 {
			fmt.Println(count, int(math.Pow(2, float64(count-1))))
			ans += int(math.Pow(2, float64(count-1)))
		}
		break
	}
	fmt.Println("Part 1:", ans)
}

func part2(input []string) {
	ans := 0
	for _, line := range input {
		break
		fmt.Println(line)
	}
	fmt.Println("Part 2: ", ans)

}

func parse_card_win_count(card string) int {
	game := strings.Split(card, ":")[1]
	card_split := strings.Split(strings.TrimSpace(game), "|")

	win_set := mapset.NewSet[int]()
	// iterate through card_split[0] after splitting at " " and add to win_set after converting to int
	for _, card := range strings.Split(card_split[0], " ") {
		val, _ := strconv.Atoi(card)
		win_set.Add(val)
	}

	fmt.Println(win_set)
	count := 0
	for _, card := range strings.Split(card_split[1], " ") {
		val, _ := strconv.Atoi(card)
		if win_set.Contains(val) {
			count++
		}
	}
	return count
}
