package bridge

type Printer interface {
	Printf(format string, a ...interface{})
}

type DrawingAPI interface {
	DrawCircle(x, y, radius float64, printer Printer)
	DrawSquare(x, y, side float64, printer Printer)
}

type WindowsAPI struct{}

func (w *WindowsAPI) DrawCircle(x, y, radius float64, printer Printer) {
	printer.Printf("Drawing a circle on Windows at (%.2f, %.2f) with radius %.2f\n", x, y, radius)
}

func (w *WindowsAPI) DrawSquare(x, y, side float64, printer Printer) {
	printer.Printf("Drawing a square on Windows at (%.2f, %.2f) with side %.2f\n", x, y, side)
}

type LinuxAPI struct{}

func (l *LinuxAPI) DrawCircle(x, y, radius float64, printer Printer) {
	printer.Printf("Drawing a circle on Linux at (%.2f, %.2f) with radius %.2f\n", x, y, radius)
}

func (l *LinuxAPI) DrawSquare(x, y, side float64, printer Printer) {
	printer.Printf("Drawing a square on Linux at (%.2f, %.2f) with side %.2f\n", x, y, side)
}

type Shape interface {
	Draw()
}

type Circle struct {
	x, y, radius float64
	api          DrawingAPI
}

func (c *Circle) Draw(printer Printer) {
	c.api.DrawCircle(c.x, c.y, c.radius, printer)
}

type Square struct {
	x, y, side float64
	api        DrawingAPI
}

func (s *Square) Draw(printer Printer) {
	s.api.DrawSquare(s.x, s.y, s.side, printer)
}
