package fg

type Foreground = uint8

const (
	Black Foreground = iota + 30
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
	DarkGray Foreground = iota + 90
	LightRed
	LightGreen
	LightYellow
	LightBlue
	LightMagenta
	LightCyan
	White
)

func IsDefault(a Foreground) bool {
	return a == Default
}
func IsColor(a Foreground) bool {
	return a >= Black && a <= LightGray || a >= DarkGray && a <= White
}
func IsCustom(a Foreground) bool {
	return a == Custom
}
