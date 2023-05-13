package model


type FunctionBase struct{
    Name string
    Forward func(mat *Matrix) *Matrix
    Backward func(mat *Matrix, gy []float64) []float64
    Input *Matrix
    Output *Matrix
}

func (f *FunctionBase) Fn(mat *Matrix) *Matrix {
    result := f.Forward(mat)
    f.Input = mat
    result.Creator = f
    f.Output = result

    return result
}

func (f *FunctionBase) BackFn(gy []float64) []float64  {
    x := f.Input
    result := f.Backward(x, gy)
    return result
}

func (f *FunctionBase) String() string {
    return f.Name
}