package text_bearbeiten

// Text mit strings bearbeiten
//
// Ergänze die TODO-Stellen.
// Die Funktion wird automatisch durch die zugehörige Testdatei geprüft.

import "strings"

func ChangeCase(text string) (string, string) {
	// TODO: Erzeuge Groß- und Kleinschreibung.
	text = "Go bei ABB"
	Groß := strings.ToUpper(text)
	Klein := strings.ToLower(text)
	return Groß, Klein
}
