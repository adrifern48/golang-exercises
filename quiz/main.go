package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	// Evaluate CSV file path
	var filename string
	flag.StringVar(&filename, "file", "./problems.csv", "file to read problems from")
	flag.Parse()

	// Used to read user input from stdin
	inputReader := bufio.NewReader(os.Stdin)

	// Read CSV
	csvFile, _ := os.Open(filename)
	reader := csv.NewReader(bufio.NewReader(csvFile))

	// Used to keep track of score
	questionsCount := 0
	correctAnsCount := 0

	for {
		line, error := reader.Read()

		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		// Normalize both question and answer to lowercase
		line[0] = strings.ToLower(line[0])
		line[1] = strings.ToLower(line[1])

		questionsCount++

		// Prompt for user input
		question := fmt.Sprintf("%s = ", line[0])
		fmt.Print(question)

		// Read user input and match against the correct answer
		answerInput, _ := inputReader.ReadString('\n')
		answerInputTrimmed := strings.TrimSpace(answerInput)
		answer, _ := strconv.Atoi(answerInputTrimmed)
		correctAnswer, _ := strconv.Atoi(strings.TrimSpace(line[1]))
		if answer == correctAnswer {
			correctAnsCount++
		}
	}

	totalScore := fmt.Sprintf("Quiz Finished! You scored %d out of %d", correctAnsCount, questionsCount)
	fmt.Println(totalScore)
}
