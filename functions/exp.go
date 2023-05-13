package functions

import (
	"github.com/syuya2036/GoZero/model"
	"math"
)

func NewExp() *model.FunctionBase {
	var e *model.FunctionBase = &model.FunctionBase{Name: "Exp", Forward: ExpForward, Backward: ExpBackward}
	return e
}

func ExpForward(mat *model.Matrix) *model.Matrix {
	result := make([]float64, mat.Rows*mat.Cols)
	for i := range result {
		result[i] = math.Exp(mat.Mat[i])
	}

	return model.NewMatrix(result, mat.Rows, mat.Cols)
}

func ExpBackward(mat *model.Matrix, gy []float64) []float64 {
	result := make([]float64, mat.Rows*mat.Cols)
	for i := range result {
		result[i] = gy[i] * math.Exp(mat.Mat[i])
	}
	return result
}