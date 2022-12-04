package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Info().Msg("Day 4")
	file, err := os.Open("day4.txt")
	if err != nil {
		log.Info().Err(err).Send()
	}

	defer file.Close()

	var overlappingAssignments int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if processLinePairs(line) {
			overlappingAssignments++
		}
	}

	log.Info().Int("total", overlappingAssignments).Send()
}

func processLinePairs(pairs string) bool {
	splitStr := strings.Split(pairs, ",")
	elf1Min, elf1Max := getMinMax(splitStr[0])
	elf2Min, elf2Max := getMinMax(splitStr[1])

	return isOverlapping(elf1Min, elf1Max, elf2Min, elf2Max)
}

func isOverlapping(elf1Min, elf1Max, elf2Min, elf2Max int) bool {
	if elf1Min >= elf2Min && elf1Min <= elf2Max {
		log.Info().Interface("ranges", map[string]string{
			"elf1": fmt.Sprintf("%d-%d", elf1Min, elf1Max),
			"elf2": fmt.Sprintf("%d-%d", elf2Min, elf2Max),
		}).Send()
		return true
	} else if elf2Min >= elf1Min && elf2Min <= elf1Max {
		log.Info().Interface("ranges", map[string]string{
			"elf1": fmt.Sprintf("%d-%d", elf1Min, elf1Max),
			"elf2": fmt.Sprintf("%d-%d", elf2Min, elf2Max),
		}).Send()
		return true
	}

	return false
}

func isFullyOverlapping(elf1Min, elf1Max, elf2Min, elf2Max int) bool {
	if elf1Min <= elf2Min && elf1Max >= elf2Max {
		return true
	} else if elf2Min <= elf1Min && elf2Max >= elf1Max {
		return true
	}

	return false

}

func getMinMax(assignents string) (min, max int) {
	results := strings.Split(assignents, "-")

	min, _ = strconv.Atoi(results[0])
	max, _ = strconv.Atoi(results[1])

	return
}
