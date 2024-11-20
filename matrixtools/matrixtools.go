package matrixtools

// GetRow liefert die i-te Zeile der Matrix m.
func GetRow(m [][]float64, i int) []float64 {
	return m[i]
}

// GetCol liefert die j-te Spalte der Matrix m.
func GetCol(m [][]float64, j int) []float64 {
	col := make([]float64, len(m))
	for i := range m {
		col[i] = m[i][j]
	}
	return col
}

// AddRows erwartet eine Matrix und zwei Zeilennummern.
// Addiert die beiden Zeilen paarweise und speichert das Ergebnis in der ersten Zeile.
func AddRows(m [][]float64, i, j int) {
	for k := range m[i] {
		m[i][k] += m[j][k]
	}
}

// ScalarMultRow erwartet eine Matrix, eine Zeilennummer und einen skalaren Faktor.
// Multipliziert die Zeile mit dem Faktor und speichert das Ergebnis in der Zeile.
func ScalarMultRow(m [][]float64, i int, factor float64) {
	for j := range m[i] {
		m[i][j] *= factor
	}
}
