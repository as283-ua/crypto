package aes

import "fmt"

/*
Expresses a matrix as an array, groups of x elements represent a column
*/
type BMatrix struct {
	Array     []byte
	Dimension int
}

func (m *BMatrix) GetRow(row int) []byte {
	res := make([]byte, m.Dimension)

	for i := 0; i < m.Dimension; i++ {
		res[i] = m.Array[m.Dimension*i+row]
	}

	return res
}

func (m *BMatrix) GetCol(col int) []byte {
	startIdx := m.Dimension * col
	return m.Array[startIdx : startIdx+m.Dimension]
}

func (m *BMatrix) Get(row, col int) byte {
	return m.GetCol(col)[row]
}

func (m *BMatrix) Set(row, col int, value byte) {
	m.Array[row+col*m.Dimension] = value
}

func GetMatrix2D(m [][]byte) BMatrix {
	colSize := len(m[0])
	res := BMatrix{Array: make([]byte, colSize*colSize), Dimension: colSize}
	for i, v := range m {
		for j, b := range v {
			res.Array[i+j*colSize] = b
		}
	}
	return res
}

func GetMatrix1D(m []byte, dim int) BMatrix {
	res := BMatrix{Array: m, Dimension: dim}
	return res
}

func GetIdentity(dim int) BMatrix {
	res := BMatrix{Array: make([]byte, dim*dim), Dimension: dim}

	for i := 0; i < dim; i++ {
		res.Set(i, i, 1)
	}

	return res
}

func GetDefaultMatrix(size int) BMatrix {
	res := BMatrix{Array: make([]byte, size*size), Dimension: size}
	return res
}

func RowMatrixMult(a []byte, b []byte, dimension int) (res byte) {
	if len(a) != dimension || len(b) != dimension {
		panic(fmt.Sprint("Matrix must be of size", dimension))
	}

	for i := 0; i < dimension; i++ {
		res ^= GaloisMult(a[i], b[i])
	}

	return
}

func MatrixMult(mat, other BMatrix) BMatrix {
	res := GetDefaultMatrix(4)

	for row := 0; row < 4; row++ {
		for col := 0; col < 4; col++ {
			rowMult := RowMatrixMult(mat.GetRow(row), other.GetCol(col), mat.Dimension)
			res.Array[col*res.Dimension+row] = rowMult
		}
	}
	return res
}
