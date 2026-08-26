package werte_mit_range

// Werte mit range durchlaufen
//
// Ergänze die TODO-Stellen.
// Die Funktion wird automatisch durch die zugehörige Testdatei geprüft.

func CopyValues(numbers []int) []int {
	var result []int
	for _, number := range numbers {
		result = append(result, number)
	}

	// TODO: Kopiere alle Werte mit range.

	return result
}
