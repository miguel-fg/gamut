package pkg

import (
	"github.com/miguel-fg/gamut/pkg/utils"
)

func Tones(base utils.HSLColor, num int) []utils.HSLColor {

	n := utils.ClampPaletteSize(num)

	const STEP = 10
	steps, stepIndex := utils.GetPalleteOffsetArrays(STEP)

	palette := make([]utils.HSLColor, 0)

	start := base.Sat + float64(steps[stepIndex[n]])
	start = max(0, min(100-float64(STEP*(n-1)), start))

	for i := range n {
		palette = append(palette,
			utils.HSLColor{
				Hue: base.Hue,
				Sat: start + float64(STEP*i),
				Lum: base.Lum,
			},
		)
	}

	return palette
}
