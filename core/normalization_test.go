package core_test

import (
	"fmt"
	. "github.com/typerandom/validator/core"
	"reflect"
	"testing"
)

func testThatValueIsNormalizedToType(t *testing.T, value interface{}, expectedValue interface{}, expectedOriginalKind reflect.Kind, expectedNormalizedKind reflect.Kind, expectNil bool) {
	normalized, err := Normalize(value)

	if err != nil {
		t.Fatalf("Didn't expect any error, but got '%s'.", err)
	}

	normalizedValueKind := reflect.ValueOf(normalized.Value).Kind()

	// Test the kind of the normalized value to the expected one
	if normalizedValueKind == expectedNormalizedKind {
		if expectNil {
			if !normalized.IsNil {
				t.Fatal("Expected value to be nil, but it wasn't.")
			}
		} else {
			if normalized.IsNil {
				t.Fatal("Exepected value to not be nil, but it was")
			}
		}
		// We know the kind is right, so just do a simple string comparison
		if fmt.Sprintf("%v", normalized.Value) != fmt.Sprintf("%v", expectedValue) {
			t.Fatalf("Expected '%v' but got '%v'.", expectedValue, normalized.Value)
		}
	} else {
		t.Fatalf("Expected '%s', but got '%s'.", normalizedValueKind, expectedNormalizedKind)
	}

	if normalized.OriginalKind != expectedOriginalKind {
		t.Fatalf("Expected kind to be '%s' but got '%s'.", expectedOriginalKind, normalized.OriginalKind)
	}
}

func TestThatIntIsNormalizedToInt64(t *testing.T) {
	var value int = 123
	var nilValue *int
	testThatValueIsNormalizedToType(t, value, int64(123), reflect.Int, reflect.Int64, false)
	testThatValueIsNormalizedToType(t, &value, int64(123), reflect.Int, reflect.Int64, false)
	testThatValueIsNormalizedToType(t, nilValue, int64(0), reflect.Int, reflect.Int64, true)
}

func TestThatInt8IsNormalizedToInt64(t *testing.T) {
	var value int8 = 123
	var nilValue *int8
	testThatValueIsNormalizedToType(t, value, int64(123), reflect.Int8, reflect.Int64, false)
	testThatValueIsNormalizedToType(t, &value, int64(123), reflect.Int8, reflect.Int64, false)
	testThatValueIsNormalizedToType(t, nilValue, int64(0), reflect.Int8, reflect.Int64, true)
}

func TestThatInt16IsNormalizedToInt64(t *testing.T) {
	var value int16 = 123
	var nilValue *int16
	testThatValueIsNormalizedToType(t, value, int64(123), reflect.Int16, reflect.Int64, false)
	testThatValueIsNormalizedToType(t, &value, int64(123), reflect.Int16, reflect.Int64, false)
	testThatValueIsNormalizedToType(t, nilValue, int64(0), reflect.Int16, reflect.Int64, true)
}

func TestThatInt32IsNormalizedToInt64(t *testing.T) {
	var value int32 = 123
	var nilValue *int32
	testThatValueIsNormalizedToType(t, value, int64(123), reflect.Int32, reflect.Int64, false)
	testThatValueIsNormalizedToType(t, &value, int64(123), reflect.Int32, reflect.Int64, false)
	testThatValueIsNormalizedToType(t, nilValue, int64(0), reflect.Int32, reflect.Int64, true)
}

func TestThatInt64IsNormalizedToInt64(t *testing.T) {
	var value int64 = 123
	var nilValue *int64
	testThatValueIsNormalizedToType(t, value, value, reflect.Int64, reflect.Int64, false)
	testThatValueIsNormalizedToType(t, &value, value, reflect.Int64, reflect.Int64, false)
	testThatValueIsNormalizedToType(t, nilValue, int64(0), reflect.Int64, reflect.Int64, true)
}

func TestThatUIntIsNormalizedToInt64(t *testing.T) {
	var value uint = 123
	var nilValue *uint
	testThatValueIsNormalizedToType(t, value, int64(123), reflect.Uint, reflect.Int64, false)
	testThatValueIsNormalizedToType(t, &value, int64(123), reflect.Uint, reflect.Int64, false)
	testThatValueIsNormalizedToType(t, nilValue, int64(0), reflect.Uint, reflect.Int64, true)
}

func TestThatUInt8IsNormalizedToInt64(t *testing.T) {
	var value uint8 = 123
	var nilValue *uint8
	testThatValueIsNormalizedToType(t, value, int64(123), reflect.Uint8, reflect.Int64, false)
	testThatValueIsNormalizedToType(t, &value, int64(123), reflect.Uint8, reflect.Int64, false)
	testThatValueIsNormalizedToType(t, nilValue, int64(0), reflect.Uint8, reflect.Int64, true)
}

func TestThatUInt16IsNormalizedToInt64(t *testing.T) {
	var value uint16 = 123
	var nilValue *uint16
	testThatValueIsNormalizedToType(t, value, int64(123), reflect.Uint16, reflect.Int64, false)
	testThatValueIsNormalizedToType(t, &value, int64(123), reflect.Uint16, reflect.Int64, false)
	testThatValueIsNormalizedToType(t, nilValue, int64(0), reflect.Uint16, reflect.Int64, true)
}

func TestThatUInt32IsNormalizedToInt64(t *testing.T) {
	var value uint32 = 123
	var nilValue *uint32
	testThatValueIsNormalizedToType(t, value, int64(123), reflect.Uint32, reflect.Int64, false)
	testThatValueIsNormalizedToType(t, &value, int64(123), reflect.Uint32, reflect.Int64, false)
	testThatValueIsNormalizedToType(t, nilValue, int64(0), reflect.Uint32, reflect.Int64, true)
}

func TestThatUInt64IsNormalizedToInt64(t *testing.T) {
	var value uint64 = 123
	var nilValue *uint64
	testThatValueIsNormalizedToType(t, value, int64(123), reflect.Uint64, reflect.Int64, false)
	testThatValueIsNormalizedToType(t, &value, int64(123), reflect.Uint64, reflect.Int64, false)
	testThatValueIsNormalizedToType(t, nilValue, int64(0), reflect.Uint64, reflect.Int64, true)
}

func TestThatFloat32IsNormalizedToFloat64(t *testing.T) {
	var value float32 = 123
	var nilValue *float32
	testThatValueIsNormalizedToType(t, value, float64(123), reflect.Float32, reflect.Float64, false)
	testThatValueIsNormalizedToType(t, &value, float64(123), reflect.Float32, reflect.Float64, false)
	testThatValueIsNormalizedToType(t, nilValue, float64(0), reflect.Float32, reflect.Float64, true)
}

func TestThatFloat64IsNormalizedToFloat64(t *testing.T) {
	var value float64 = 123
	var nilValue *float64
	testThatValueIsNormalizedToType(t, value, value, reflect.Float64, reflect.Float64, false)
	testThatValueIsNormalizedToType(t, &value, value, reflect.Float64, reflect.Float64, false)
	testThatValueIsNormalizedToType(t, nilValue, float64(0), reflect.Float64, reflect.Float64, true)
}

func TestThatStringIsNormalizedToString(t *testing.T) {
	var value string = "abc"
	var nilValue *string
	testThatValueIsNormalizedToType(t, value, value, reflect.String, reflect.String, false)
	testThatValueIsNormalizedToType(t, &value, value, reflect.String, reflect.String, false)
	testThatValueIsNormalizedToType(t, nilValue, "", reflect.String, reflect.String, true)
}

func TestThatBooleanIsNormalizedToBoolean(t *testing.T) {
	var value bool = true
	var nilValue *bool
	testThatValueIsNormalizedToType(t, value, value, reflect.Bool, reflect.Bool, false)
	testThatValueIsNormalizedToType(t, &value, value, reflect.Bool, reflect.Bool, false)
	testThatValueIsNormalizedToType(t, nilValue, false, reflect.Bool, reflect.Bool, true)
}

func TestThatSliceIsNormalizedToSlice(t *testing.T) {
	var value []string = []string{"abc", "def", "123"}
	var nilValue *[]string
	testThatValueIsNormalizedToType(t, value, value, reflect.Slice, reflect.Slice, false)
	testThatValueIsNormalizedToType(t, &value, value, reflect.Slice, reflect.Slice, false)
	testThatValueIsNormalizedToType(t, nilValue, []string{}, reflect.Slice, reflect.Slice, true)
}

func TestThatMapIsNormalizedToMap(t *testing.T) {
	var value map[string]string = map[string]string{"abc": "def"}
	var nilValue *map[string]string
	testThatValueIsNormalizedToType(t, value, value, reflect.Map, reflect.Map, false)
	testThatValueIsNormalizedToType(t, &value, value, reflect.Map, reflect.Map, false)
	testThatValueIsNormalizedToType(t, nilValue, map[string]string{}, reflect.Map, reflect.Map, true)
}

func TestThatDeepPointerValuesCanBeNormalized(t *testing.T) {
	var value uint32 = 123

	ptrA := &value
	ptrB := &ptrA
	ptrC := &ptrB

	var ptrValue ****uint32 = &ptrC

	testThatValueIsNormalizedToType(t, ptrValue, int64(123), reflect.Uint32, reflect.Int64, false)
}

func TestThatDeepPointerNilValuesCanBeNormalized(t *testing.T) {
	var ptrValue *****uint32
	testThatValueIsNormalizedToType(t, ptrValue, int64(0), reflect.Uint32, reflect.Int64, true)
}

func TestThatInvalidValuesCanBeNormalized(t *testing.T) {
	testThatValueIsNormalizedToType(t, nil, nil, reflect.Invalid, reflect.Invalid, true)
}
