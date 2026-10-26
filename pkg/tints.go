package pkg

import "github.com/miguel-fg/gamut/pkg/utils"

func Tints(base utils.HSLColor, num int) []utils.HSLColor {

	n := utils.ClampPaletteSize(num)

	palette := make([]utils.HSLColor, 0)

	const STEP = 10

	i := 1
	light := base.Lum + STEP

	for i <= n && light <= 95 {
		palette = append(palette, utils.HSLColor{
			Hue: base.Hue,
			Sat: base.Sat,
			Lum: light,
		})

		light += STEP
		i++
	}

	return palette
}
