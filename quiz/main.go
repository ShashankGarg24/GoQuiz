package main

import (
	"GoQuiz/quiz/utils"
	"fmt"
)

func main() {
	quizFileName, quizTimeLimit := utils.InitializeFlags()
	fmt.Printf("time%d", quizTimeLimit)
	quizFile := utils.LoadFile(quizFileName)
	problems := utils.LoadProblemsFromFile(quizFile)

	correctAnswers := 0
	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s\n", (i + 1), problem.Question)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == problem.Answer {
			fmt.Printf("Correct Answer!\n")
			correctAnswers++
		} else {
			fmt.Printf("Incorrect Answer!\n")
		}
	}
	fmt.Printf("You scored %d out of %d.", correctAnswers, len(problems))
}
