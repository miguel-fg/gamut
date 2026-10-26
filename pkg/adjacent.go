package pkg

import "github.com/miguel-fg/gamut/pkg/utils"

func Adjacent(base utils.HSLColor, num int, offset int) []utils.HSLColor {

	n := utils.ClampPaletteSize(num)

	const OFFSET = 15
	angles, angleIndex := utils.GetPalleteOffsetArrays(OFFSET)

	palette := make([]utils.HSLColor, 0)

	for i := range n {
		newHue := base.Hue + angles[angleIndex[n]+i]

		newHue = utils.ClampAngle(newHue)

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
