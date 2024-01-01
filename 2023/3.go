package main

import (
	util "advent-go/utils"
	"fmt"
	"strconv"

	// "strings"
	mapset "github.com/deckarep/golang-set/v2"
	// "regexp"
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

var special_set = mapset.NewSet[Coord]()

func part1(input []string) {
	ans := 0
	bound_x := len(input[0])
	bound_y := len(input)

	for i, line := range input {
		num_string := ""
		for j, char := range line {
			if char == '.' && num_string == "" {
				continue

			} else if char == '.'{
				length := len(num_string)
				coords := generate_surrounding_coords(j, i, length, bound_y, bound_x)

				special_char := check_coords(input, coords, i)

				if special_char {
					val := convert_to_int(num_string)
					fmt.Println("YOO", val, ans)
					ans += val
				}

				num_string = ""



			} else if unicode.IsDigit(char) {
				num_string += string(char)
				if j == bound_x-1 {
					length := len(num_string)
					coords := generate_surrounding_coords(j, i, length, bound_y, bound_x)
					special_char := check_coords(input, coords, i)
					if special_char {
						val := convert_to_int(num_string)
						fmt.Println("YOO", val, ans)
						ans += val
					}
				}
			} else if num_string == "" {
				special_set.Add(Coord{j, i})
			} else {
				special_set.Add(Coord{j, i})
				val := convert_to_int(num_string)
				fmt.Println("YOO", val, ans)
				ans += val
				num_string = ""
			}

		}
		// break
	}

	fmt.Println("Part 1:", ans)

}

func part2(input []string) {
	ans := 0
	// for _, line := range input {

	// }
	fmt.Println("Part 2: ", ans)

}

func generate_surrounding_coords(x2 int, y2 int, length int, bound_y int, bound_x int) []Coord {
	coords := []Coord{}
	y_up := (y2 - 1) > 0
	y_down := (y2 + 1) < bound_y
	for x := x2 - (length + 1); x < x2; x++ {

		if x < 0 || x >= bound_x {
			continue

		}
		if x == x2-(length+1) {
			coords = append(coords, Coord{x2 - (length + 1), y2})
		}

		if !y_up && !y_down {
			break
		}

		if y_up {
			coords = append(coords, Coord{x, y2 - 1})
		}
		if y_down {
			coords = append(coords, Coord{x, y2 + 1})
		}
	}
	return coords
}

func convert_to_int(str string) int {
	ans, _ := strconv.Atoi(str)
	return ans
}

func check_coords(input []string, coords []Coord, i int) bool {
	flag := false
	for _, coord := range coords {
		if special_set.Contains(coord) {
			flag = true
			break
		} else if coord.Y > i {
			val := input[coord.Y][coord.X]
			if val == '.' || unicode.IsDigit(rune(val)) {
				continue
			} else {
				special_set.Add(coord)
				flag = true
				break
			}
		}
	}
	return flag
}
