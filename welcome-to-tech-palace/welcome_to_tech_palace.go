package techpalace

import(
    "fmt"
    "strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	starString := strings.Repeat("*", numStarsPerLine)
    return fmt.Sprintf("%s\n%s\n%s", starString, welcomeMsg, starString)
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	cleanUpStars := strings.Replace(oldMsg, "*","", -1)
    cleanUpNewLines :=strings.Replace(cleanUpStars, "\n","", -1)
    cleanUpWhiteSpace := strings.TrimSpace(cleanUpNewLines)
    return cleanUpWhiteSpace
}
