package main

import (
	"GoQuiz/quiz/utils"
	"fmt"
	"time"
)

func main() {
	quizFileName, quizTimeLimit := utils.InitializeFlags()
	quizFile := utils.LoadFile(quizFileName)
	problems := utils.LoadProblemsFromFile(quizFile)

	timer := time.NewTimer(time.Duration(quizTimeLimit) * time.Second)
	correctAnswers := 0

questionLoop:
	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s\n", (i + 1), problem.Question)

		answerChannel := make(chan string)
		go utils.GetUserInput(answerChannel)

		select {
		case <-timer.C:
			fmt.Printf("Time's up!\n")
			break questionLoop
		case answer := <-answerChannel:
			if utils.EvaluateAnswer(answer, problem.Answer) {
				correctAnswers++
			}
		}
	}
	fmt.Printf("You scored %d out of %d.", correctAnswers, len(problems))
}
