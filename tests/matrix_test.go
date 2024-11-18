package tests

import (
	"bytes"
	"testing"

	"github.com/as283-ua/crypto/aes"
)

func assertEqualsBytes(t *testing.T, expected, result []byte) {
	if !bytes.Equal(expected, result) {
		t.Errorf("expected %v but got %v", expected, result)
	}
	// fmt.Println(expected, result)
}

func TestGetMatrix1And2D(t *testing.T) {
	a := aes.GetMatrix2D([][]byte{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})

	b := aes.GetMatrix1D([]byte{
		1, 4, 7,
		2, 5, 8,
		3, 6, 9,
	}, 3)

	assertEqualsBytes(t, a.Array, b.Array)
}

func TestMatrixGetCol(t *testing.T) {
	a := aes.GetMatrix2D([][]byte{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})

	expected := []byte{2, 5, 8}
	result := a.GetCol(1)

	assertEqualsBytes(t, expected, result)
}

func TestMatrixGetRow(t *testing.T) {
	a := aes.GetMatrix2D([][]byte{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})

	expected := []byte{4, 5, 6}
	result := a.GetRow(1)

	assertEqualsBytes(t, expected, result)
}

func TestMatrixMult(t *testing.T) {
	// a := aes.GetMatrix2D([][]byte{
	// 	{1, 2, 3, 4},
	// 	{5, 6, 7, 8},
	// 	{9, 10, 11, 12},
	// 	{13, 14, 15, 0},
	// })

	// var b = aes.GetMatrix2D([][]byte{
	// 	{2, 3, 1, 1},
	// 	{1, 2, 3, 1},
	// 	{1, 1, 2, 3},
	// 	{3, 1, 1, 2},
	// })

}

func TestIdentity(t *testing.T) {
	var id = aes.GetIdentity(3)

	a := aes.GetMatrix2D([][]byte{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	})

	assertEqualsBytes(t, id.Array, a.Array)
}

func TestInvMatrixMult(t *testing.T) {
	var mixColsTable = aes.GetMatrix2D([][]byte{
		{2, 3, 1, 1},
		{1, 2, 3, 1},
		{1, 1, 2, 3},
		{3, 1, 1, 2},
	})

	var invMixColsTable = aes.GetMatrix2D([][]byte{
		{14, 11, 13, 9},
		{9, 14, 11, 13},
		{13, 9, 14, 11},
		{11, 13, 9, 14},
	})

	res := aes.MatrixMult(mixColsTable, invMixColsTable)
	expect := aes.GetIdentity(4)
	assertEqualsBytes(t, expect.Array, res.Array)
}

func TestInvMatrixMultAes(t *testing.T) {
	var mixColsTable = aes.GetMatrix2D([][]byte{
		{2, 3, 1, 1},
		{1, 2, 3, 1},
		{1, 1, 2, 3},
		{3, 1, 1, 2},
	})

	var invMixColsTable = aes.GetMatrix2D([][]byte{
		{14, 11, 13, 9},
		{9, 14, 11, 13},
		{13, 9, 14, 11},
		{11, 13, 9, 14},
	})

	a := getState()

	b := aes.MatrixMult(aes.BMatrix{Array: a, Dimension: 4}, mixColsTable)

	c := aes.MatrixMult(b, invMixColsTable)

	if !bytes.Equal(a, c.Array) {
		t.Errorf("expected state to be %v but got %v", a, c)
	}
}
