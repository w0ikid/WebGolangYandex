package main

import "fmt"

type Rectangle struct {
	width, height int
}

func (r Rectangle) SetWidth(width int) {
	r.width = width
}

func (r Rectangle) SetHeight(height int) {
	r.height = height
}

func (r Rectangle) Area() int {
	return r.width * r.height
}

type Square struct {
	Rectangle
}

func (s Square) SetWidth(width int) {
	s.width = width
	s.height = width
}

func (s Square) SetHeight(height int) {
	s.width = height
	s.height = height
}

func ResizeRectangle(r *Rectangle) {
	r.SetWidth(5)
	r.SetHeight(10)

	fmt.Println("expected area", 50)
	fmt.Println("actual area", r.Area())
}

func main() {
	rect := &Rectangle{}

	ResizeRectangle(rect)

	Square := &Square{}
	ResizeRectangle(&Square.Rectangle)


}
