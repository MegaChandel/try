package handlers

import (
	"a4_OnlineExaminationSystem/model"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func TakeQuiz(questions []model.Question) int {
	reader := bufio.NewReader(os.Stdin)
	score := 0

	fmt.Println("Welcome to the quiz! Type 'exit' to quit anytime.")
	fmt.Println()

	for i, question := range questions {
		fmt.Printf("Question %d: %s\n", i+1, question.QuestionText)

		for j, option := range question.Options {
			fmt.Printf("%d. %s\n", j+1, option)
		}

		fmt.Print("Enter your answer (1-4): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			fmt.Println("You left the quiz.")
			break
		}

		answer, err := strconv.Atoi(input)
		if err != nil || answer < 1 || answer > 4 {
			fmt.Println("Invalid input. Try again.")
			fmt.Println()
			continue
		}

		if answer == question.Answer {
			fmt.Println("Correct!")
			score++
		} else {
			fmt.Println("Wrong answer.")
		}
		fmt.Println()
	}

	return score
}

func GradePerformance(score int, totalQuestions int) {
	fmt.Printf("\nYour score: %d out of %d.\n", score, totalQuestions)

	percentage := float64(score) / float64(totalQuestions) * 100

	if percentage >= 90 {
		fmt.Println("Excellent")
	} else if percentage >= 70 {
		fmt.Println("Great job")
	} else if percentage >= 50 {
		fmt.Println("Not bad")
	} else {
		fmt.Println("Needs improvement")
	}
}
