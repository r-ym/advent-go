package main

import (
	"advent-go/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var input, _ = util.ReadURL("https://adventofcode.com/2023/day/2/input")

	part1(input)
	part2(input)
}

func part1(input []string) {
	r_max := 12
	g_max := 13
	b_max := 14
	ans := 0
	for _, game := range input {
		r, g, b := parse_game_max(game)
		if r <= r_max && g <= g_max && b <= b_max {
			game_id := parse_game_id(game)
			ans += game_id
		}
	}
	fmt.Println(ans)
}

func part2(input []string) {
	running_total := 0
	for _, game := range input {
		r, g, b := parse_game_max(game)
		running_total += r * g * b
	}
	fmt.Println(running_total)

}

func parse_game_max(game string) (int, int, int) {
	r, g, b := 0, 0, 0
	game_info := strings.Split(game, ":")[1]
	game_info_split := strings.Split(strings.TrimSpace(game_info), ";")
	for _, info := range game_info_split {
		info_split := strings.Split(strings.TrimSpace(info), ",")
		for _, color := range info_split {
			color_split := strings.Split(strings.TrimSpace(color), " ")
			color_val, _ := strconv.Atoi(color_split[0])
			switch color := color_split[1]; color {
			case "red":
				r = max(r, color_val)
			case "green":
				g = max(g, color_val)
			case "blue":
				b = max(b, color_val)
			}
		}
	}
	return r, g, b
}

func parse_game_id(game string) int {
	game_id_str := strings.Split(game, ":")[0][5:]
	game_id, _ := strconv.Atoi(game_id_str)
	return game_id
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
