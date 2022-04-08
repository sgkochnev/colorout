package bg

type Background = uint8

const (
	Black Background = iota + 40
	Red
	Green
	Yellow
	Blue
	Purple
	Cyan
	LightGray
	Custom
	Default
)

const (
	DarkGray Background = iota + 100
	LightRed
	LightGreen
	LightYellow
	LightBlue
	LightMagenta
	LightCyan
	White
)

func IsDefault(a Background) bool {
	return a == Default
}
func IsColor(a Background) bool {
	return a >= Black && a <= LightGray || a >= DarkGray && a <= White
}
func IsCustom(a Background) bool {
	return a == Custom
}
