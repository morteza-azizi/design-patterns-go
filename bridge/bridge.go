package bridge

import "fmt"

type DrawingAPI interface {
	DrawCircle(x, y, radius float64)
	DrawSquare(x, y, side float64)
}

type WindowsAPI struct{}

func (w *WindowsAPI) DrawCircle(x, y, radius float64) {
	fmt.Printf("Drawing a circle on Windows at (%.2f, %.2f) with radius %.2f\n", x, y, radius)
}

func (w *WindowsAPI) DrawSquare(x, y, side float64) {
	fmt.Printf("Drawing a square on Windows at (%.2f, %.2f) with side %.2f\n", x, y, side)
}

type LinuxAPI struct{}

func (l *LinuxAPI) DrawCircle(x, y, radius float64) {
	fmt.Printf("Drawing a circle on Linux at (%.2f, %.2f) with radius %.2f\n", x, y, radius)
}

func (l *LinuxAPI) DrawSquare(x, y, side float64) {
	fmt.Printf("Drawing a square on Linux at (%.2f, %.2f) with side %.2f\n", x, y, side)
}

type Shape interface {
	Draw()
}

type Circle struct {
	x, y, radius float64
	api          DrawingAPI
}

func (c *Circle) Draw() {
	c.api.DrawCircle(c.x, c.y, c.radius)
}

type Square struct {
	x, y, side float64
	api        DrawingAPI
}

func (s *Square) Draw() {
	s.api.DrawSquare(s.x, s.y, s.side)
}
