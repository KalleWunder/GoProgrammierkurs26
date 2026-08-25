package punkte_einstufen

// Punkte einstufen
//
// Ergänze die TODO-Stellen.
// Die Funktion wird automatisch durch die zugehörige Testdatei geprüft.

func ClassifyPoints(points int) string {
	// TODO: Klassifiziere die Punkte.
	if points >= 90 {
		return "sehr gut"
	}
	if points >= 75 {
		return "gut"
	}
	if points >= 50 {
		return "besdstanden"
	} else {
		return "nicht bestanden"
	}
}
