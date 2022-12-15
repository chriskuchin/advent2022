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
	log.Info().Msg("Day 7")
	file, err := os.Open("day7.txt")
	if err != nil {
		log.Info().Err(err).Send()
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	currentLoc := []string{}
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "$") {
			cmd := strings.TrimPrefix(line, "$ ")
			if strings.HasPrefix(cmd, "cd ") {
				dir := strings.TrimPrefix(cmd, "cd ")
				if dir != ".." && dir != "/" {
					currentLoc = append(currentLoc, dir)
				} else if dir == ".." {
					currentLoc = currentLoc[:len(currentLoc)-1]
				}
			}
			log.Info().Interface("loc", currentLoc).Str("folder", fmt.Sprintf("/%s", strings.Join(currentLoc, "/"))).Send()
		} else if strings.HasPrefix(line, "dir") {
			dir := strings.TrimPrefix(line, "dir ")
			log.Info().Str("dir", dir).Send()
		} else {
			split := strings.Split(line, " ")
			size, _ := strconv.Atoi(split[0])
			log.Warn().Str("file", split[1]).Int("size", size).Send()
		}
	}
}
