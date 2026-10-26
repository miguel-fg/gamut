package pkg

import (
	"math"
	"slices"

	"github.com/miguel-fg/gamut/pkg/utils"
)

func Monochromatic(base utils.HSLColor, num int) []utils.HSLColor {

	n := utils.ClampPaletteSize(num)

	palette := make([]utils.HSLColor, 0)

	shadesLength := int(math.Floor(float64(n-1) / 2))
	tintsLength := n - 1 - shadesLength

	shades := Shades(base, shadesLength)
	tints := Tints(base, tintsLength)

	for _, tint := range slices.Backward(tints) {
		palette = append(palette, tint)
	}

	palette = append(palette, base)
	palette = append(palette, shades...)

	return palette
}
