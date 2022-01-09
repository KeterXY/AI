package matrix

import (
	"fmt"
	"math/rand"
	"time"
	"unsafe"
)

type Vector struct {
	Length int
	Vector []*float64
}


type Matrix struct {
	Shape [2]int
	Matrix [][]*float64
}


func RandNormVector(length int) *Vector {
	var vector []*float64
	rand.Seed(time.Now().Unix())
	for i:=0;i<length;i++{
		value := rand.NormFloat64()
		vector = append(vector, &value)
	}
	return &Vector{Length: length, Vector: vector}
}

func RandVector(length int) *Vector {
	var vector []*float64
	rand.Seed(time.Now().Unix())
	for i:=0;i<length;i++{
		value := rand.Float64() - 0.5
		vector = append(vector, &value)
	}
	return &Vector{Length: length, Vector: vector}
}

func GenerateOneVector(length int)*Vector{
	var vector []*float64
	for i:= 0;i<length;i++{
		var one float64 = 1
		vector = append(vector, &one)
	}
	return &Vector{Length: length, Vector: vector}
}

func GenerateZeroVector(length int)*Vector{
	var vector []*float64
	for i:= 0;i<length;i++{
		var one float64 = 0
		vector = append(vector, &one)
	}
	return &Vector{Length: length, Vector: vector}
}

func(vector *Vector) Print()  {
	fmt.Printf("Length: %d",vector.Length)
	printVector(vector)
}

func printVector(vector *Vector){
	fmt.Print(" [ ")
	if vector.Length > 10{
		for i := 0; i < 5; i++{
			fmt.Printf("%f ",*vector.Vector[i])
		}
		fmt.Print("... ")
		for i := 5; i > 0;i--{
			fmt.Printf("%f ",*vector.Vector[vector.Length-i])
		}
	}else {
		for i := 0; i< vector.Length; i++{
			fmt.Printf("%f ",*vector.Vector[i])
		}
	}
	fmt.Print(" ]\n")
}

func(vector *Vector) SetValue(index int,value float64) error {
	if index >= vector.Length{
		return IndexError
	}
	ptr := (*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(vector.Vector[0])) + uintptr(index)*unsafe.Sizeof(vector.Vector[0])))
	*ptr = value
	return nil
}

func RandNormMatrix(row,col int) *Matrix {
	rand.Seed(time.Now().Unix())
	var matrix [][]*float64
	for i:=0 ;i < row; i++{
		var vector []*float64
		for j:=0;j<col;j++{
			value := rand.NormFloat64()
			vector = append(vector, &value)
		}
		matrix = append(matrix, vector)
	}
	return &Matrix{Shape: [2]int{row,col}, Matrix: matrix}
}

func RandMatrix(row,col int) *Matrix {
	rand.Seed(time.Now().Unix())
	var matrix [][]*float64
	for i:=0 ;i < row; i++{
		var vector []*float64
		for j:=0;j<col;j++{
			value := rand.Float64() - 0.5
			vector = append(vector, &value)
		}
		matrix = append(matrix, vector)
	}
	return &Matrix{Shape: [2]int{row,col}, Matrix: matrix}
}

func GenerateOneMatrix(row,col int) *Matrix {
	rand.Seed(time.Now().Unix())
	var matrix [][]*float64
	for i:=0 ;i < row; i++{
		var vector []*float64
		for j:=0;j<col;j++{
			value := 1.0
			vector = append(vector, &value)
		}
		matrix = append(matrix, vector)
	}
	return &Matrix{Shape: [2]int{row,col}, Matrix: matrix}
}

func GenerateZeroMatrix(row,col int) *Matrix {
	rand.Seed(time.Now().Unix())
	var matrix [][]*float64
	for i:=0 ;i < row; i++{
		var vector []*float64
		for j:=0;j<col;j++{
			value := 0.0
			vector = append(vector, &value)
		}
		matrix = append(matrix, vector)
	}
	return &Matrix{Shape: [2]int{row,col}, Matrix: matrix}
}

func (matrix *Matrix)Print()  {
	fmt.Printf("Shape: %d rows %d columns\n",matrix.Shape[0],matrix.Shape[1])
	if matrix.Shape[0] > 10{
		for i:=0; i<5; i++{
			printVector(&Vector{Length: matrix.Shape[1],Vector: matrix.Matrix[i]})
		}
		fmt.Println(" ...")
		for i:= 5; i>0; i--{
			printVector(&Vector{Length: matrix.Shape[1],Vector:matrix.Matrix[matrix.Shape[0] - i]})
		}
	}else{
		for i:= 0; i<matrix.Shape[0]; i++{
			printVector(&Vector{Length: matrix.Shape[1],Vector: matrix.Matrix[i]})
		}
	}
}

func (matrix *Matrix)SetValue(rowIndex,colIndex int, value float64) error {
	if rowIndex >= matrix.Shape[0] || colIndex >= matrix.Shape[1]{
		return IndexError
	}else if rowIndex < 0 || colIndex < 0{
		return IndexError
	}
	matrix.Matrix[rowIndex][colIndex] = &value
	return nil
}

func (matrix *Matrix)GetColumn(index int) ([]float64,error) {
	if index < 0 || index >= matrix.Shape[1]{
		return nil,IndexError
	}
	output := make([]float64,matrix.Shape[0])
	for temp:=0;temp<matrix.Shape[0];temp++{
		output[temp] = *matrix.Matrix[temp][index]
	}
	return output,nil
}

func (matrix *Matrix)getColumn(index int) []*float64 {
	output := make([]*float64,matrix.Shape[0])
	for temp:=0;temp<matrix.Shape[0];temp++{
		output[temp] = matrix.Matrix[temp][index]
	}
	return output
}

func (matrix *Matrix)SetArray(input []float64) *Matrix {
	newMatrix := make([][]*float64, 1)
	newVector := make([]*float64, len(input))
	for index:=0; index < len(input); index++{
		var value = input[index]
		newVector[index] = &value
	}
	newMatrix[0] = newVector
	return &Matrix{Shape: [2]int{1,len(input)}, Matrix: newMatrix}
}

func (matrix *Matrix)SetMatrix(input *Matrix)*Matrix{
	newMatrix := make([][]*float64, input.Shape[0])

	for index:=0; index < input.Shape[0]; index++{
		newVector := make([]*float64, input.Shape[1])
		for key,val := range input.Matrix[index]{
			var newValue float64 = *val
			newVector[key] = &newValue
		}
		newMatrix[index] = newVector
	}
	return &Matrix{Shape: [2]int{input.Shape[0],input.Shape[1]}, Matrix: newMatrix}
}

func (matrix *Matrix)SetMatrixValue(input [][]float64)*Matrix{
	newMatrix := make([][]*float64, len(input))
	for index:=0; index < len(input); index++{
		newVector := make([]*float64, len(input[0]))
		for index1:=0; index1<len(input[0]);index1++ {
			var value = input[index][index1]
			newVector[index1] = &value
		}
		newMatrix[index] = newVector
	}
	return &Matrix{Shape: [2]int{len(input),len(input[0])}, Matrix: newMatrix}
}

func (matrix *Matrix)ToValue()[][]float64{
	output := make([][]float64, matrix.Shape[0])
	for index,list := range matrix.Matrix{
		newVector := make([]float64, matrix.Shape[1])
		for key,val := range list{
			newVector[key] = *val
		}
		output[index] = newVector
	}
	return output
}