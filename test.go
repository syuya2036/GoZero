package main

import (
	// "fmt"
	"log"

	"github.com/syuya2036/GoZero/functions"
	"github.com/syuya2036/GoZero/matrix"
)

func main() {
	x := matrix.NewMatrixFromArray([][]float64{
		{1, 2},
		{3, 4}})
	
	f := functions.NewSquare()
	y := f.Forward(x)

	err := y.PrintArray()
	if err != nil {
		log.Fatal(err)
	}
}