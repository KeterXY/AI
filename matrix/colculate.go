package matrix

func VectorCross(){

}

func VectorMultiply(vector *Vector, number float64)(*Vector,error){
	var length = vector.Length
	var values []*float64
	for i:=0; i<length; i++{
		value := *vector.Vector[i] * number
		values = append(values, &value)
	}
	return &Vector{Length: length, Vector: values},nil
}

func VectorDot(vector1,vector2 *Vector) (*float64,error) {
	if vector1.Length != vector2.Length{
		return nil,ShapeError
	}
	var valueSum float64
	for key,value1 := range vector1.Vector{
		value := *value1 * *vector2.Vector[key]
		valueSum += value
	}
	return &valueSum,nil
}

func MatrixCross(matrix1,matrix2 *Matrix)(*Matrix,error){
	if matrix1.Shape[1] != matrix2.Shape[0]{
		return nil,ShapeError
	}
	newShape := [2]int{matrix1.Shape[0],matrix2.Shape[1]}
	output := make([][]*float64, matrix1.Shape[0])
	for i:=0; i<matrix1.Shape[0]; i++ {
		newVector := make([]*float64, matrix2.Shape[1])
		for j:= 0; j< matrix2.Shape[1]; j++{
			col,_ := matrix2.GetColumn(j)
			newVector[j] = cross(matrix1.Matrix[i],col)
		}
		output[i] = newVector
	}
	return &Matrix{Shape: newShape,Matrix: output},nil
}

func MatrixMultiply(matrix *Matrix,number float64)(*Matrix,error){
	var newMatrix [][]*float64
	for i:=0;i<matrix.Shape[0];i++{
		var newVector []*float64
		for j:=0;j<matrix.Shape[1];j++{
			newValue := *matrix.Matrix[i][j] * number
			newVector = append(newVector, &newValue)
		}
		newMatrix = append(newMatrix, newVector)
	}
	return &Matrix{Shape: matrix.Shape, Matrix: newMatrix},nil
}

func MatrixDot(matrix1,matrix2 *Matrix) (*Matrix,error) {
	if matrix1.Shape[0] != matrix2.Shape[0] || matrix1.Shape[1] != matrix2.Shape[1]{
		return nil,ShapeError
	}
	var newMatrix [][]*float64
	for i:=0; i<matrix1.Shape[0]; i++{
		var newVector []*float64
		for j:=0; j<matrix1.Shape[1]; j++{
			newValue := *matrix1.Matrix[i][j] * *matrix2.Matrix[i][j]
			newVector = append(newVector, &newValue)
		}
		newMatrix = append(newMatrix, newVector)
	}
	return &Matrix{Shape: matrix1.Shape, Matrix: newMatrix}, nil
}

func cross(list1 []*float64,list2 []float64) *float64 {
	var op float64
	for i:=0;i<len(list1);i++{
		op += *list1[i]*list2[i]
	}
	return &op
}

func Sum(lis []float64) float64 {
	var op float64
	for _,val := range lis {
		op += val
	}
	return op
}

func (matrix *Matrix)Transport()*Matrix{
	newMatrix := make([][]*float64,matrix.Shape[1])
	newShape := [2]int{matrix.Shape[1],matrix.Shape[0]}
	for i:=0;i<matrix.Shape[1];i++{
		vec := make([]*float64,matrix.Shape[0])
		for j:=0;j<matrix.Shape[0];j++{
			vec[j] = matrix.Matrix[j][i]
		}
		newMatrix[i] = vec
	}
	return &Matrix{Shape:newShape, Matrix: newMatrix}
}

func (matrix *Matrix)Sum() float64 {
	var total = 0.0
	for index1:=0; index1<matrix.Shape[0]; index1++{
		for index2:=0; index2<matrix.Shape[1]; index2++ {
			total += *matrix.Matrix[index1][index2]
		}
	}
	return total
}

func (matrix *Matrix)Mean() float64 {
	count := float64(matrix.Shape[0] * matrix.Shape[1])
	return matrix.Sum() / count
}

func (matrix *Matrix)Max() float64 {
	var max = *matrix.Matrix[0][0]
	for index1:=0; index1<matrix.Shape[0]; index1++{
		for index2:=0; index2<matrix.Shape[1]; index2++ {
			if *matrix.Matrix[index1][index2] > max{
				max = *matrix.Matrix[index1][index2]
			}
		}
	}
	return max
}

func (matrix *Matrix)Min() float64 {
	var max = *matrix.Matrix[0][0]
	for index1:=0; index1<matrix.Shape[0]; index1++{
		for index2:=0; index2<matrix.Shape[1]; index2++ {
			if *matrix.Matrix[index1][index2] < max{
				max = *matrix.Matrix[index1][index2]
			}
		}
	}
	return max
}