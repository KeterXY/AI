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

func (sigmoid *Sigmoid)GetMatrix()*matrix.Matrix{
	newMatrix := new(matrix.Matrix).SetMatrix(sigmoid.layer)
	return newMatrix
}

func (sigmoid *Sigmoid)Forward(input *matrix.Matrix)*Sigmoid{
	newMatrix := make([][]*float64,input.Shape[0])
	ch := make(chan bool, input.Shape[0])
	for index1:=0;index1<input.Shape[0];index1++{
		go sigmoidForward(input.Matrix[index1],index1,ch,newMatrix)
	}
	for index2:=0;index2<input.Shape[0];index2++{
		<- ch
	}
	sigmoid.layer = &matrix.Matrix{Shape: [2]int{input.Shape[0],input.Shape[1]}, Matrix: newMatrix}
	return sigmoid
}

func sigmoidForward(lis []*float64, index int, ch chan bool, output [][]*float64){
	newVector := make([]*float64,len(lis))
	for index2:=0;index2<len(lis);index2++{
		value := math.Exp(-*lis[index2])		//e ^ (-x)
		value += 1; value = 1/value		// 1 / (1 + e ^ (-x))
		newVector[index2] = &value
	}
	output[index] = newVector
	ch <- true
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
	ch := make(chan bool, dOut.Shape[0])
	for index1:=0; index1<sigmoid.layer.Shape[0]; index1++{
		go sigmoidBack(sigmoid.layer.Matrix[index1], dOut.Matrix[index1], index1, ch, newMatrix)
	}
	for i:=0;i<sigmoid.layer.Shape[0];i++{
		<- ch
	}
	return &matrix.Matrix{Shape: [2]int{sigmoid.layer.Shape[0],sigmoid.layer.Shape[1]}, Matrix: newMatrix}
}

func sigmoidBack(layerVec, dOutVec []*float64, index int, ch chan bool, outputMatrix [][]*float64){
	newVector := make([]*float64, len(layerVec))
	for index2,val := range layerVec{
		value := *val*(1- *val)
		value = value * *dOutVec[index2]
		newVector[index2] = &value
	}
	outputMatrix[index] = newVector
	ch <- true
}

func (relu *Relu)GetMatrix()*matrix.Matrix{
	newMatrix := new(matrix.Matrix).SetMatrix(relu.layer)
	return newMatrix
}

func (sigmoid *Sigmoid)CheckValue()  {
	sigmoid.layer.Print()
}