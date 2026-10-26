package pkg

import "github.com/miguel-fg/gamut/pkg/utils"

func bindTintsParams(num int) int {
	if num < 1 {
		num = 1
	}

	return num
}

func Tints(base utils.HSLColor, num int) []utils.HSLColor {

	n := bindTintsParams(num)

	palette := make([]utils.HSLColor, 0)

	const STEP = 5

	i := 1
	light := base.Lum

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
