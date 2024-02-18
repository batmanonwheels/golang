package perimeter

import "testing"

func TestPerimeter(t *testing.T) {

	checkPerimeter := func(t testing.TB, got, want float64) {
		if got != want {
			t.Helper()
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}
	t.Run("add all sides up", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := Perimeter(rectangle)
		want := 40.0
		checkPerimeter(t, got, want)

	})
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{shape: Circle{Radius: 10}, want: 314.1592653589793},
		{shape: Triangle{Base: 12, Height: 6}, want: 36},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.want)
		}
	}

}
