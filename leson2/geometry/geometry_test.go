package geometry

import (
	"testing"
)

// Test calc diameter of circle
func TestCircleArea(t *testing.T) {
	var r float64 = 10
	expA := 314.15926535897932581065

	c := Circle{r}
	if c.radius == 0 {
		t.Error("Expected struct Circle, got empty struct")
	}

	if a := c.Area(); a != expA {
		t.Errorf("Expected area of circle %5.20f, got %5.20f", expA, a)
	}
}


func TestCircleDiameter(t *testing.T) {
	var r float64 = 10
	expD := 20.0

	c := Circle{r}
	if c.radius == 0 {
		t.Error("Expected struct Circle, got empty struct")
	}

	if d := c.Diameter(); d != expD {
		t.Errorf("Expected diameter of circle %5.20f, got %5.20f", expD, d)
	}
}

func TestCircleLength(t *testing.T) {
	var r float64 = 10
	expL := 62.83185307179586231996

	c := Circle{r}
	if c.radius == 0 {
		t.Error("Expected struct Circle, got empty struct")
	}

	if l := c.Length(); l != expL {
		t.Errorf("Expected length of citcle %5.20f, got %5.20f", expL, l)
	}
}


func TestRectangleArea(t *testing.T) {
	var w float64 = 10
	var h float64 = 5
	expA := 50.0

	r := Rectangle{w, h}
	if r.w == 0 || r.h == 0 {
		t.Error("Expected struct Rectangle, got empty struct")
	}

	if a := r.Area(); a != expA {
		t.Errorf("Expected area of rectangle %5.20f, got %5.20f", expA, a)
	}
}

