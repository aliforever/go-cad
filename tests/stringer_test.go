package tests

import (
	"fmt"
	"testing"

	"github.com/aliforever/go-cad"
)

func TestStringer(t *testing.T) {
	c := cad.Cents(204)

	expected := "CAD$2.04"
	result := fmt.Sprintf("%s", c)

	if expected != result {
		t.Fatal(fmt.Sprintf("Test Failed.\nExpected: %s\nGot: %s", expected, result))
	}
}

func TestParseCadStringer(t *testing.T) {
	c, err := cad.ParseCAD("-9¢")
	if err != nil {
		t.Fatal("Test Failed.", err)
		return
	}

	expected := "CAD$0.09"
	result := fmt.Sprintf("%s", c)

	if expected != result {
		t.Fatal(fmt.Sprintf("Test Failed.\nExpected: %s\nGot: %s", expected, result))
	}
}
