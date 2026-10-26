package pkg

import "github.com/miguel-fg/gamut/pkg/utils"

func bindPaletteParams(num int, offset int) (int, int) {
	if num < 3 {
		num = 3
	}

	if num > 9 {
		num = 3
	}

	for offset < 0 {
		offset = 360 - offset
	}

	for offset > 360 {
		offset = offset - 360
	}

	return num, offset
}

func buildAngleArray(angle int) ([9]int, [11]int) {
	angles := [9]int{
		angle * -4,
		angle * -3,
		angle * -2,
		angle * -1,
		0,
		angle * 1,
		angle * 2,
		angle * 3,
		angle * 4,
	}

	angleStartIndex := [11]int{-1, -1, -1, 3, 3, 2, 2, 1, 1, 0, 0}

	return angles, angleStartIndex
}

func Adjacent(base utils.HSLColor, num int, offset int) []utils.HSLColor {

	n, a := bindPaletteParams(num, offset)

	angles, angleIndex := buildAngleArray(a)

	palette := make([]utils.HSLColor, 0)

	for i := range n {
		newHue := base.Hue + angles[angleIndex[n]+i]

		if newHue > 360 {
			newHue = newHue % 360
		} else if newHue < 0 {
			newHue = newHue % 360
			newHue = 360 + newHue
		}

		palette = append(palette,
			utils.HSLColor{
				Hue: newHue,
				Sat: base.Sat,
				Lum: base.Lum,
			},
		)
	}

	return palette
}
