package matrix

import (
	"errors"
)

var IndexError = errors.New("The index out of shape! ")
var ShapeError = errors.New("Inputs has error shape! ")