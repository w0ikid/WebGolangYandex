// package openclosed

// import "math"

// type Shapes interface {
// 	// some data
// }

// type Rectangle struct {
// 	Length int
// 	Breadth int
// }

// type Circle struct {
// 	Radius int
// }

// func Area(s []Shapes) float64 {
// 	var totalArea float64
// 	for _, shape := range s {
// 		switch shape.(type) {
// 			case Circle:
// 				totalArea += math.Pi * math.Pow(float64(shape.(Circle).Radius), 2)
// 			case Rectangle:
// 				totalArea += float64(shape.(Rectangle).Length * shape.(Rectangle).Breadth)
// 		}
// 	}
// 	return totalArea
// }


// // Это пример нарушения принципа открытости/закрытости.
// // Если мы хотим добавить новую фигуру, то нам придется изменить функцию Area.


// // Чтобы исправить это, мы можем использовать интерфейс.
// // Теперь, если мы хотим добавить новую фигуру, нам нужно просто реализовать интерфейс Shapes.

// --------------------------------------------------


package openclosed

import "math"

type Shapes interface {
	Area() float64
}

type Rectangle struct {
	Length int
	Breadth int
}

func (r Rectangle) Area() float64 {
	return float64(r.Breadth) * float64(r.Length)
}

type Circle struct {
	Radius int
}

func (c Circle) Area() float64 {
	return math.Pi * float64(c.Radius)
}

func TotalArea(shapes []Shapes) float64 {
	var totalArea float64

	for _, shape := range shapes {
		totalArea += shape.Area()
	}

	return totalArea
}