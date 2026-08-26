package countdown

// Countdown
//
// Ergänze die TODO-Stellen.
// Die Funktion wird automatisch durch die zugehörige Testdatei geprüft.

func Countdown(start int) []int {
	var result []int
	for i := 5; i > 0; i-- {
		result = append(result, i)
	}

	// TODO: Zähle von start bis 1 rückwärts.

	return result
}
