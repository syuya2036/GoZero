package model

import (
	"fmt"
	"math"
)

type Matrix struct {
	Mat  []float64
	Grad []float64
	Creator *FunctionBase
	Rows int
	Cols int
}

func NewMatrix(mat []float64, rows, cols int) *Matrix {
	var self *Matrix = new(Matrix)

	self.Mat = mat
	self.Rows, self.Cols = rows, cols

	return self
}

func NewScaler(value float64) *Matrix {
	var self *Matrix = new(Matrix)

	scalerMat := make([]float64, 1)
	scalerMat[0] = value
	self.Mat = scalerMat

	self.Rows, self.Cols = 1, 1

	return self
}

// 二次元配列から初期化
func NewMatrixFromArray(mat [][]float64) *Matrix {

	rows, cols := len(mat), len(mat[0])

	matC := make([]float64, rows*cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			matC[i*cols+j] = mat[i][j]
		}
	}

	return NewMatrix(matC, rows, cols)
}

// 自信を出力した関数を覚える
func (mat *Matrix) SetCreator(function FunctionBase) {
	mat.Creator = &function
}

// 要素がすべて0の配列
func Zeros(rows, cols int) (*Matrix, error) {
	if rows < 1 || cols < 1 {
		return NewMatrix(make([]float64, 0), 0, 0), fmt.Errorf("invalid matrix shape: %d x %d", rows, cols)
	}

	mat := make([]float64, rows*cols)
	for i := range mat {
		mat[i] = 0
	}

	return NewMatrix(mat, rows, cols), nil
}

func (mat *Matrix) Backward() {
	f := mat.Creator
	if f != nil {
		x := f.Input
		x.Grad = f.BackFn(mat.Grad)
		x.Backward()
		fmt.Print(x.Grad)
		fmt.Println(f.Name)
	}
}

// 行列積を計算
func Dot(matA, matB Matrix) (*Matrix, error) {
	if matA.Cols != matB.Rows {
		return NewMatrix(make([]float64, 0), 0, 0), fmt.Errorf("invalid matrix shape: %d x %d and %d x %d", matA.Rows, matA.Cols, matB.Rows, matB.Cols)
	}

	result := make([]float64, matA.Rows*matB.Cols)

	for i := 0; i < matA.Rows; i++ {
		for j := 0; j < matB.Cols; j++ {

			// matAのi行k列とmatBのk行j列の積の和
			var sum float64
			for k := 0; k < matA.Cols; k++ {
				sum += matA.Mat[i*matA.Cols+k] * matB.Mat[k*matB.Cols+j]
			}

			result[i*matB.Cols+j] = sum
		}
	}

	return NewMatrix(result, matA.Rows, matB.Cols), nil
}

// 配列の要素同士の足し算
func Add(matA, matB Matrix) (*Matrix, error) {
	if matA.Rows != matB.Rows || matA.Cols != matB.Cols {
		return NewMatrix(make([]float64, 0), 0, 0), fmt.Errorf("matrices must have same dimensions")
	}

	result := make([]float64, matA.Rows*matA.Cols)
	for i := range result {
		result[i] = matA.Mat[i] + matB.Mat[i]
	}

	return NewMatrix(result, matA.Rows, matA.Cols), nil
}

// 配列の要素同士の引き算
func Sub(matA, matB Matrix) (*Matrix, error) {
	if matA.Rows != matB.Rows || matA.Cols != matB.Cols {
		return NewMatrix(make([]float64, 0), 0, 0), fmt.Errorf("matrices must have same dimensions")
	}

	result := make([]float64, matA.Rows*matA.Cols)
	for i := range result {
		result[i] = matA.Mat[i] - matB.Mat[i]
	}

	return NewMatrix(result, matA.Rows, matA.Cols), nil
}

// 配列の要素同士の掛ぎ算
func Mul(matA, matB Matrix) (*Matrix, error) {
	if matA.Rows != matB.Rows || matA.Cols != matB.Cols {
		return NewMatrix(make([]float64, 0), 0, 0), fmt.Errorf("matrices must have same dimensions")
	}
	
	result := make([]float64, matA.Rows*matA.Cols)
	for i := range result {
		result[i] = matA.Mat[i] * matB.Mat[i]
	}

	return NewMatrix(result, matA.Rows, matA.Cols), nil
}

// 配列の要素同士の割り算
func Div(matA, matB Matrix) (*Matrix, error) {
	if matA.Rows != matB.Rows || matA.Cols != matB.Cols {
		return NewMatrix(make([]float64, 0), 0, 0), fmt.Errorf("matrices must have same dimensions")
	}

	result := make([]float64, matA.Rows*matA.Cols)
	for i := range result {
		if abs(matB.Mat[i]) < 1e-14 {
			return NewMatrix(make([]float64, 0), 0, 0), fmt.Errorf("division by zero")
		} else {
			result[i] = matA.Mat[i] / matB.Mat[i]
		}
	}

	return NewMatrix(result, matA.Rows, matA.Cols), nil
}

// 配列の全要素にscalerを足す
func (mat Matrix) AddScaler(scaler float64) (*Matrix, error) {
	if mat.Rows == 0 {
		return NewMatrix(make([]float64, 0), 0, 0), fmt.Errorf("invalid matrix shape: %d x %d", mat.Rows, mat.Cols)
	}
	result := make([]float64, mat.Rows*mat.Cols)

	for i := range result {
		result[i] = mat.Mat[i] + scaler
	}

	return NewMatrix(result, mat.Rows, mat.Cols), nil
}

// 配列の全要素からscalerを引く
func (mat Matrix) SubScaler(scaler float64) (*Matrix, error) {
	if mat.Rows == 0 {
		return NewMatrix(make([]float64, 0), 0, 0), fmt.Errorf("invalid matrix shape: %d x %d", mat.Rows, mat.Cols)
	}
	result := make([]float64, mat.Rows*mat.Cols)

	for i := range result {
		result[i] = mat.Mat[i] - scaler
	}

	return NewMatrix(result, mat.Rows, mat.Cols), nil
}

// 配列の全要素にscalerを掛ける
func (mat Matrix) MulScaler(scaler float64) (*Matrix, error) {
	if mat.Rows == 0 {
		return NewMatrix(make([]float64, 0), 0, 0), fmt.Errorf("invalid matrix shape: %d x %d", mat.Rows, mat.Cols)
	}
	result := make([]float64, mat.Rows*mat.Cols)

	for i := range result {
		result[i] = mat.Mat[i] * scaler
	}

	return NewMatrix(result, mat.Rows, mat.Cols), nil
}

// 配列の全要素をscalerで割る
func (mat Matrix) DivScaler(scaler float64) (*Matrix, error) {
	if mat.Rows == 0 {
		return NewMatrix(make([]float64, 0), 0, 0), fmt.Errorf("invalid matrix shape: %d x %d", mat.Rows, mat.Cols)
	}
	result := make([]float64, mat.Rows*mat.Cols)

	for i := range result {
		if abs(scaler) < 1e-14 {
			return NewMatrix(make([]float64, 0), 0, 0), fmt.Errorf("division by zero")
		} else {
			result[i] = mat.Mat[i] / scaler
		}
	}

	return NewMatrix(result, mat.Rows, mat.Cols), nil
}

// 配列の全要素をscaler乗する
func (mat Matrix) Pow(x float64) (*Matrix, error) {
	if mat.Rows == 0 {
		return NewMatrix(make([]float64, 0), 0, 0), fmt.Errorf("invalid matrix shape: %d x %d", mat.Rows, mat.Cols)
	}
	result := make([]float64, mat.Rows*mat.Cols)

	for i := range result {
		result[i] = math.Pow(mat.Mat[i], x)
	}

	return NewMatrix(result, mat.Rows, mat.Cols), nil
}

// 行列を転置する
func (mat Matrix) T() (*Matrix, error) {
	if mat.Rows == 0 {
		return NewMatrix(make([]float64, 0), 0, 0), fmt.Errorf("invalid matrix shape: %d x %d", mat.Rows, mat.Cols)
	}
	result := make([]float64, mat.Cols*mat.Rows)

	for i := 0; i < mat.Rows; i++ {
		for j := 0; j < mat.Cols; j++ {
			result[j*mat.Rows+i] = mat.Mat[i*mat.Cols+j]
		}
	}

	return NewMatrix(result, mat.Cols, mat.Rows), nil
}

// 引数の絶対値を返す
func abs(n float64) float64 {
	if n < 0 {
		return -n
	} else {
		return n
	}
}

func (mat Matrix) PrintArray() (error) {
	if mat.Rows == 0 {
		return fmt.Errorf("matrix is nil")
	}
	for i := 0; i < mat.Rows; i++ {
		for j := 0; j < mat.Cols; j++ {
			fmt.Printf("%f, ", mat.Mat[i*mat.Cols+j])
		}
		fmt.Print("\n")
	}

	return nil
}

func (mat Matrix) PrintGrad() (error) {
	if mat.Rows == 0 {
		return fmt.Errorf("matrix is nil")
	}
	for i := 0; i < mat.Rows; i++ {
		for j := 0; j < mat.Cols; j++ {
			fmt.Printf("%f, ", mat.Grad[i*mat.Cols+j])
		}
		fmt.Print("\n")
	}

	return nil
}