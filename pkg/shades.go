package pkg

import "github.com/miguel-fg/gamut/pkg/utils"

func bindShadesParams(num int) int {
	if num < 1 {
		num = 1
	}

	return num
}

func Shades(base utils.HSLColor, num int) []utils.HSLColor {

	n := bindShadesParams(num)

	palette := make([]utils.HSLColor, 0)

	const STEP = 5

	i := 1
	light := base.Lum

	for i <= n && light >= 5 {
		palette = append(palette, utils.HSLColor{
			Hue: base.Hue,
			Sat: base.Sat,
			Lum: light,
		})

		light -= STEP
		i++
	}

	return palette
}
