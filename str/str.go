package str

import (
	"reflect"
	"sort"
	"unsafe"
)

// TrimQuotes removes quotes from the beginning and end of a string.
func TrimQuotes(s string) string {
	if len(s) >= 2 {
		if s[0] == '"' && s[len(s)-1] == '"' {
			return s[1 : len(s)-1]
		}
	}
	return s
}

// SortAndRemoveDuplicates will sort the string slice while removing duplicates.
func SortAndRemoveDuplicates(s []string) []string {
	if len(s) < 1 {
		return s
	}

	sort.Strings(s)
	prev := 1
	for curr := 1; curr < len(s); curr++ {
		if s[curr-1] != s[curr] {
			s[prev] = s[curr]
			prev++
		}
	}

	return s[:prev]
}

// UnsafeByteToString converts a byte to a string without memory allocations.
// They will both share the same bytes, so make sure not to modify the bytes
// while the string still survives.
func UnsafeByteToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// UnsafeStringToByte converts a string to a byte without memory allocations.
// They will both share the same bytes, so make sure not to modify the bytes
// while the string still survives.
func UnsafeStringToByte(s string) (b []byte) {
	sHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	bHeader.Data = sHeader.Data
	bHeader.Cap = sHeader.Len
	bHeader.Len = sHeader.Len
	return b
}
