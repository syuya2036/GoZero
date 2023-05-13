package functions

import (
	"github.com/syuya2036/GoZero/model"
	"math"
)

func NewSquare() *model.FunctionBase {
	var s *model.FunctionBase = &model.FunctionBase{Name: "Square", Forward: SquareForward, Backward: SquareBackward}
	return	s
}

func SquareForward(mat *model.Matrix) *model.Matrix {
	result := make([]float64, mat.Rows*mat.Cols)
	for i := range result  {
		result[i] = math.Pow(mat.Mat[i], 2)
	}

	return model.NewMatrix(result, mat.Rows, mat.Cols)
}

func SquareBackward(mat *model.Matrix, gy []float64) []float64{
	result := make([]float64, mat.Rows*mat.Cols)
	for i := range result  {
		result[i] = 2 * mat.Mat[i] * gy[i]
	}
	return result
}