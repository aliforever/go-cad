package cad

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
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

	r := regexp.MustCompile(`[^$\dCAD¢.-]`)
	if r.MatchString(s) {
		err = possibleErr
		return
	}

	dollarSign := "$"
	centSign := "¢"

	dollarSignIndex := strings.Index(s, dollarSign)
	centSignIndex := strings.Index(s, centSign)

	if dollarSignIndex == -1 && centSignIndex == -1 {
		err = errors.New("$_or_¢_not_defined")
		return
	}

	if dollarSignIndex != -1 && centSignIndex != -1 {
		err = errors.New("should_not_contain_dollar_and_cent_signs_together")
		return
	}

	if dollarSignIndex > 1 {
		err = errors.New("invalid_position_for_dollar_sing")
		return
	}

	if strings.Count(s, dollarSign) > 1 || strings.Count(s, centSign) > 1 || strings.Contains(s, "-") {
		err = errors.New("more_than_one_dollar_or_cent_sign")
		return
	}

	s = strings.ReplaceAll(s, "$", "")
	s = strings.ReplaceAll(s, "CAD", "")
	s = strings.ReplaceAll(s, "¢", "")
	s = strings.ReplaceAll(s, ",", "")
	s = strings.TrimSpace(s)

	isNegative := false
	negativeIndex := strings.Index(s, "-")
	if negativeIndex != -1 && negativeIndex != 0 {
		err = possibleErr
		return
	} else if negativeIndex == 0 {
		isNegative = true
		s = strings.ReplaceAll(s, "-", "")
	}

	if s == "" {
		err = possibleErr
		return
	}

	var whole, decimal int

	split := strings.Split(s, ".")
	if len(split) > 2 {
		err = possibleErr
		return
	}

	if split[0] == "" {
		split[0] = "0"
	}

	if len(split) == 2 {
		if split[1] == "" {
			split[1] = "0"
		}
	}

	var parseErr error

	whole, parseErr = strconv.Atoi(split[0])
	if parseErr != nil {
		err = possibleErr
		return
	}

	if len(split) == 2 {
		if centSignIndex != -1 {
			err = possibleErr
			return
		}
		decimal, parseErr = strconv.Atoi(split[1])
		if parseErr != nil {
			err = possibleErr
			return
		}
		if decimal > 99 {
			err = possibleErr
			return
		}
	}

	if centSignIndex != -1 && whole > 0 && decimal > 0 {
		err = possibleErr
		return
	}

	var cents int64
	if centSignIndex != -1 && decimal == 0 {
		cents = negativeOnFlag(isNegative, int64(whole))
	} else {
		cents = negativeOnFlag(isNegative, int64((whole*100)+decimal))
	}

	cad = Cents(cents)
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
	whole := c.cents / 100
	return whole, c.cents - whole
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
	dollars, cents := c.CanonicalForm()
	return fmt.Sprintf("CAD$%d.%02d", dollars, cents)
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
	return c.String(), nil
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
