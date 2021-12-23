package main

import "fmt"

func main() {
	fmt.Println("Welcome to a mini quiz game!")

	var name string
	fmt.Printf("What's your name? ")
	fmt.Scan(&name)

	fmt.Printf("Hello, %s! I want to welcome you to the game! \n", name)

	var age uint
	fmt.Printf("How old are you? ")
	fmt.Scan(&age)

	if age > 10 {
		fmt.Println("You are good to play the game!")
	} else {
		fmt.Println("I'm sorry. You cannot play the game at your age.")
		return
	}

	score := 0
	num_questions := 3

	fmt.Printf("What element does 'O' represent on the periodic table? ")
	var answer1 string
	fmt.Scan(&answer1)
	if answer1 == "Oxygen" || answer1 == "oxygen" {
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Println("Incorrect! The correct answer was Oxygen or oxygen!")
	}

	fmt.Printf("Which city is the golden gate bridge in? hint: SF ")
	var answer2 string
	var answer3 string
	fmt.Scan(&answer2, &answer3)

	if answer2+" "+answer3 == "San Francisco" || answer2+" "+answer3 == "San francisco" || answer2+" "+answer3 == "san Francisco" || answer2+" "+answer3 == "san francisco" {
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Println("Incorrect! The correct answer was San Francisco! (San francisco, san Francisco, san francisco were all acceptable answers)")
	}

	var answer4 string
	fmt.Printf("What is Sweden's capital city? ")
	fmt.Scan(&answer4)
	if answer4 == "Stockholm" || answer4 == "stockholm" {
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Println("Incorrect! The correct answer was Stockholm or stockholm!")
	}

	percentage := (float64(score) / float64(num_questions)) * 100
	fmt.Printf("%s, you got %d out of %d questions right! \n", name, score, num_questions)
	fmt.Printf("You scored %f%%! \n", percentage)
}
