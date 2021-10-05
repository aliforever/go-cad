package tests

import (
	"fmt"
	"testing"

	"github.com/aliforever/go-cad"
)

func TestGoString(t *testing.T) {
	c := cad.Cents(224)

	expected := "cad.Cents(224)"
	result := fmt.Sprintf("%#v", c)

	if expected != result {
		t.Fatal(fmt.Sprintf("Test Failed.\nExpected: %s\nGot: %s", expected, result))
	}
}

func TestParseCadGoString(t *testing.T) {
	c, err := cad.ParseCAD("-9¢")
	if err != nil {
		t.Fatal("Test Failed.", err)
		return
	}

	expected := "cad.Cents(-9)"
	result := fmt.Sprintf("%#v", c)

	if expected != result {
		t.Fatal(fmt.Sprintf("Test Failed.\nExpected: %s\nGot: %s", expected, result))
	}
}
