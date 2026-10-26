package utils

type ColorFormat int

const MIN_PALETTE_SIZE = 3
const MAX_PALETTE_SIZE = 9
const MAX_ANGLE = 360
const MIN_ANGLE = 0

const (
	FormatHex ColorFormat = iota
	FormatRGB
	FormatHSL
	FormatUnknown
)

var formatName = map[ColorFormat]string{
	FormatHex: "hex",
	FormatRGB: "rgb",
	FormatHSL: "hsl",
}

func (fmt ColorFormat) String() string {
	return formatName[fmt]
}

type Color struct {
	Value  string
	Format ColorFormat
}

type HSLColor struct {
	Hue int
	Sat float64
	Lum float64
}
