package main

import (
	"fmt"
	// "github.com/syuya2036/GoZero/steps"
	"github.com/syuya2036/GoZero/matrix"
	"log"
)

func main() {
	matA := GoZero.NewMatrixFromArray([][]float64{
		{1, 2},
		{3, 4}})

	matB := GoZero.NewMatrixFromArray([][]float64{
		{1, 2},
		{3, 4}})

	matC, err := GoZero.Dot(matB, matA)
	if err != nil {
		log.Fatal(err)
	}

	matZero, err := GoZero.Zeros(10, 3)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println(matC)
	fmt.Println(matZero)
}