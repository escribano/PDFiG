package pdf

const (
	regularCharacter = iota
	whiteSpaceCharacter
	delimiterCharacter
)

var (
	whiteSpaceCharacters = "\000\t\n\f\r "
	delimiterCharacters  = "%()/<>[]{}"
	characterTypes       [256]byte
)

func init() {
	for i := range characterTypes {
		characterTypes[i] = regularCharacter
	}

	for _, w := range whiteSpaceCharacters {
		characterTypes[w] = whiteSpaceCharacter
	}

	for _, d := range delimiterCharacters {
		characterTypes[d] = delimiterCharacter
	}
}

// Is the passed byte PDF white space?
func IsWhiteSpace(b byte) bool {
	return characterTypes[b] == whiteSpaceCharacter
}

// Is the passed byte a PDF delimiter?
func IsDelimiter(b byte) bool {
	return characterTypes[b] == delimiterCharacter
}

// Is the passed byte a regular PDF character?
func IsRegular(b byte) bool {
	return characterTypes[b] == regularCharacter
}

func IsAlpha(b byte) bool {
	return b>='a' && b<='z' ||
		b>='A' && b<='Z'
}

func IsDigit(b byte) bool {
	return b>='0' && b<='9'
}

func IsOctalDigit(b byte) bool {
	return b>='0' && b<='7'
}

func IsHexDigit(b byte) bool {
	return b>='0' && b<='9' ||
		b>='a' && b<='f' ||
		b>='A' && b<='F'
}

func HexDigit(b byte) (result byte) {
	switch {
	case b < 10:
		return  b + '0'
	case b < 16:
		return (b - 10) + 'A'
	}
	panic("value out of range")
}
