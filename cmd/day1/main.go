package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/rs/zerolog/log"
)

func main() {
	fmt.Println("Hello Worldgo")

	file, err := os.Open("day1.txt")
	if err != nil {
		log.Info().Err(err).Send()
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var elfTotal int
	var currentMax []int = []int{}
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			sort.Ints(currentMax)
			if len(currentMax) < 3 {
				currentMax = append(currentMax, elfTotal)
			} else {
				for idx, val := range currentMax {
					if elfTotal > val {
						currentMax[idx] = elfTotal
						break
					}
				}
			}

			// log.Info().Interface("max", currentMax).Send()

			elfTotal = 0
		} else {
			currentCal, _ := strconv.Atoi(line)
			elfTotal += currentCal
		}
	}

	var total int
	for _, calories := range currentMax {
		total += calories
	}

	log.Info().Int("calories", total).Send()
}

func partOne() {
	file, err := os.Open("day1.txt")
	if err != nil {
		log.Info().Err(err).Send()
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var elfTotal int
	var elfCount int
	var currentMax int
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			elfCount++
			if currentMax < elfTotal {
				currentMax = elfTotal
			}
			log.Info().Int("total", currentMax).Send()
			elfTotal = 0
		} else {
			currentCal, _ := strconv.Atoi(line)
			elfTotal += currentCal
		}
		// log.Info().Str("line", line).Send()
	}
}
