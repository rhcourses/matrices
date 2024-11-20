package arraytools

// Mult erwartet ein Array und einen skalaren Faktor.
// Multipliziert jedes Element des Arrays mit dem Faktor.
func ScalarMult(a []float64, factor float64) {
	// SOLUTION
	for i := range a {
		a[i] *= factor
	}
	// SOLUTION_END
}

// HINT
// Iterieren Sie über alle Elemente des Arrays und multiplizieren Sie sie mit dem Faktor.

// Add erwartet zwei Arrays der gleichen Länge.
// Addiert die Elemente der Arrays paarweise.
func Add(a, b []float64) {
	// SOLUTION
	for i := range a {
		a[i] += b[i]
	}
	// SOLUTION_END
}

// HINT
// Iterieren Sie über alle Elemente der Arrays und addieren Sie sie paarweise.
// Schreiben Sie das Ergebnis zurück in das erste Array.
