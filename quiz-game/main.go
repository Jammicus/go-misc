package main

/**
TODO
Timers
**/

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	iteration := 0
	score := 0
	filename := flag.String("filename", "problems.csv", "Specify CSV File to use")
	timervalue := flag.Duration("duration", 30, "Specify the time in seconds per question")
	// Convert time to seconds.
	fmt.Println(timervalue)
	flag.Parse()

	contents := parseFile(*filename)
	contents.Comma = ' '

	for {
		record, err := contents.Read()
		if err == io.EOF {
			break
		}
		entry := strings.Split(record[0], ",")
		question, answer := entry[0], entry[1]
		fmt.Println("# Problem", iteration, question)
		iteration++
		score = score + checkAnswer(answer)

	}

	fmt.Println("Final Score:", score, "out of", iteration)
}

func parseFile(path string) *csv.Reader {
	file, err := os.Open(path)
	check(err)
	return csv.NewReader(file)
}

func checkAnswer(expected string) int {
	value := bufio.NewReader(os.Stdin)
	text, err := value.ReadString('\n')
	check(err)

	expected = strings.Trim(expected, "\n")
	text = strings.Trim(text, "\n")

	if expected == text {
		return 1
	}
	return 0
}
