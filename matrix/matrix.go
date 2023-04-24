package GoZero

import (
	"math"
    "fmt"
)

type Matrix struct {
	Mat []float64
	Rows int
	Cols int
}


func NewMatrix(mat []float64, rows, cols int) (*Matrix) {
	return &Matrix{Mat: mat, Rows: rows, Cols: cols}
}

func NewMatrixFromArray(mat [][]float64) (*Matrix) {
    rows := len(mat)
    cols := len(mat[0])
    matC := make([]float64, rows*cols)
    for i:=0; i<rows; i++ {
        for j:=0; j<cols; j++  {
            matC[i*cols+j] = mat[i][j]
        }
    }
    return NewMatrix(matC, rows, cols)
}

func Zeros(rows, cols int) (*Matrix, error) {
	if rows < 1 || cols < 1 {
		return nil, fmt.Errorf("invalid matrix shape: %d x %d", rows, cols)
	}

	mat := make([]float64, rows*cols)
	for i := range mat {
		mat[i] = 0
	}

	return &Matrix{Mat: mat, Rows: rows, Cols: cols}, nil
}


func Dot(matA, matB *Matrix) (*Matrix, error) {
    if matA.Cols != matB.Rows {
        return nil, fmt.Errorf("invalid matrix shape: %d x %d and %d x %d", matA.Rows, matA.Cols, matB.Rows, matB.Cols)
    }

    matC := make([]float64, matA.Rows*matB.Cols)
    for i:=0; i<matA.Rows; i++  {
        for j:=0; j<matB.Cols; j++  {
            var sum float64
            for k:=0; k<matA.Cols; k++  {
                sum += matA.Mat[i*matA.Cols+k] * matB.Mat[k*matB.Cols+j]
            }
            matC[i*matB.Cols+j] = sum
        }
    }

    return NewMatrix(matC, matA.Rows, matB.Cols), nil
}



// 配列同士
func Add(matA, matB *Matrix) (*Matrix, error) {
    if matA.Rows != matB.Rows || matA.Cols != matB.Cols {
        return nil, fmt.Errorf("matrices must have same dimensions")
    }

    result := make([]float64, matA.Rows*matA.Cols)
    for i := range result {
        result[i] = matA.Mat[i] + matB.Mat[i]
    }

    return NewMatrix(result, matA.Rows, matA.Cols), nil
}

func Sub(matA, matB *Matrix) (*Matrix, error) {
    if matA.Rows != matB.Rows || matA.Cols != matB.Cols {
        return nil, fmt.Errorf("matrices must have same dimensions")
    }

    result := make([]float64, matA.Rows*matA.Cols)
    for i := range result {
        result[i] = matA.Mat[i] - matB.Mat[i]
    }

    return NewMatrix(result, matA.Rows, matA.Cols), nil
}

func Mul(matA, matB *Matrix) (*Matrix, error) {
    if matA.Rows != matB.Rows || matA.Cols != matB.Cols {
        return nil, fmt.Errorf("matrices must have same dimensions")
    }

    result := make([]float64, matA.Rows*matA.Cols)
    for i := range result {
        result[i] = matA.Mat[i] * matB.Mat[i]
    }

    return NewMatrix(result, matA.Rows, matA.Cols), nil
}

func Div(matA, matB *Matrix) (*Matrix, error) {
    if matA.Rows != matB.Rows || matA.Cols != matB.Cols {
        return nil, fmt.Errorf("matrices must have same dimensions")
    }

    result := make([]float64, matA.Rows*matA.Cols)
    for i := range result {
        if abs(matB.Mat[i]) < 1e-14 {
            return nil, fmt.Errorf("division by zero")
        } else {
        result[i] = matA.Mat[i] / matB.Mat[i]
        }
    }

    return NewMatrix(result, matA.Rows, matA.Cols), nil
}


func (mat *Matrix) AddScaler(scaler float64) (*Matrix, error) {
    result := make([]float64, mat.Rows*mat.Cols)
    for i := range result {
        result[i] = mat.Mat[i] + scaler
    }

    return NewMatrix(result, mat.Rows, mat.Cols), nil
}

func (mat *Matrix) SubScaler(scaler float64) (*Matrix, error) {
    result := make([]float64, mat.Rows*mat.Cols)
    for i := range result {
        result[i] = mat.Mat[i] - scaler
    }

    return NewMatrix(result, mat.Rows, mat.Cols), nil
}

func (mat *Matrix) MulScaler(scaler float64) (*Matrix, error) {
    result := make([]float64, mat.Rows*mat.Cols)
    for i := range result {
        result[i] = mat.Mat[i] * scaler
    }

    return NewMatrix(result, mat.Rows, mat.Cols), nil
}

func (mat *Matrix) DivScaler(scaler float64) (*Matrix, error) {
    result := make([]float64, mat.Rows*mat.Cols)
    for i := range result {
        if abs(scaler) < 1e-14 {
            return nil, fmt.Errorf("division by zero")
            } else {
        result[i] = mat.Mat[i] / scaler
            }
    }

    return NewMatrix(result, mat.Rows, mat.Cols), nil
}

func (mat *Matrix) Pow(x float64) (*Matrix, error) {
    result := make([]float64, mat.Rows*mat.Cols)
    for i := range result {
        result[i] = math.Pow(mat.Mat[i], x)
    }

    return NewMatrix(result, mat.Rows, mat.Cols), nil
}


func (mat *Matrix) T() (*Matrix, error) {
	result := make([]float64, mat.Cols*mat.Rows)
	for i:=0; i<mat.Rows; i++ {
        for j:=0; j<mat.Cols; j++ {
            result[j*mat.Rows+i] = mat.Mat[i*mat.Cols+j]
        }
	}

	return NewMatrix(result, mat.Cols, mat.Rows), nil
}

func abs(n float64) float64 {
	if n < 0 {
		return -n
	} else {
		return n
	}
}