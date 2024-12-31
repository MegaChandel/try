package main

import (
	handlers "a4_OnlineExaminationSystem/handler"
	"a4_OnlineExaminationSystem/model"
	"fmt"
)

func main() {

	questions := []model.Question{
		{
			QuestionText: "What is the capital of India?",
			Options:      []string{"1. Paris", "2. Rome", "3. Delhi", "4. Madrid"},
			Answer:       3,
		},
		{
			QuestionText: "What is 2 + 2?",
			Options:      []string{"1. 3", "2. 4", "3. 5", "4. 6"},
			Answer:       2,
		},
		{
			QuestionText: "Which planet is known as the Red Planet?",
			Options:      []string{"1. Earth", "2. Mars", "3. Jupiter", "4. Venus"},
			Answer:       2,
		},
		{
			QuestionText: "Who wrote 'Romeo and Juliet'?",
			Options:      []string{"1. Charles Dickens", "2. Jane Austen", "3. William Shakespeare", "4. Mark Twain"},
			Answer:       3,
		},
	}

	fmt.Println("Welcome to the Online Examination System!")
	score := handlers.TakeQuiz(questions)

	handlers.GradePerformance(score, len(questions))
}
