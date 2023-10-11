package num

import "strconv"

// ParseInt32 is a convenience function to parse a string into an int32
func ParseInt32(num string) (int32, error) {
	val, err := strconv.ParseInt(num, 10, 32)
	return int32(val), err
}

// ParseUint32 is a convenience function to parse a string into an uint32
func ParseUint32(num string) (uint32, error) {
	val, err := strconv.ParseUint(num, 10, 32)
	return uint32(val), err
}

// ParseInt64 is a convenience function to parse a string into an int64
func ParseInt64(num string) (int64, error) {
	return strconv.ParseInt(num, 10, 64)
}

// ParseUint64 is a convenience function to parse a string into an uint64
func ParseUint64(num string) (uint64, error) {
	return strconv.ParseUint(num, 10, 64)
}
