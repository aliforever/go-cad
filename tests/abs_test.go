package tests

import (
	"fmt"
	"math"
	"testing"

	"github.com/aliforever/go-cad"
)

func TestAbs(t *testing.T) {
	x := int64(-105)

	c := cad.Cents(x)

	result := c.Abs().AsCents()
	expected := int64(math.Abs(float64(x)))

	if result != expected {
		t.Fatal(fmt.Sprintf("Test Failed. Expected Cents: %d\nCents: %d", expected, result))
	}
}

func TestAbsParse(t *testing.T) {
	x := "-$1.05"

	xCad, err := cad.ParseCAD(x)
	if err != nil {
		t.Fatal(err)
		return
	}

	result := xCad.Abs().AsCents()
	expected := int64(105)

	if result != expected {
		t.Fatal(fmt.Sprintf("Test Failed. Expected Cents: %d\nCents: %d", expected, result))
	}
}
