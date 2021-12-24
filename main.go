package main

import (
	"fmt"
	"src/matrix"
)

func main() {
	t1 := []float64{3,2,1}
	t2 := []float64{1,2,3}
	t3 := []float64{5,3,1}
	t := [][]float64{t1,t2,t3}

	test1 := new(matrix.Matrix).SetMatrix(t)

	fmt.Println(test1.Max())
}

