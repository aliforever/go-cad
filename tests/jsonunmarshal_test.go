package tests

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/aliforever/go-cad"
)

func TestJsonUnmarshal(t *testing.T) {
	type Data struct {
		Name    string  `json:"name"`
		Balance cad.CAD `json:"balance"`
	}

	jsonBytes := []byte(`{"name":"Jow Blow","balance":"CAD$123.45"}`)
	var d *Data
	err := json.Unmarshal(jsonBytes, &d)
	if err != nil {
		t.Fatal("Test Failed Unmarshalling.", err)
		return
	}

	expectedBalance := int64(12345)
	result := d.Balance.AsCents()

	if expectedBalance != result {
		t.Fatal(fmt.Sprintf("Test Failed.\nExpected: %d\nGot: %d", expectedBalance, result))
	}
}
