package matrix

import (
	"errors"
	"strconv"
	"strings"
)

var ()

//Matrice ...
type Matrice interface {
	Rows() [][]int
	Cols() [][]int
	Set(row int, col int, val int) bool
}

type matrice struct {
	matrixRowVar [][]int
	rowsNum      int
	colsNum      int
}

// New ...
func New(inputNumbers string) (Matrice, error) {
	var err error

	rowsCol := 0
	rows := strings.Split(inputNumbers, "\n")
	matrixRows := [][]int{}

	for _, rowValues := range rows {
		rowVals := strings.Split(strings.TrimSpace(rowValues), " ")
		aRow := make([]int, len(rowVals))

		if rowsCol != 0 && rowsCol != len(aRow) {
			return nil, errors.New("Different Dimension")
		} else {
			rowsCol = len(aRow)
		}

		for index, rowVal := range rowVals {
			aRow[index], err = strconv.Atoi(rowVal)
			if err != nil {
				return nil, errors.New(err.Error())
			}
		}

		matrixRows = append(matrixRows, aRow)
	}

	mtx := &matrice{matrixRows, len(matrixRows), len(matrixRows[0])}

	return mtx, nil
}

// Rows will duplicate the processed matrix and return a copy
func (m *matrice) Rows() [][]int {
	duplicate := make([][]int, len(m.matrixRowVar))
	for i := range m.matrixRowVar {
		duplicate[i] = make([]int, len(m.matrixRowVar[i]))
		copy(duplicate[i], m.matrixRowVar[i])
	}

	return duplicate
}

// Cols will transpose the processed matrix
func (m *matrice) Cols() [][]int {
	matrixcols := [][]int{}

	for i := 0; i < m.colsNum; i++ {
		newRow := make([]int, m.rowsNum)
		for j := 0; j < m.rowsNum; j++ {
			newRow[j] = m.matrixRowVar[j][i]
		}

		matrixcols = append(matrixcols, newRow)
	}

	return matrixcols
}

// Set will check if row and col inside boundary
func (m *matrice) Set(row int, col int, val int) bool {

	if row >= 0 && row < m.rowsNum && col >= 0 && col < m.colsNum {
		m.matrixRowVar[row][col] = val

		return true
	}

	return false
}
