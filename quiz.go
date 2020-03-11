package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	fileFlagPtr := flag.String("csv_file", "problems.csv", "file path to problems")

	timeFlagPtr := flag.Int("timer", 30, "timer for questions")

	flag.Parse()

	file, err := os.Open(*fileFlagPtr)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	r := csv.NewReader(file)

	records, err := r.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	problems := make([]question, len(records))

	for i, record := range records {
		problems[i].q = record[0]
		problems[i].a = record[1]
	}

	var correctAnswers int
	timer := time.NewTimer(time.Duration(*timeFlagPtr) * time.Second)

Loop:
	for _, problem := range problems {
		answerChan := make(chan string)
		fmt.Printf("What is %s?\n", problem.q)

		go func() {
			var input string
			_, err := fmt.Scanf("%s\n", &input)

			if err != nil {
				log.Fatal(err)
			}

			answerChan <- input

		}()

		select {
		case <-timer.C:
			break Loop
		case answer := <-answerChan:
			if answer == problem.a {
				correctAnswers++

			}
		}

	}
	fmt.Printf("Score: %d/%d\n", correctAnswers, len(problems))

}

type question struct {
	q string
	a string
}
