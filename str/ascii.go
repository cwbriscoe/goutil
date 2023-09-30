package str

// ToASCII will return a string that is stripped of non ASCII readable characters.
func ToASCII(s string) string {
	rs := make([]rune, 0, len(s))
	for _, r := range s {
		if IsReadable(r) {
			rs = append(rs, r)
		}
	}
	return string(rs)
}

// IsLower returns true if the rune is an ASCII lower case character [a-z].
func IsLower(r rune) bool {
	return r >= 97 && r <= 122
}

// IsUpper returns true if the rune is an ASCII upper case character [A-Z].
func IsUpper(r rune) bool {
	return r >= 65 && r <= 90
}

// IsDigit returns true if the rune is an ASCII numerical digit [0-9].
func IsDigit(r rune) bool {
	return r >= 48 && r <= 57
}

// IsSpace returns true if the rune is an ASCII space [ ].
func IsSpace(r rune) bool {
	return r == 32
}

// IsQuote returns true if the rune is an ASCII quote character ["'`].
func IsQuote(r rune) bool {
	return r == 34 || r == 39 || r == 96
}

// IsBackslash returns true if the rune is an ASCII backslash [\].
func IsBackslash(r rune) bool {
	return r == 92
}

// IsReadable returns true if the rune is an ASCII readable character.
func IsReadable(r rune) bool {
	return r >= 32 && r <= 126
}

// IsSpecial returns true if the rune is an ASCII special character and not one of the 3 quote characters or backslash.
// [!#$%&()*+,-./:;<=>?@^_{|}~] and not ["'`\]
func IsSpecial(r rune) bool {
	return r == 33 || (r >= 35 && r <= 38) || (r >= 40 && r <= 47) || (r >= 58 && r <= 64) || r == 91 || (r >= 93 && r <= 95) || (r >= 123 && r <= 126)
}
