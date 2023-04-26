package functions

import (
	"github.com/syuya2036/GoZero/matrix"
)


type FunctionBase struct{
    Name string
    Forward func(mat *matrix.Matrix) *matrix.Matrix
    Backward func(mat *matrix.Matrix) *matrix.Matrix
}

func (f *FunctionBase) Fn(mat *matrix.Matrix) (*matrix.Matrix, error) {
    result := f.Forward(mat)
    return result, nil
}

func (f *FunctionBase) ToString() string {
    return f.Name
}