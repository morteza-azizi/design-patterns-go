package bridge

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockPrinter struct {
	mock.Mock
}

func (mp *MockPrinter) Printf(format string, a ...interface{}) {
	mp.Called(format, a)
}

func TestCircle_WindowsAPIDraw(t *testing.T) {
	mockPrinter := new(MockPrinter)

	circle := Circle{
		x:      1.0,
		y:      2.0,
		radius: 3.0,
		api:    &WindowsAPI{},
	}

	mockPrinter.On("Printf", "Drawing a circle on Windows at (%.2f, %.2f) with radius %.2f\n", []interface{}{circle.x, circle.y, circle.radius})

	circle.Draw(mockPrinter)

	mockPrinter.AssertCalled(t, "Printf", "Drawing a circle on Windows at (%.2f, %.2f) with radius %.2f\n", []interface{}{circle.x, circle.y, circle.radius})
}

func TestCircle_LinuxAPIDraw(t *testing.T) {
	mockPrinter := new(MockPrinter)

	circle := Circle{
		x:      1.0,
		y:      2.0,
		radius: 3.0,
		api:    &LinuxAPI{},
	}

	mockPrinter.On("Printf", "Drawing a circle on Linux at (%.2f, %.2f) with radius %.2f\n", []interface{}{circle.x, circle.y, circle.radius})

	circle.Draw(mockPrinter)

	mockPrinter.AssertCalled(t, "Printf", "Drawing a circle on Linux at (%.2f, %.2f) with radius %.2f\n", []interface{}{circle.x, circle.y, circle.radius})
}

func TestSquare_DrawWindowsAPI(t *testing.T) {
	mockPrinter := new(MockPrinter)

	square := Square{
		x:    1.0,
		y:    2.0,
		side: 3.0,
		api:  &WindowsAPI{},
	}

	mockPrinter.On("Printf", "Drawing a square on Windows at (%.2f, %.2f) with side %.2f\n", []interface{}{square.x, square.y, square.side})

	square.Draw(mockPrinter)

	mockPrinter.AssertCalled(t, "Printf", "Drawing a square on Windows at (%.2f, %.2f) with side %.2f\n", []interface{}{square.x, square.y, square.side})
}

func TestSquare_DrawLinuxAPI(t *testing.T) {
	mockPrinter := new(MockPrinter)

	square := Square{
		x:    1.0,
		y:    2.0,
		side: 3.0,
		api:  &LinuxAPI{},
	}

	mockPrinter.On("Printf", "Drawing a square on Linux at (%.2f, %.2f) with side %.2f\n", []interface{}{square.x, square.y, square.side})

	square.Draw(mockPrinter)

	mockPrinter.AssertCalled(t, "Printf", "Drawing a square on Linux at (%.2f, %.2f) with side %.2f\n", []interface{}{square.x, square.y, square.side})
}
