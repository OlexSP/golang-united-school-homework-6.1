package golang_united_school_homework

import (
	"errors"
	"fmt"
)

var (
	errorOutOfCapacity = errors.New("out of the shapesCapacity range")
	errorOutOfRange    = errors.New("index is out of range")
	errorNotExist      = errors.New("shape not exist in the list")
	inputErrorString   = "input error: %w" // for quality test
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) == b.shapesCapacity {
		return fmt.Errorf(inputErrorString, errorOutOfCapacity)
	}
	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if len(b.shapes) <= i {
		return nil, fmt.Errorf(inputErrorString, errorOutOfRange)
	}
	return b.shapes[i], nil

}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if len(b.shapes) <= i {
		return nil, fmt.Errorf(inputErrorString, errorOutOfRange)
	}
	RemovedShape := b.shapes[i]
	b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
	return RemovedShape, nil

}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if len(b.shapes) <= i {
		return nil, fmt.Errorf(inputErrorString, errorOutOfRange)
	}
	RemovedShape := b.shapes[i]
	b.shapes[i] = shape
	return RemovedShape, nil

}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var sum float64
	for _, v := range b.shapes {
		sum += v.CalcPerimeter()
	}
	return sum
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var sum float64
	for _, v := range b.shapes {
		sum += v.CalcArea()
	}
	return sum
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	var counterSlice []int
	for i, v := range b.shapes {
		switch v.(type) {
		case Circle:
			counterSlice = append(counterSlice, i)
		}
	}
	if len(counterSlice) == 0 {
		return fmt.Errorf(inputErrorString, errorNotExist)
	}
	for _, v := range counterSlice {
		b.shapes = append(b.shapes[:v], b.shapes[v+1:]...)
	}
	return nil
}
