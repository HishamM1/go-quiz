package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type problem struct {
	question string
	answer   string
}

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	problems, err := readProblems(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFileName))
	}

	correctAnswers := startQuiz(problems, *timeLimit)

	fmt.Printf("You scored %d out of %d\n", correctAnswers, len(problems))
}

func readProblems(csvFileName string) ([]problem, error) {
	file, err := os.Open(csvFileName)
	if err != nil {
		return nil, err
	}

	problemsCSV := csv.NewReader(file)
	lines, err := problemsCSV.ReadAll()
	if err != nil {
		return nil, err
	}

	problems := make([]problem, len(lines))
	for i, line := range lines {
		problems[i] = problem{
			question: strings.TrimSpace(line[0]),
			answer:   strings.TrimSpace(line[1]),
		}
	}

	return problems, nil
}

func startQuiz(problems []problem, timeLimit int) int {
	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)
	correctAnswers := 0
	scanner := bufio.NewScanner(os.Stdin)

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.question)
		answerCh := make(chan string)

		go func() {
			scanner.Scan()
			answerCh <- scanner.Text()
		}()

		select {
		case <-timer.C:
			return correctAnswers
		case answer := <-answerCh:
			if answer == "exit" {
				return correctAnswers
			}

			if answer == p.answer {
				fmt.Println("Correct!")
				correctAnswers++
			} else {
				fmt.Println("Incorrect!")
			}
		}
	}

	return correctAnswers
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
