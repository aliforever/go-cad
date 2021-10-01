package tests

import (
	"fmt"
	"testing"

	"github.com/aliforever/go-cad"
)

func TestAddCents(t *testing.T) {
	x := int64(105)
	y := int64(106)

	c := cad.Cents(x)
	d := cad.Cents(y)

	result := c.Add(d).AsCents()
	expected := x + y

	if result != expected {
		t.Fatal(fmt.Sprintf("Test Failed. Expected Cents: %d\nCents: %d", expected, result))
	}
}

func TestAddParse(t *testing.T) {
	x := "$1.05"
	y := "$1.06"

	xCad, err := cad.ParseCAD(x)
	if err != nil {
		t.Fatal(err)
		return
	}

	yCad, err := cad.ParseCAD(y)
	if err != nil {
		t.Fatal(err)
		return
	}

	result := xCad.Add(yCad).AsCents()
	expected := int64(211)

	if result != expected {
		t.Fatal(fmt.Sprintf("Test Failed. Expected Cents: %d\nCents: %d", expected, result))
	}
}
