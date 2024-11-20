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

// Add erwartet zwei Arrays der gleichen Länge.
// Addiert die Elemente der Arrays paarweise.
func Add(a, b []float64) {
	// SOLUTION
	for i := range a {
		a[i] += b[i]
	}
	// SOLUTION_END
}
