package cad

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

type CAD struct {
	cents int64
}

// Cents returns a CAD that represents ‘n’ cents.
//
// For example, if one was to call:
//
//	cad := Cents(105)
//
// Then ‘cad’ would be: $1.05
func abs(n int64) int64 {
	if n < 0 {
		return -1 * n
	}
	return n
}

func Cents(n int64) CAD {
	return CAD{
		cents: n,
	}
}

// ParseCAD parses the string ‘s’ and return the equivalent CAD.
//
// If ‘s’ does not contain a money amount, then ParseCAD returns an error.
//
// Some example valid strings include:
//
// • -$1234.56
// • $-1234.56
// • -$1,234.56
// • $-1,234.56
// • CAD -$1234.56
// • CAD $-1234.56
// • CAD-$1,234.56
// • CAD$-1,234.56
// • $1234.56
// • $1,234.56
// • CAD $1234.56
// • CAD $1,234.56
// • CAD$1234.56
// • CAD$1,234.56
// • $0.09
// • $.09
// • -$0.09
// • -$.09
// • $-0.09
// • $-.09
// • CAD $0.09
// • CAD $.09
// • CAD -$0.09
// • CAD -$.09
// • CAD $-0.09
// • CAD $-.09
// • CAD$0.09
// • CAD$.09
// • CAD-$0.09
// • CAD-$.09
// • CAD$-0.09
// • CAD$-.09
// • 9¢
// • -9¢
// • 123456¢
// • -123456¢
func ParseCAD(s string) (cad CAD, err error) {
	possibleErr := errors.New("invalid_cad")

	s = strings.TrimSpace(s)

	cadStr := "CAD"
	cadIndex := strings.Index(s, cadStr)

	if cadIndex > 0 {
		err = possibleErr
		return
	}

	s = strings.Replace(s, "CAD", "", 1)
	s = strings.TrimSpace(s)

	r := regexp.MustCompile(`[^$\d¢,.-]`)
	if r.MatchString(s) {
		err = possibleErr
		return
	}

	const (
		dollarSign = "$"
		centSign   = "¢"
		dot        = "."
		minus      = "-"
	)

	if strings.Count(s, dollarSign) > 1 || strings.Count(s, centSign) > 1 || strings.Count(s, dot) > 1 || strings.Count(s, minus) > 1 || strings.Count(s, cadStr) > 1 {
		err = possibleErr
		return
	}

	dollarSignIndex := strings.Index(s, dollarSign)
	centSignIndex := strings.Index(s, centSign)
	minusIndex := strings.Index(s, minus)
	dotIndex := strings.Index(s, dot)

	if dollarSignIndex == -1 && centSignIndex == -1 {
		err = possibleErr
		return
	}

	if dollarSignIndex != -1 && centSignIndex != -1 {
		err = possibleErr
		return
	}

	if dollarSignIndex != -1 && dollarSignIndex > 1 {
		err = possibleErr
		return
	}

	if centSignIndex != -1 && centSignIndex != utf8.RuneCountInString(s)-1 {
		err = possibleErr
		return
	}

	hasMinus := false

	if minusIndex > 1 {
		err = possibleErr
		return
	} else if minusIndex != -1 {
		hasMinus = true
	}

	s = strings.ReplaceAll(s, dollarSign, "")
	s = strings.ReplaceAll(s, centSign, "")
	s = strings.ReplaceAll(s, minus, "")
	s = strings.ReplaceAll(s, ",", "")
	split := strings.Split(s, ".")
	if len(split) == 2 {
		if number, numberErr := strconv.Atoi(split[1]); numberErr != nil || number > 99 {
			err = possibleErr
			return
		}
	}
	s = strings.ReplaceAll(s, ".", "")

	var number int
	number, err = strconv.Atoi(s)
	if err != nil {
		err = possibleErr
		return
	}

	if dotIndex == -1 && dollarSignIndex != -1 {
		number *= 100
	}

	cad = Cents(negativeOnFlag(hasMinus, int64(number)))
	return
}

//
// Abs returns the absolute value.
func (c CAD) Abs() CAD {
	return CAD{
		cents: abs(c.cents),
	}
}

// AsCents returns CAD as the number of pennies it is equivalent to.
func (c CAD) AsCents() int64 {
	return c.cents
}

// Add adds two CAD and returns the result.
func (c CAD) Add(other CAD) CAD {
	// return Cents((c.whole+other.whole)*100 + (c.cents + other.cents))
	return Cents(c.AsCents() + other.AsCents())
}

func negativeOnFlag(f bool, n int64) (result int64) {
	result = n
	if f {
		result = -n
	}
	return
}

// CanonicalForm returns the number of dollars and cents that CAD represents.
//
// ‘cents’ is always less than for equal to 99. I.e.,:
//	cents ≤ 99
func (c CAD) CanonicalForm() (dollars int64, cents int64) {
	return c.cents / 100, c.cents % 100
}

// Mul multiplies CAD by a scalar (number) and returns the result.
func (c CAD) Mul(scalar int64) CAD {
	return Cents(c.AsCents() * scalar)
}

// Sub subtracts two CAD and returns the result.
func (c CAD) Sub(other CAD) CAD {
	return Cents(c.AsCents() - other.AsCents())
}

func (c CAD) GoString() string {
	return fmt.Sprintf("cad.Cents(%d)", c.cents)
}

func (c CAD) String() string {
	var sign string = ""
	{
		if c.cents < 0 {
			sign = "-"
		}
	}

	var whole int64
	var fractional int64
	{
		whole, fractional = c.CanonicalForm()

		whole = abs(whole)
		fractional = abs(fractional)
	}

	return fmt.Sprintf("CAD$%s%d.%02d", sign, whole, fractional)
}

func (c CAD) MarshalJSON() (b []byte, err error) {
	return json.Marshal(c.String())
}

func (c *CAD) UnmarshalJSON(b []byte) (err error) {
	if c == nil {
		err = errors.New("nil_receiver")
		return
	}
	var cadStr string
	if err = json.Unmarshal(b, &cadStr); err != nil {
		return
	}
	var value CAD
	value, err = ParseCAD(cadStr)
	if err == nil {
		*c = value
	}
	return
}

func (c CAD) Value() (driver.Value, error) {
	data := c.String()
	if strings.HasPrefix(data, "CAD") {
		data = data[len("CAD"):]
	}
	return data, nil
}

func (c *CAD) Scan(value interface{}) (err error) {
	if c == nil {
		return errors.New("nil_receiver")
	}
	var bv driver.Value
	bv, err = driver.String.ConvertValue(value)
	if err != nil {
		return
	}

	var s string
	switch casted := bv.(type) {
	case []byte:
		s = string(casted)
	case string:
		s = casted
	default:
		return errors.New("internal_error")
	}

	var val CAD
	val, err = ParseCAD(s)
	if err != nil {
		return
	}
	*c = val
	return
}
