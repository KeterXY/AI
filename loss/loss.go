package loss

import "src/matrix"

func MeanSquare(inputMatrix,label *matrix.Matrix)(float64,error){
	total := 0.0
	count := float64(inputMatrix.Shape[0] * inputMatrix.Shape[1])
	for index1:=0;index1<inputMatrix.Shape[0];index1++{
		for index2:=0;index2<inputMatrix.Shape[1];index2++{
			var value float64
			value = (*inputMatrix.Matrix[index1][index2] - *label.Matrix[index1][index2])*(*inputMatrix.Matrix[index1][index2] - *label.Matrix[index1][index2])
			total += value
		}
	}
	output := total / (count)
	return output, nil
}

func MeanSquareBack(inputMatrix,label *matrix.Matrix)(*matrix.Matrix,error)  {
	if inputMatrix.Shape[0] != label.Shape[0] || inputMatrix.Shape[1] != label.Shape[1]{
		return nil,matrix.ShapeError
	}
	total := 0.0
	count := float64(inputMatrix.Shape[0] * inputMatrix.Shape[1])
	newMatrix := make([][]*float64, inputMatrix.Shape[0])
	for index1:=0;index1<inputMatrix.Shape[0];index1++{
		newVector := make([]*float64, inputMatrix.Shape[1])
		for index2:=0;index2<inputMatrix.Shape[1];index2++{
			var value float64
			value = *inputMatrix.Matrix[index1][index2] - *label.Matrix[index1][index2]
			newVector[index2] = &value
			total += value * value
		}
		newMatrix[index1] = newVector
	}
	output := total / (count)
	deltaOutput,_ := matrix.MatrixMultiply(&matrix.Matrix{Shape: inputMatrix.Shape, Matrix: newMatrix}, output)
	return deltaOutput, nil
}