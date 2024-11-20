package matrix

import "fmt"

func ExampleMatrix_Row() {
	m := Matrix{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(m.Row(1))

	// Output: [4 5 6]
}

func ExampleGetCol() {
	m := Matrix{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(m.Col(1))

	// Output: [2 5]
}
