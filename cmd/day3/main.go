package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Day 3")
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Info().Msg("Hello World")
	file, err := os.Open("day3.txt")
	if err != nil {
		log.Info().Err(err).Send()
	}

	defer file.Close()

	var sum int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line1 := scanner.Text()
		scanner.Scan()
		line2 := scanner.Text()
		scanner.Scan()
		line3 := scanner.Text()
		// sum += getKnapsackSharedItem(line)
		sum += getGroupSharedItem(line1, line2, line3)
	}

	log.Info().Int("sum", sum).Send()
}

func getGroupSharedItem(sack1, sack2, sack3 string) (sum int) {

	for _, item := range sack1 {
		if strings.ContainsRune(sack2, item) && strings.ContainsRune(sack3, item) {
			sum = getItemPriority(item)
			return
		}
	}
	return
}

func getKnapsackSharedItem(sack string) (sum int) {
	compartment1 := sack[:len(sack)/2]
	compartment2 := sack[len(sack)/2:]

	var found int
	var common []string
	for _, item := range compartment1 {
		if strings.Contains(compartment2, string(item)) {
			found++
			sum += getItemPriority(item)
			common = append(common, string(item))
			break
		}
	}
	log.Info().Interface("common", common).Int("found", found).Str("cmpt1", compartment1).Str("cmpt2", compartment2).Send()
	return
}

func getItemPriority(item rune) (priority int) {
	// A = 65 a = 97 diff = 32
	// a - z = 1 - 26
	// A - Z = 27 - 52
	priority = int(item) - 96
	if priority < 0 {
		priority = int(item) - 38
	}

	return
}
