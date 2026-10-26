package utils

type ColorFormat int

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
