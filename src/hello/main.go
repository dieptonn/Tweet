package main

import "fmt"

// var dayscomplete = 11
var TwitterName string
var DaysCompleted uint //so nguyen duong
var remainingDays uint = 90

func main() {

	var challenge = "90DaysOfDevOps"
	const daystotal = 90
	fmt.Printf("Welcome to %s, what your name? \n", challenge)
	fmt.Scan(&TwitterName)
	fmt.Printf("How many days have you completed?: ")
	fmt.Scan(&DaysCompleted)

	remainingDays = remainingDays - DaysCompleted

	fmt.Printf("This is a %v challenge and you (%v) have completed %v days \n", daystotal, TwitterName, DaysCompleted)
	fmt.Printf("You have %v days remaining for the %v challenge\n", remainingDays, daystotal)
	fmt.Println("Great work!")
}
