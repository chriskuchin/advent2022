package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	re         *regexp.Regexp = regexp.MustCompile(`move (?P<count>\d+) from (?P<src>\d+) to (?P<dest>\d+)`)
	containers [][]string     = [][]string{
		{}, // 1
		{}, // 2
		{}, // 3
		{}, // 4
		{}, // 5
		{}, // 6
		{}, // 7
		{}, // 8
		{}, // 9
	}
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
	var reversedContainerStacks bool
	tempContainer := [][]string{{}, {}, {}, {}, {}, {}, {}, {}, {}}
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
						tempContainer[column-1] = append(tempContainer[column-1], string(entry))
						continueCount = 2
					}
				} else {
					continueCount--
				}
			}
		} else {
			if !reversedContainerStacks {
				reversedContainerStacks = true
				reverseContainerStacks(tempContainer)
			}
			if line == "" {
				continue
			}
			count, src, dest := getCommand(line)
			processMultipleCommand(count, src, dest)
		}
	}
	log.Info().Interface("containers", containers).Str("solution", getResults()).Send()
}

func reverseContainerStacks(tempContainer [][]string) {
	for stackIdx, stack := range tempContainer {
		for i := len(stack); i > 0; i-- {
			containers[stackIdx] = append(containers[stackIdx], stack[i-1])
		}
	}
}

func getCommand(line string) (count, src, dest int) {
	match := re.FindStringSubmatch(line)
	count, _ = strconv.Atoi(match[1])
	src, _ = strconv.Atoi(match[2])
	dest, _ = strconv.Atoi(match[3])

	return
}

func processSingleMoveCommand(count, sc, dest int) {

}

func processCommand(count, src, dest int) {
	for i := 0; i < count; i++ {
		var item string
		item, containers[src-1] = pop(containers[src-1])
		containers[dest-1] = append(containers[dest-1], item)
	}
}

func processMultipleCommand(count, src, dest int) {
	var item []string
	item, containers[src-1] = popMultiple(containers[src-1], count)
	containers[dest-1] = append(containers[dest-1], item...)
}

func pop(src []string) (val string, list []string) {
	val = src[len(src)-1]
	list = src[:len(src)-1]
	return
}

func popMultiple(src []string, count int) (val, list []string) {
	val = src[len(src)-count:]
	list = src[:len(src)-(count)]

	return
}

func getResults() string {
	var results string = ""

	for _, stack := range containers {
		results += stack[len(stack)-1]
	}

	return results
}
