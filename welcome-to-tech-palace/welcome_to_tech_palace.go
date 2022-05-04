package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	output := ""

	for i := 0; i < numStarsPerLine; i++ {
		output += "*"
	}

	output += "\n"
	output += welcomeMsg
	output += "\n"

	for i := 0; i < numStarsPerLine; i++ {
		output += "*"
	}

	return output
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	output := oldMsg
	output = strings.ReplaceAll(output, "*", "")
	output = strings.TrimSpace(output)
	return output
}
