package main

import (
	pkg "github.com/miguel-fg/gamut/pkg"
	utils "github.com/miguel-fg/gamut/pkg/utils"
)

func main() {
	input := pkg.ReadInput()
	base := utils.GenerateColor(input)

	hsl := utils.ToHSL(base)

	palette := pkg.Tones(hsl, 9)

	utils.PrintPalette(palette, utils.FormatHSL)
}
