package main

// If integer, then tries to resolve the value to either a int64 value or pointer.
// If it fails to do that, then the original value is returned.
// Note: This could have been done with reflection. But this is probably clearer and better performance wise.
// ALSO: Note that floats will be truncated according to http://golang.org/ref/spec#Conversions (float -> int)
func tryResolveInt64(value interface{}) interface{} {
	switch typedValue := value.(type) {
	case int:
		return int64(typedValue)
	case *int:
		int64Val := int64(*typedValue)
		return &int64Val
	case int8:
		return int64(typedValue)
	case *int8:
		int64Val := int64(*typedValue)
		return &int64Val
	case int16:
		return int64(typedValue)
	case *int16:
		int64Val := int64(*typedValue)
		return &int64Val
	case int32:
		return int64(typedValue)
	case *int32:
		int64Val := int64(*typedValue)
		return &int64Val
	case uint:
		return int64(typedValue)
	case *uint:
		int64Val := int64(*typedValue)
		return &int64Val
	case uint8:
		return int64(typedValue)
	case *uint8:
		int64Val := int64(*typedValue)
		return &int64Val
	case uint16:
		return int64(typedValue)
	case *uint16:
		int64Val := int64(*typedValue)
		return &int64Val
	case uint32:
		return int64(typedValue)
	case *uint32:
		int64Val := int64(*typedValue)
		return &int64Val
	case uint64:
		return int64(typedValue)
	case *uint64:
		int64Val := int64(*typedValue)
		return &int64Val
	// Note that converting a float to int will discard any fraction.
	// It's important that this effect is documented.
	case float32:
		return int64(typedValue)
	case *float32:
		int64Val := int64(*typedValue)
		return &int64Val
	case float64:
		return int64(typedValue)
	case *float64:
		int64Val := int64(*typedValue)
		return &int64Val
	}
	return value
}
