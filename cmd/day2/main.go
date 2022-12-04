package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Info().Msg("Hello World")
	file, err := os.Open("day2.txt")
	if err != nil {
		log.Info().Err(err).Send()
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var totalScore int
	for scanner.Scan() {
		line := scanner.Text()
		totalScore += scoreYourGamePartTwo(line)
	}

	log.Info().Int("score", totalScore).Send()
}

func scoreYourGame(game string) int {
	score := 0
	parsedGame := strings.Split(game, " ")

	opponent := parsedGame[0]
	mine := parsedGame[1]

	if mine == "X" {
		score += 1
	} else if mine == "Y" {
		score += 2
	} else if mine == "Z" {
		score += 3
	}

	if (mine == "X" && opponent == "A") ||
		(mine == "Y" && opponent == "B") ||
		(mine == "Z" && opponent == "C") {
		// tie
		score += 3
		log.Info().Str("game", game).Int("score", score).Msg("tie")
	} else if (mine == "X" && opponent == "B") ||
		(mine == "Y" && opponent == "C") ||
		(mine == "Z" && opponent == "A") {
		// lose
		score += 0
		log.Info().Str("game", game).Int("score", score).Msg("lose")
	} else {
		score += 6
		log.Info().Str("game", game).Int("score", score).Msg("win")
	}

	return score
}

func scoreYourGamePartTwo(game string) int {
	score := 0
	parsedGame := strings.Split(game, " ")

	opponent := parsedGame[0]
	mine := parsedGame[1]

	if opponent == "A" { // Opponent Rock
		if mine == "X" { // Mine Scissors
			// lose
			return 3
		} else if mine == "Y" {
			// draw
			return 3 + 1
		} else if mine == "Z" {
			// win
			return 6 + 2
		}
	} else if opponent == "B" { // Opponenet Paper
		if mine == "X" { // Mine Rock
			// lose
			return 1
		} else if mine == "Y" { // Mine Paper
			// draw
			return 3 + 2
		} else if mine == "Z" { // Mine Scissors
			// win
			return 6 + 3
		}
	} else if opponent == "C" { // Oponent Scissors
		if mine == "X" { // Mine Paper
			// lose
			return 2
		} else if mine == "Y" { // Mine Scissors
			// draw
			return 3 + 3
		} else if mine == "Z" { // Mine Rock
			// win
			return 6 + 1
		}
	}

	return score
}
