package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

// Define the Matrix type here.
type Matrix [][]int

func New(s string) (*Matrix, error) {
	rowStrings := strings.Split(s, "\n")
	rowCount := len(rowStrings)
	matrix := make(Matrix, rowCount)

	for i, rowString := range rowStrings {
		colStrings := strings.Split(strings.TrimSpace(rowString), " ")
		matrix[i] = make([]int, len(colStrings))
		for j, colString := range colStrings {
			value, err := strconv.Atoi(colString)
			if err != nil {
				return nil, fmt.Errorf("atoi colString: %w, %s", err, colString)
			}
			if ok := matrix.Set(i, j, value); !ok {
				return nil, fmt.Errorf("failed to set value: %v", value)
			}
		}
	}

	if err := matrix.validate(); err != nil {
		return nil, fmt.Errorf("invalid matrix: %w", err)
	}

	return &matrix, nil
}

func (m *Matrix) validate() error {
	// check rows
	rowL := len((*m)[0])

	for _, row := range *m {
		if len(row) != rowL {
			return fmt.Errorf("uneven matrix row lengths")
		}
	}
	return nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m *Matrix) Cols() [][]int {
	rowL := len((*m)[0])

	res := make([][]int, rowL)

	for i := range res {
		res[i] = make([]int, len(*m))
	}

	for i, row := range *m {
		for j, _ := range row {
			res[j][i] = (*m)[i][j]
		}
	}

	return res
}

func (m *Matrix) Rows() [][]int {
	res := make([][]int, len(*m))

	for i, row := range *m {
		res[i] = make([]int, len(row))
		for j, col := range row {
			res[i][j] = col
		}
	}

	return res
}

func (m *Matrix) Set(row, col, val int) bool {
	if row >= len(*m) || col >= len((*m)[0]) || row < 0 || col < 0 {
		return false // index out of bounds
	}
	(*m)[row][col] = val
	return true
}
