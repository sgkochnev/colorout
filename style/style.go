package style

type Style = uint8

const (
	Reset Style = iota
	Bold
	Dim
	Italic
	Underline
	Blink
	FastBlink
	Reverse
	Hidden
	Strikeout
	DoubleUnderline Style = 21
	Overline        Style = 53
)

//cancel style
const (
	NotBoldDim Style = iota + 22
	NotItalic
	NotUnderline
	NotBlink
	_
	NotReverse
	NotHidden
	NotStrikeout
	NotOverline Style = 55
)

func IsStyle(a Style) bool {
	return (a >= Reset && a <= Strikeout || a == DoubleUnderline || a == Overline)
}
