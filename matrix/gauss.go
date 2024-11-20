package matrix

// Normalize erwartet eine Spaltennummer.
// Falls das Diagonalelement [col][col] nicht 0 ist, wird die Zeile durch das Diagonalelement normiert.
// D.h. die gesamte Zeile col wird durch das Diagonalelement geteilt.
func (m Matrix) Normalize(col int) {
	// SOLUTION
	if almostEqual(m[col][col], 0) {
		return
	}
	factor := 1 / m[col][col]
	m.ScalarMultRow(col, factor)
	// SOLUTION_END
}

// EliminateBelow erwartet eine Zeilennummer `row`.
// Multipliziert alle Zeilen unter der Zeile row mit -1/Matrix[row][row] und addiert sie zur Zeile row.
// Dadurch wird jeweils das Element unter dem Diagonalelement 0.
// Voraussetzung: Die Zeile row ist bereits normiert, d.h. das Diagonalelement ist 1.
func (m Matrix) EliminateBelow(row int) {
	// SOLUTION
	for i := row + 1; i < len(m); i++ {
		factor := -1 / m[i][row]
		m.ScalarMultRow(i, factor)
		m.AddRows(i, row)
	}
	// SOLUTION_END
}

// EliminateAbove erwartet eine Zeilennummer `row`.
// Multipliziert alle Zeilen über der Zeile row mit -1/Matrix[row][row] und addiert sie zur Zeile row.
// Dadurch wird jeweils das Element über dem Diagonalelement 0.
// Voraussetzung: Die Zeile row ist bereits normiert, d.h. das Diagonalelement ist 1.
func (m Matrix) EliminateAbove(row int) {
	// SOLUTION
	for i := 0; i < row; i++ {
		factor := -1 / m[i][row]
		m.ScalarMultRow(i, factor)
		m.AddRows(i, row)
	}
	// SOLUTION_END
}

// UpperTriangular führt die Gauß-Elimination für alle Zeilen der Matrix durch.
// So entsteht im linken Bereich eine obere Dreiecksmatrix, bei der die Diagonalelemente 1 sind.
func (m Matrix) UpperTriangular() {
	// SOLUTION
	for i := range m {
		m.Normalize(i)
		m.EliminateBelow(i)
	}
	fixZeros(m)
	// SOLUTION_END
}

// LowerTriangular führt die Gauß-Elimination für alle Zeilen der Matrix durch.
// So entsteht im linken Bereich eine untere Dreiecksmatrix, bei der die Diagonalelemente 1 sind.
func (m Matrix) LowerTriangular() {
	// SOLUTION
	l := len(m)
	for i := range m {
		m.Normalize(l - i - 1)
		m.EliminateAbove(l - i - 1)
	}
	fixZeros(m)
	// SOLUTION_END
}

// Gauss transformiert die Matrix im linken Bereich in die Einheitsmatrix.
func (m Matrix) Gauss() {
	// SOLUTION
	m.UpperTriangular()
	m.LowerTriangular()
	// SOLUTION_END
}

// almostEqual ist eine Hilfsfunktion, die zwei float64-Werte auf Gleichheit prüft.
// Da float64-Werte nicht exakt vergleichbar sind, wird ein Toleranzwert von 1e-10 verwendet.
func almostEqual(a, b float64) bool {
	return a-b < 1e-10 && b-a < 1e-10
}

// fixZeros erwartet eine Matrix und setzt alle Elemente, die kleiner als 1e-10 sind, auf 0.
func fixZeros(m Matrix) {
	for i := range m {
		for j := range m[i] {
			if almostEqual(m[i][j], 0) {
				m[i][j] = 0
			}
		}
	}
}
