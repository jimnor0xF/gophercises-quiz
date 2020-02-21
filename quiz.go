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

	lines, err := r.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(lines)

}

type question struct {
	q string
	a string
}
