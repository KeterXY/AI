package main

import (
	"src/matrix"
	"src/optimizer"
)

func main() {
	t1 := []float64{3, 2, 1, 1}
	//t2 := []float64{1, 2, 3, 1}
	//t3 := []float64{5, 3, 1, 1}
	//t4 := []float64{1, 2, 1, 1}
	//t := [][]float64{t1, t2, t3, t4}
	t := [][]float64{t1}
	test1 := new(matrix.Matrix).SetMatrixValue(t)
	label := matrix.GenerateOneMatrix(1, 3)

	model := new(optimizer.Model).Init(test1, label)
	model.SetLayer(new(optimizer.FullConnect).SetCell(8))
	model.SetLayer(new(optimizer.Sigmoid))
	model.SetLayer(new(optimizer.FullConnect).SetCell(5))
	model.SetLayer(new(optimizer.Sigmoid))
	model.SetLayer(new(optimizer.FullConnect).SetCell(3))
	model.SetLayer(new(optimizer.MeanSquareLoss))

	model.NormInitLayers()

	model.GraduateOptimize(0.3,10000)

	model.Calculate()

}