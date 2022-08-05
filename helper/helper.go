package helper

import (
	"strings"
)

func ValidateUserInput(firstName string, lastName string, email string, userRooms uint, remainingRooms uint) (bool, bool, bool) {
	isValidName := len(firstName) > 2 && len(lastName) > 2
	isValidEmail := strings.Contains(email, "@")
	isValidRoomQuantity := userRooms <= remainingRooms
	return isValidName, isValidEmail, isValidRoomQuantity
}
