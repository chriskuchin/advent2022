package main

import (
	"bufio"
	"io"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Info().Msg("Day 6")
	file, err := os.Open("day6.txt")
	if err != nil {
		log.Info().Err(err).Send()
	}

	defer file.Close()
	r := bufio.NewReader(file)

	var sequence string
	var charCount int
	for {
		charCount++
		if c, _, err := r.ReadRune(); err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal().Err(err).Send()
			}
		} else {
			if len(sequence) < 14 {
				sequence += string(c)
			} else {
				sequence += string(c)
				sequence = sequence[1:]
				log.Info().Str("sequence", sequence).Int("chars", charCount).Send()
				if hasOnlyUniqueCharacters(sequence) {
					log.Info().Int("count", charCount).Str("seq", sequence).Send()
					break
				}
			}
		}
	}
}

func hasOnlyUniqueCharacters(seq string) bool {
	charMap := map[string]bool{}
	for _, ch := range seq {
		charMap[string(ch)] = true
	}
	return len(charMap) == len(seq)
}
