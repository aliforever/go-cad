package tests

import (
	"fmt"
	"testing"

	"github.com/aliforever/go-cad"
)

func TestMulCents(t *testing.T) {
	x := int64(2)
	n := int64(106)

	c := cad.Cents(n)

	result := c.Mul(x).AsCents()
	expected := x * n

	if result != expected {
		t.Fatal(fmt.Sprintf("Test Failed. Expected Cents: %d\nCents: %d", expected, result))
	}
}

func TestMulParse(t *testing.T) {
	n := "$1.07"
	x := int64(2)

	xCad, err := cad.ParseCAD(n)
	if err != nil {
		t.Fatal(err)
		return
	}

	result := xCad.Mul(x).AsCents()
	expected := int64(214)

	if result != expected {
		t.Fatal(fmt.Sprintf("Test Failed. Expected Cents: %d\nCents: %d", expected, result))
	}
}
