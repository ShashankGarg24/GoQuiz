package main

import (
	"GoQuiz/quiz/utils"
	"fmt"
	"time"
)

func main() {
	quizFileName, quizTimeLimit := utils.InitializeFlags()
	fmt.Printf("time%d", quizTimeLimit)
	quizFile := utils.LoadFile(quizFileName)
	problems := utils.LoadProblemsFromFile(quizFile)

	timer := time.NewTimer(time.Duration(quizTimeLimit) * time.Second)
	correctAnswers := 0
	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s\n", (i + 1), problem.Question)

		answerChannel := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerChannel <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("Time's up!\n")
			fmt.Printf("You scored %d out of %d.", correctAnswers, len(problems))
			return
		case answer := <- answerChannel:
			
			if answer == problem.Answer {
				fmt.Printf("Correct Answer!\n")
				correctAnswers++
			} else {
				fmt.Printf("Incorrect Answer!\n")
			}
		}
	}
	fmt.Printf("You scored %d out of %d.", correctAnswers, len(problems))
}
