package gerade_zahlen_filtern

// Gerade Zahlen filtern
//
// Ergänze die TODO-Stellen.
// Die Funktion wird automatisch durch die zugehörige Testdatei geprüft.

func FilterEven(numbers []int) []int {
	evenNumbers := []int{}
	for _, value := range numbers {
		if value%2 == 0 {
			evenNumbers = append(evenNumbers, value)
		}
	}

	// TODO: Füge nur gerade Zahlen hinzu.

	return evenNumbers
}
