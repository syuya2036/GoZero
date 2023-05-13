package steps

import (
	// "fmt"
	"fmt"
	"log"

	F "github.com/syuya2036/GoZero/functions"
	"github.com/syuya2036/GoZero/model"
)

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

	// backward
	y.Grad = []float64{1}
	b.Grad = C.BackFn(y.Grad)
	a.Grad = B.BackFn(b.Grad)
	x.Grad = A.BackFn(a.Grad)
	fmt.Println(x.Grad)
}