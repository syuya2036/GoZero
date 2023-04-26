package functions

import (
	"github.com/syuya2036/GoZero/matrix"
	"math"
)

type Exp struct {
	FunctionBase
}

func NewExp() *Exp {
	var e *Exp = &Exp{}
	return e
}

func (e *Exp) Forward(mat *matrix.Matrix) *matrix.Matrix {
	result := make([]float64, mat.Rows*mat.Cols)
	for i := range result {
		result[i] = math.Exp(mat.Mat[i])
	}

	return matrix.NewMatrix(result, mat.Rows, mat.Cols)
}

func (e *Exp) Backward(mat *matrix.Matrix) *matrix.Matrix {
	return nil
}