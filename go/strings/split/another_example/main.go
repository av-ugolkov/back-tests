package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	score1 := part1()
	fmt.Println(score1)

	score2 := part2()
	fmt.Println(score2)
}

func part1() int {
	sumOfGameIndices := 0

	inputLines := []string{
		"Game 1: 4 blue, 7 red, 5 green; 3 blue, 4 red, 5 green; 3 red, 11 green",
		"Game 2: 2 blue, 8 red, 1 green; 15 blue, 2 green, 8 red;",
		"Game 3: 1 red, 6 green, 4 blue",
	}

	for gameIndex, line := range inputLines {

		// Remove the "Game x:" prefix
		parts := strings.Split(line, ": ")
		game := parts[1]

		// Process each round of the current game
		rounds := strings.Split(game, "; ")

		isPossible := false
		for _, round := range rounds {
			red, green, blue := findScores(round)
			if isRoundPossible(red, green, blue) {
				isPossible = true
			} else {
				isPossible = false

				// If only one round is not possible, there is
				// no need to process the other ones.
				break
			}
		}
		if isPossible {
			sumOfGameIndices += gameIndex + 1
		}
	}
	return sumOfGameIndices
}

func findScores(round string) (int, int, int) {

	// initialize the red, green, blue values to 0
	var red, green, blue int

	individualScores := strings.Split(round, ", ")
	for _, score := range individualScores {
		scoreWithColor := strings.Split(score, " ")
		scoreStr := scoreWithColor[0]
		color := scoreWithColor[1]

		// convert score to integer
		score, _ := strconv.Atoi(scoreStr)
		if color == "red" {
			red += score
		}
		if color == "green" {
			green += score
		}
		if color == "blue" {
			blue += score
		}
	}
	return red, blue, green
}

func isRoundPossible(red, green, blue int) bool {
	return red <= 12 && green <= 13 && blue <= 14
}

func part2() int {
	sumOfGameIndices := 0

	inputLines := []string{
		"Game 1: 4 blue, 7 red, 5 green; 3 blue, 4 red, 5 green; 3 red, 11 green",
		"Game 2: 2 blue, 8 red, 1 green; 15 blue, 2 green, 8 red;",
		"Game 3: 1 red, 6 green, 4 blue",
	}

	for gameIndex, line := range inputLines {
		// Remove the "Game x:" prefix
		gameContentIndex := strings.Index(line, ": ") + 2
		game := line[gameContentIndex:]

		// Process each round of the current game
		rounds := strings.Split(game, "; ")

		isPossible := false
		for _, round := range rounds {
			red, green, blue := findScores2(round)
			if isRoundPossible2(red, green, blue) {
				isPossible = true
			} else {
				isPossible = false

				// If only one round is not possible, there is
				// no need to process the other ones.
				break
			}
		}
		if isPossible {
			sumOfGameIndices += gameIndex + 1
		}
	}
	return sumOfGameIndices
}

func findScores2(round string) (int, int, int) {
	var red, green, blue int

	// end of content
	var eoc int

	// while we haven't reached the end of the round
	for eoc != -1 {
		// get the score with the color
		var currentScoreWithColor string
		eoc = strings.Index(round, ", ")
		if eoc == -1 {
			currentScoreWithColor = round
		} else {
			currentScoreWithColor = round[:eoc]
			round = round[eoc+2:]
		}

		// and assign that score to the matching
		// variable
		space := strings.Index(currentScoreWithColor, " ")
		scoreStr := currentScoreWithColor[:space]
		color := currentScoreWithColor[space+1:]

		// conversion logic
		score, _ := strconv.Atoi(scoreStr)
		if color == "red" {
			red += score
		}
		if color == "green" {
			green += score
		}
		if color == "blue" {
			blue += score
		}
	}
	return red, blue, green
}

func isRoundPossible2(red, green, blue int) bool {
	return red <= 12 && green <= 13 && blue <= 14
}
