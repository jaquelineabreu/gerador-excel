package internal

import "time"

func convertToStringInterface(slice []string) []interface{} {
	interfaceSlice := make([]interface{}, len(slice))
	for i, d := range slice {
		interfaceSlice[i] = d
	}
	return interfaceSlice
}

func treatsString(ptr *string) string {
	if ptr == nil {
		return ""
	}
	return *ptr
}

func treatsInt64(ptr *int64) int64 {
	if ptr == nil {
		return 0
	}
	return *ptr
}

func treatsFloat64(ptr *float64) float64 {
	if ptr == nil {
		return 0.0
	}
	return *ptr
}

func treatsTime(ptr *time.Time) time.Time {
	if ptr == nil {
		return time.Time{}
	}
	return *ptr
}
