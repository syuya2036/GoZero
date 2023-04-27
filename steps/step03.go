package main

import (
	// "fmt"
	"log"

	F "github.com/syuya2036/GoZero/functions"
	"github.com/syuya2036/GoZero/matrix"
)

func main() {
	A := F.NewSquare()
	B := F.NewExp()
	C := F.NewSquare()

	x := matrix.NewMatrixFromArray([][]float64{
		{1, 2},
		{3, 4}})

	a := A.Fn(x)
	b := B.Fn(a)
	y := C.Fn(b)
	
	err := y.PrintArray()
	if err != nil {
		log.Fatal(err)
	}
}