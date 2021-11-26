package pkg

func ConvertToInt64(val interface{}) int64 {
	if v, ok := val.(int64); ok {
		return v
	}
	if v, ok := val.(int); ok {
		return int64(v)
	}
	panic("invalid int64 value")
}

func ConvertToString(val interface{}) string {
	if v, ok := val.(string); ok {
		return v
	}
	if v, ok := val.([]byte); ok {
		return string(v)
	}
	panic("invalid string value")
}
