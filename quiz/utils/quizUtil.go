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

func InitializeFlags() (string, int) {
	fileName := FetchFileName()
	timeLimit := FetchTimeLimit()
	flag.Parse()

	return fileName, timeLimit
}

func FetchFileName() string {
	fileName := flag.String("csv", constants.DEFAULT_FILE, constants.CSV_HELP)
	return *fileName
}

func FetchTimeLimit() int {
	timeLimit := flag.Int("time", constants.DEFAULT_TIME, constants.TIME_HELP)
	return *timeLimit
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

func GetUserInput(channel chan string) {
	var answer string
	fmt.Scanf("%s\n", &answer)
	channel <- answer
}

func EvaluateAnswer(input string, actualAnswer string) bool {
	if input == actualAnswer {
		fmt.Printf("Correct Answer!\n")
		return true
	}
	fmt.Printf("Incorrect Answer!\n")
	return false
}
