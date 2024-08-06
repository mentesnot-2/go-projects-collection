package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func calculateAverage(grades []float64) float64 {
	total := 0.0
	for _, grade := range grades {
		total += grade
	}
	return total / float64(len(grades))
}

func main() {
	var name string
	var numSubjects int

	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)

	fmt.Print("Enter the number of subjects: ")
	_, err := fmt.Scanf("%d\n", &numSubjects) 
	if err != nil {
		log.Fatal(err)
	}

	subjects := make([]string, numSubjects)
	grades := make([]float64, numSubjects)

	for i := 0; i < numSubjects; i++ {
		fmt.Printf("Enter the name of subject %d: ", i+1)
		fmt.Scanln(&subjects[i])
		var grade float64
		fmt.Printf("Enter the grade for %s: ", subjects[i])
		_, err := fmt.Scanln(&grade)
		if err != nil || grade < 0 || grade > 100 {
			log.Fatalf("Invalid grade for %s: %v", subjects[i], err)
		}
		grades[i] = grade
	}

	average := calculateAverage(grades)
	fmt.Printf("\nStudent Name: %s\n", strings.ToUpper(name))
	fmt.Println("Subject Grades: ")
	for i := 0; i < numSubjects; i++ {
		fmt.Printf("%s: %.2f\n", subjects[i], grades[i])
	}
	fmt.Printf("Average Grade: %.2f\n", average)
	var word string
	fmt.Print("Enter a word: ")
	fmt.Scanln(&word)

	if IsPalindrome(word) {
		fmt.Printf("%s is a palindrome\n", word)
	} else {
		fmt.Printf("%s is not a palindrome\n", word)
	}
	var sentence string
	fmt.Print("Enter a sentence: ")
	read:=bufio.NewReader(os.Stdin)
	sentence, _ = read.ReadString('\n')
	sentence = strings.TrimSpace(sentence)

	wordFrequency := WordFrequency(sentence)
	fmt.Printf("Word Frequency: %v\n", wordFrequency)
}
