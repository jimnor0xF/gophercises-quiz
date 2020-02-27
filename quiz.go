package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	flagPtr := flag.String("csv_file", "problems.csv", "file path to problems")
	flag.Parse()

	file, err := os.Open(*flagPtr)

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

	fmt.Printf("Score: %d/%d\n", correctAnswers, len(problems))

}

type question struct {
	q string
	a string
}
