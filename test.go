package main

import (
	// "fmt"
	"github.com/syuya2036/GoZero/matrix"
)

func main() {
	var mat matrix.Matrix = matrix.NewMatrix([]float64{1,2,3,4,5,6}, 2, 3)
	mat.Mat = []float64{1,3,5,7,9,11}
	mat.PrintArray()
}