package functions

import (
	"github.com/syuya2036/GoZero/matrix"
	"math"
)

type Square struct {
	FunctionBase
}

func NewSquare() Square {
	var s Square = Square{}
	return	s
}

func (s Square) Forward(mat *matrix.Matrix) *matrix.Matrix {
	result := make([]float64, mat.Rows*mat.Cols)
	for i := range result {
		result[i] = math.Pow(mat.Mat[i], 2)
	}

	return matrix.NewMatrix(result, mat.Rows, mat.Cols)
}

func (s Square) Backward(mat *matrix.Matrix) *matrix.Matrix {
	return  nil
}