package main

import (
	"advent-go/utils"
	"fmt"
	// "strconv"
	// "strings"
	"github.com/deckarep/golang-set/v2"
	"regexp"
	"unicode"
)

func main() {
	var input, _ = util.ReadURL("https://adventofcode.com/2023/day/3/input")

	part1(input)
	part2(input)
}

type Coord struct {
	X, Y int
}

var special_re = regexp.MustCompile(`[^\d.]`)

var special_set = mapset.NewSet[Coord]()

func part1(input []string) {
	ans := 0
	for i, line := range input {
		fmt.Println(line)
		for j, char := range line {
			if char == '.' {
				continue
			} else if unicode.IsDigit(char) {
				continue
			} else {
				special_set.Add(Coord{i, j})
			}
		}
		fmt.Println(special_set)
		break
	}
	fmt.Println("Part 1:", ans)
}

func part2(input []string) {
	ans := 0
	// for _, line := range input {

	// }
	fmt.Println("Part 2: ", ans)

}
