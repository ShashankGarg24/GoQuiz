package utils

import (
	"GoQuiz/quiz/constants"
	"GoQuiz/quiz/models"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func FetchFileName() string {
	fileName := flag.String("csv", constants.DEFAULT_FILE, constants.CSV_HELP)
	flag.Parse()
	return *fileName
}

func LoadFile(name string) *os.File {
	file, err := os.Open(constants.QUIZ_PATH + name)
	if err != nil {
		gracefulExit(fmt.Sprintf("Error in opening the CSV file: %s", err))
	}
	return file
}

func LoadProblemsFromFile(file *os.File) []models.Problem {
	fileReader := csv.NewReader(file)
	lines, err := fileReader.ReadAll()
	if err != nil {
		gracefulExit(fmt.Sprintf("Error in reading the file: %s", err))
	}
	problems := parseLines(lines)
	return problems
}

func parseLines(lines [][]string) []models.Problem {
	problems := make([]models.Problem, len(lines))
	for i, line := range lines {
		problems[i] = models.Problem{
			Question: strings.TrimSpace(line[0]),
			Answer:   strings.TrimSpace(line[1]),
		}
	}
	return problems
}
