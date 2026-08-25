package hello_world

import "fmt"

// Hello, World!
//
// Ergänze die TODO-Stellen.
// Die Funktion wird automatisch durch die zugehörige Testdatei geprüft.

func HelloWorld() string {
	// TODO: Gib "Hello, World!" zurück.
	greets := "Hello, World!"
	fmt.Sprintf("%s", greets)
	return greets
}
