package optimizer

import (
	"context"
	"fmt"
	"src/activation"
	"src/loss"
	"src/matrix"
)

func(model *Model)GraduateOptimize(learningRate float64, chunk int)  {
	for i:=0;i<chunk;i++ {
		model.graduateOptimize(learningRate)
	}
}

func(model *Model)graduateOptimize(learningRate float64){
	temp := new(matrix.Matrix).SetMatrix(model.input)
	label := new(matrix.Matrix).SetMatrix(model.label)
	srcStack := generateStack()
	count := 0

	var dOut *matrix.Matrix

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
			outputLoss,err := loss.MeanSquare(temp, label)
			if err != nil{
				fmt.Println("loss error")
				context.TODO()	//均方误差错误
			}else {
				fmt.Println(outputLoss)
			}
			dOut,err = loss.MeanSquareBack(temp, label)
			if err != nil{
				fmt.Println(err)
			}
		default:
			fmt.Println("unknown type")
		}
	}

	var body interface{}
	for index:=count-1;index>0;index--{
		srcStack,body = srcStack.pop()
		switch t:=body.(type) {
		case *FullConnect:
			inputTrans := t.input.Transport()
			bodyTrans := t.layer.Transport()
			deltaMatrix,err := new(matrix.Matrix).MatrixCross(inputTrans,dOut)
			if err != nil{
				fmt.Println(err)
			}
			deltaMatrix,err = matrix.MatrixMultiply(deltaMatrix,learningRate)
			if err != nil{
				fmt.Println(err)
			}
			newLayer,_ := new(matrix.Matrix).Sub(t.layer,deltaMatrix)
			t.layer = newLayer
			newLoss,_ := new(matrix.Matrix).MatrixCross(dOut,bodyTrans)
			dOut = newLoss
			model.layers[index] = t
		case *Relu:
			dOut = t.body.Backup(dOut)
		case *Sigmoid:
			dOut = t.body.Back(dOut)
		default:
			fmt.Println("unknown type")
		}
	}
}
