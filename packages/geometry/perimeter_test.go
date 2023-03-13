package geometry

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{6.0, 6.0}
	got := Perimeter(rectangle)
	want := 24.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

// Using TableDrivenTests https://github.com/golang/go/wiki/TableDrivenTests
// Good fit when you wish to test various implementations of an interface,
// or if the data being passed in to a function has lots of different requirements that need testing.
func TestArea(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{12, 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.hasArea {
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
		}
	}
}
