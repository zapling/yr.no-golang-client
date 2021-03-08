package utils

func StringValue(s *string) string {
	if s == nil {
		return ""
	}

	return *s
}

func Float64Value(f *float64) float64 {
	if f == nil {
		return 0
	}

	return *f
}
