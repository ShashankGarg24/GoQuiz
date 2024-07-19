package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	quizFileName := fetchFileName()
	quizFile := loadFile(quizFileName)
	fmt.Print(quizFile)

}

func fetchFileName() string {
	fileName := flag.String("csv", "math-problems.csv", "Provide with a csv file of problems in the format of Question, Answer")
	flag.Parse()
	return *fileName
}

func loadFile(name string) *os.File {
	file, err := os.Open("../quizes/" + name)
	if err != nil {
		gracefulExit(fmt.Sprintf("Error in opening the CSV file: %s", err))
	}
	return file
}

func gracefulExit(msg string) {
	fmt.Printf(msg)
	os.Exit(1)
}
