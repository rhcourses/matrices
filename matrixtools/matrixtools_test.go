package matrixtools

import "fmt"

func ExampleGetRow() {
	m := [][]float64{
		{1, 2, 3},
		{4, 5, 6},
	}
	row := GetRow(m, 1)
	fmt.Println(row)

	// Output: [4 5 6]
}

func ExampleGetCol() {
	m := [][]float64{
		{1, 2, 3},
		{4, 5, 6},
	}
	col := GetCol(m, 1)
	fmt.Println(col)

	// Output: [2 5]
}
