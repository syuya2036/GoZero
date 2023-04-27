package functions

import (
	"github.com/syuya2036/GoZero/matrix"
	"math"
)

func NewSquare() *FunctionBase {
	var s *FunctionBase = &FunctionBase{Name: "Square", Forward: forward, Backward: backward}
	return	s
}

func forward(mat matrix.Matrix) matrix.Matrix {
	result := make([]float64, mat.Rows*mat.Cols)
	for i := range result  {
		result[i] = math.Pow(mat.Mat[i], 2)
	}

	return matrix.NewMatrix(result, mat.Rows, mat.Cols)
}

func backward(mat matrix.Matrix) matrix.Matrix {
	return  matrix.NewMatrix(make([]float64, 0), 0, 0)
}