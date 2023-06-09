package main

import (
	"fmt"
	"log"

	F "github.com/syuya2036/GoZero/functions"
	"github.com/syuya2036/GoZero/model"
)

func main() {
	step06()
}

func step06() {
	// forward
	A := F.NewSquare()
	B := F.NewExp()
	C := F.NewSquare()

	x := model.NewScaler(0.5)

	a := A.Fn(x)
	b := B.Fn(a)
	y := C.Fn(b)

	err := y.PrintArray()
	if err != nil {
		log.Fatal(err)
	}

	y.Grad = []float64{1}
	y.Backward()
	fmt.Println(x.Grad)
}

