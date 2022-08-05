package main

import (
	"fmt"
	"strings"
	"the-giuseppe-app/helper"
)

const hotelRooms = 5

var hotelName = "The Giuseppe"
var remainingRooms uint = 5
var bookings = []string{}

func main() {

	greetUsers()

	for {

		firstName, lastName, email, userRooms := getUserInput()
		isValidName, isValidEmail, isValidRoomQuantity := helper.ValidateUserInput(firstName, lastName, email, userRooms, remainingRooms)

		if isValidEmail && isValidName && isValidRoomQuantity {

			bookRooms(userRooms, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("The following customers have booked: %v\n", firstNames)

			if remainingRooms == 0 {
				fmt.Println("Sorry, all rooms are booked")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Invalid name")
			}
			if !isValidEmail {
				fmt.Println("Invalid email")
			}
			if !isValidRoomQuantity {
				fmt.Printf("Invalid room quantity. You can book up to %v rooms\n", remainingRooms)
			}
			fmt.Printf("Your input is invalid. Please try again\n")
			continue
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v\n", hotelName)
	fmt.Printf(hotelName+" has %v luxury rooms\n", hotelRooms)
	fmt.Printf("There are %v rooms remaining\n", remainingRooms)
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userRooms uint

	fmt.Print("What is your first name? ")
	fmt.Scan(&firstName)

	fmt.Print("What is your last name? ")
	fmt.Scan(&lastName)

	fmt.Print("What is your email? ")
	fmt.Scan(&email)

	fmt.Print("How many Rooms would you like? ")
	fmt.Scan(&userRooms)

	return firstName, lastName, email, userRooms
}

func bookRooms(userRooms uint, firstName string, lastName string, email string) {
	remainingRooms = remainingRooms - userRooms
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("%v %v, your booking has been confirmed.\n", firstName, lastName)
	fmt.Printf("Your reservation details for %v rooms has been sent to %v.\n", userRooms, email)
	fmt.Printf("%v rooms remaining\n", remainingRooms)
}
