package main

import (
	"GoQuiz/quiz/utils"
	"fmt"
)

func main() {
	quizFileName := utils.FetchFileName()
	quizFile := utils.LoadFile(quizFileName)
	problems := utils.LoadProblemsFromFile(quizFile)
	fmt.Print(problems)
}
