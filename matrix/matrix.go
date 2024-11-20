package matrix

import "matrices/matrixtools"

type Row = []float64
type Col = []float64
type Matrix []Row

// Row liefert die i-te Zeile der Matrix m.
func (m Matrix) Row(i int) Row {
	return matrixtools.GetRow(m, i)
}

// Col liefert die j-te Spalte der Matrix m.
func (m Matrix) Col(j int) Col {
	return matrixtools.GetCol(m, j)
}
