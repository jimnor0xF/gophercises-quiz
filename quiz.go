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

	//timeFlagPtr := flag.Int("timer", 30, "timer for questions")

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

	ticker := time.NewTicker(1 * time.Second)
	var correctAnswers int
	go runQuiz(correctAnswers, problems)
	go tickTime(ticker)
	time.Sleep(30 * time.Second)
	ticker.Stop()

	fmt.Printf("Score: %d/%d\n", correctAnswers, len(problems))

}

func runQuiz(correctAnswers int, problems []question) {
	for _, problem := range problems {

		fmt.Printf("What is %s?\n", problem.q)

		var input string
		_, err := fmt.Scanf("%s\n", &input)

		if err != nil {
			log.Fatal(err)
		}

		if input == problem.a {
			correctAnswers++
		}

	}

}

func tickTime(ticker *time.Ticker) {
	i := 1
	for t := range ticker.C {
		i = i + 1
		fmt.Println("Tick at", t)
	}

}

type question struct {
	q string
	a string
}
