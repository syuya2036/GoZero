package functions

import (
	"github.com/syuya2036/GoZero/matrix"
	"math"
)

func NewExp() *FunctionBase {
	var e *FunctionBase = &FunctionBase{Name: "Square", Forward: forward, Backward: backward}
	return e
}

func Forward(mat matrix.Matrix) matrix.Matrix {
	result := make([]float64, mat.Rows*mat.Cols)
	for i := range result {
		result[i] = math.Exp(mat.Mat[i])
	}

	return matrix.NewMatrix(result, mat.Rows, mat.Cols)
}

func Backward(mat *matrix.Matrix) matrix.Matrix {
	return matrix.NewMatrix(make([]float64, 0), 0, 0)
}