package main

import (
	"bufio"
	"os"
	"regexp"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	re *regexp.Regexp = regexp.MustCompile(`move (?P<count>\d+) from (?P<src>\d+) to (?P<dest>\d+)`)
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Info().Msg("Day 5")
	file, err := os.Open("day5.txt")
	if err != nil {
		log.Info().Err(err).Send()
	}

	defer file.Close()

	var finishedInitial bool
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			finishedInitial = true
		}

		if !finishedInitial {
			var continueCount int
			var column int
			for _, entry := range line {
				if continueCount == 0 {
					if entry == ' ' {
						column++
						continueCount = 3
					} else if entry == '[' {
						column++
					} else {
						log.Info().Int("column", column).Str("entry", string(entry)).Send()
						continueCount = 2
					}
				} else {
					continueCount--
				}
			}
		} else {

		}
	}
}
