package main

import (
	"fmt"

	pkg "github.com/miguel-fg/gamut/pkg"
	utils "github.com/miguel-fg/gamut/pkg/utils"
)

func main() {
	input := pkg.ReadInput()
	base := utils.GenerateColor(input)

	hsl := utils.ToHSL(base)

	adjacent := pkg.Adjacent(hsl, 9, 15) // Adjacent
	shades := pkg.Shades(hsl, 9)         // Shades
	tints := pkg.Tints(hsl, 9)           // Tints

	fmt.Println("Adjacent")
	utils.PrintPalette(adjacent, utils.FormatHSL)
	fmt.Println()

	fmt.Println("Shades")
	utils.PrintPalette(shades, utils.FormatHSL)
	fmt.Println()

	fmt.Println("Tints")
	utils.PrintPalette(tints, utils.FormatHSL)
	fmt.Println()
}
