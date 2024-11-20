package matrixtools

// GetRow liefert die i-te Zeile der Matrix m.
func GetRow(m [][]float64, i int) []float64 {
	// SOLUTION
	return m[i]
	// SOLUTION_END
}

// GetCol liefert die j-te Spalte der Matrix m.
func GetCol(m [][]float64, j int) []float64 {
	// SOLUTION
	col := make([]float64, len(m))
	for i := range m {
		col[i] = m[i][j]
	}
	return col
	// SOLUTION_END
}

// AddRows erwartet eine Matrix und zwei Zeilennummern.
// Addiert die beiden Zeilen paarweise und speichert das Ergebnis in der ersten Zeile.
func AddRows(m [][]float64, i, j int) {
	// SOLUTION
	for k := range m[i] {
		m[i][k] += m[j][k]
	}
	// SOLUTION_END
}

// ScalarMultRow erwartet eine Matrix, eine Zeilennummer und einen skalaren Faktor.
// Multipliziert die Zeile mit dem Faktor und speichert das Ergebnis in der Zeile.
func ScalarMultRow(m [][]float64, i int, factor float64) {
	// SOLUTION
	for j := range m[i] {
		m[i][j] *= factor
	}
	// SOLUTION_END
}

// Transpose erwartet eine Matrix und liefert ihre Transponierte.
// D.h. alle Zeilen der ersten Matrix werden zu Spalten der Transponierten und umgekehrt.
func Transposed(m [][]float64) [][]float64 {
	// SOLUTION
	transposed := make([][]float64, len(m[0]))
	for i := range transposed {
		transposed[i] = make([]float64, len(m))
		for j := range transposed[i] {
			transposed[i][j] = m[j][i]
		}
	}
	return transposed
	// SOLUTION_END
}
