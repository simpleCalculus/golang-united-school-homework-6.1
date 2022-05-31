package golang_united_school_homework

import "errors"

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
		return errors.New("out of the shapesCapacity range")
	}
	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if i < 0 || i >= b.shapesCapacity {
		return nil, errors.New("index went out of the range")
	}
	if i >= len(b.shapes) {
		return nil, errors.New("index doesn't exist")
	}
	return b.shapes[i], nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	shape, err := b.GetByIndex(i)
	if err == nil {
		b.shapes = append(b.shapes[0:i], b.shapes[i+1:]...)
	}
	return shape, err
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	sh, err := b.GetByIndex(i)
	if err == nil {
		b.shapes[i] = shape
	}
	return sh, err
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	sum := float64(0)
	for _, v := range b.shapes {
		sum += v.CalcPerimeter()
	}
	return sum
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	sum := float64(0)
	for _, v := range b.shapes {
		sum += v.CalcArea()
	}
	return sum
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	var existCircle bool
	for i := 0; i < len(b.shapes); i++ {
		_, ok := b.shapes[i].(Circle)
		_, ok2 := b.shapes[i].(*Circle)
		if ok || ok2 {
			b.ExtractByIndex(i)
			i--
			existCircle = true
		}
	}
	if existCircle {
		return nil
	}
	return errors.New("circles are not exist")
}
