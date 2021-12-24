package activation

import (
	"math"
	"src/matrix"
)

type Relu struct {
	layer *matrix.Matrix
}

type Sigmoid struct {
	layer *matrix.Matrix
}

func (relu *Relu)Forward(input *matrix.Matrix)*Relu{
	newMatrix := make([][]*float64,input.Shape[0])
	for index1:=0; index1<input.Shape[0]; index1++{
		newVector := make([]*float64, input.Shape[1])
		for index2:=0; index2<input.Shape[1];index2++{
			var value float64
			if  *input.Matrix[index1][index2] > 0{
				value = *input.Matrix[index1][index2]
			}else {
				value = 0
			}
			newVector[index2] = &value
		}
		newMatrix[index1] = newVector
	}
	relu.layer = &matrix.Matrix{Shape: [2]int{input.Shape[0],input.Shape[1]}, Matrix: newMatrix}
	return relu
}

func (sigmoid *Sigmoid)Forward(input *matrix.Matrix)*Sigmoid{
	newMatrix := make([][]*float64,input.Shape[0])
	for index1:=0;index1<input.Shape[0];index1++{
		newVector := make([]*float64,input.Shape[1])
		for index2:=0;index2<input.Shape[1];index2++{
			value := math.Exp(-*input.Matrix[index1][index2])		//e ^ (-x)
			value += 1; value = 1/value		// 1 / (1 + e ^ (-x))
			newVector[index2] = &value
		}
		newMatrix[index1] = newVector
	}
	sigmoid.layer = &matrix.Matrix{Shape: [2]int{input.Shape[0],input.Shape[1]}, Matrix: newMatrix}
	return sigmoid
}

func (relu *Relu)Backup(dOut *matrix.Matrix) *matrix.Matrix {
	newMatrix := make([][]*float64,relu.layer.Shape[0])
	for index1:=0; index1<relu.layer.Shape[0]; index1++{
		newVector := make([]*float64, relu.layer.Shape[1])
		for index2:=0; index2<relu.layer.Shape[1];index2++{
			var value float64
			if  *relu.layer.Matrix[index1][index2] > 0{
				value = *dOut.Matrix[index1][index2]
			}else {
				value = 0
			}
			newVector[index2] = &value
		}
		newMatrix[index1] = newVector
	}
	return &matrix.Matrix{Shape: [2]int{relu.layer.Shape[0],relu.layer.Shape[1]}, Matrix: newMatrix}
}

func(sigmoid *Sigmoid)Back(dOut *matrix.Matrix) *matrix.Matrix {
	newMatrix := make([][]*float64,sigmoid.layer.Shape[0])
	for index1:=0; index1<sigmoid.layer.Shape[0]; index1++{
		newVector := make([]*float64, sigmoid.layer.Shape[1])
		for index2:=0; index2<sigmoid.layer.Shape[1];index2++{
			value := *sigmoid.layer.Matrix[index1][index2]*(1- *sigmoid.layer.Matrix[index1][index2])
			value = value * *dOut.Matrix[index1][index2]
			newVector[index2] = &value
		}
		newMatrix[index1] = newVector
	}
	return &matrix.Matrix{Shape: [2]int{sigmoid.layer.Shape[0],sigmoid.layer.Shape[1]}, Matrix: newMatrix}
}