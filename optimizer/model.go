package optimizer

import (
	"context"
	"fmt"
	"src/activation"
	"src/matrix"
)

type Model struct {
	layers []interface{}
	input *matrix.Matrix
	label *matrix.Matrix
}

type FullConnect struct {
	layer *matrix.Matrix
	cellCount int
	input *matrix.Matrix
}

type Relu struct {
	body *activation.Relu
}

type Sigmoid struct {
	body *activation.Sigmoid
}

type MeanSquareLoss struct {

}

func (layer *FullConnect)SetCell(cellCount int) *FullConnect {
	layer.cellCount = cellCount
	return layer
}

func (model *Model)SetLayer(body interface{})  {
	switch body.(type) {
	case *Sigmoid:
		context.TODO()
	case *FullConnect:
		context.TODO()
	case *Relu:
		context.TODO()
	case *MeanSquareLoss:
		context.TODO()
	default:
		fmt.Println("unknown type")
	}
	model.layers = append(model.layers, body)
}

func (model *Model)Init(input,label *matrix.Matrix) *Model {
	model.label = label
	model.input = input
	return model
}

func (model *Model)NormInitLayers()  {
	temp := new(matrix.Matrix).SetMatrix(model.input)
	for _,layer := range model.layers{
		switch t := layer.(type) {
		case *FullConnect:
			body := matrix.RandNormMatrix(temp.Shape[1],t.cellCount)
			t.layer = body
			temp,_ = new(matrix.Matrix).MatrixCross(temp,body)
		case *Relu:
			context.TODO()
		case *Sigmoid:
			context.TODO()
		case *MeanSquareLoss:
			context.TODO()
		default:
			fmt.Println("Unknown type!")
		}
	}
}

func (model *Model)RandInitLayers()  {
	temp := new(matrix.Matrix).SetMatrix(model.input)
	for _,layer := range model.layers{
		switch t := layer.(type) {
		case *FullConnect:
			body := matrix.RandMatrix(temp.Shape[1],t.cellCount)
			t.layer = body
			temp,_ = new(matrix.Matrix).MatrixCross(temp,body)
		case *Relu:
			context.TODO()
		case Sigmoid:
			context.TODO()
		case MeanSquareLoss:
			context.TODO()
		default:
			fmt.Println("Unknown type!")
		}
	}
}

func (model *Model)Check()  {
	fmt.Printf("input shape %d * %d \n", model.input.Shape[0], model.input.Shape[1])
	temp := new(matrix.Matrix).SetMatrix(model.input)
	for _,layer := range model.layers{
		switch t := layer.(type) {
		case *FullConnect:
			if temp.Shape[1] != t.layer.Shape[0]{
				context.TODO()	//纬度错误
			}
			temp,_ = new(matrix.Matrix).MatrixCross(temp,t.layer)		//全链接层
			fmt.Printf("full connection layer with matrix shape %d * %d \n",t.layer.Shape[0], t.layer.Shape[1] )
		case *Relu:
			fmt.Println("relu activation")
		case *Sigmoid:
			fmt.Println("sigmoid activation")
		case *MeanSquareLoss:
			fmt.Printf("mean square loss with label shape %d * %d \n", model.label.Shape[0], model.label.Shape[1])
		default:
			fmt.Println("unknown type")
		}
	}
}

func (model *Model)Calculate()  {
	temp := new(matrix.Matrix).SetMatrix(model.input)
	srcStack := generateStack()
	count := 0
	for _,layer := range model.layers{
		switch t := layer.(type) {
		case *FullConnect:
			if temp.Shape[1] != t.layer.Shape[0]{
				context.TODO()	//纬度错误
			}
			t.input = new(matrix.Matrix).SetMatrix(temp)
			temp,_ = new(matrix.Matrix).MatrixCross(temp,t.layer)		//全链接层
			srcStack = srcStack.push(t)		//推入栈内
			count += 1
		case *Relu:
			relu := new(activation.Relu).Forward(temp)
			t.body = relu
			temp = relu.GetMatrix()
			srcStack = srcStack.push(t)
			count += 1
		case *Sigmoid:
			sigmoid := new(activation.Sigmoid).Forward(temp)
			t.body = sigmoid
			temp = sigmoid.GetMatrix()
			srcStack = srcStack.push(t)
			count += 1
		case *MeanSquareLoss:
			fmt.Println("final output:")
			temp.Print()
		default:
			fmt.Println("unknown type")
		}
	}

}

