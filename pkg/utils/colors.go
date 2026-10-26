package utils

import (
	"errors"
	"log"
	"math"
	"strconv"
	"strings"
)

func GenerateColor(str string) Color {
	format, err := getColorFormat(str)

	if err != nil {
		log.Fatal(err.Error())
	}

	color := Color{
		Value:  str,
		Format: format,
	}

	return color
}

func getColorFormat(str string) (ColorFormat, error) {
	if Hex.MatchString(str) {
		return FormatHex, nil
	}

	if Rgb.MatchString((str)) {
		return FormatRGB, nil
	}

	if Hsl.MatchString((str)) {
		return FormatHSL, nil
	}

	return FormatUnknown, errors.New("Input: `" + str + "` is not a valid color format.")
}

func ToHSL(color Color) HSLColor {
	hsl := GetHSLOutput(color.Value)

	hue, sat, lum := extractHSLValues(hsl)

	return HSLColor{
		Hue: hue,
		Sat: sat,
		Lum: lum,
	}
}

func extractHSLValues(hslStr string) (h int, s float64, l float64) {
	split := strings.Split(hslStr, ",")

	hStr := split[0][4:]
	sStr := split[1][1 : len(split[1])-1]
	lStr := split[2][1 : len(split[2])-3]

	hue, err := strconv.Atoi(hStr)

	if err != nil {
		panic(err)
	}

	sVal, err := strconv.ParseFloat(sStr, 32)

	sat := math.Round(sVal*10) / 10

	if err != nil {
		panic(err)
	}

	lVal, err := strconv.ParseFloat(lStr, 32)
	lum := math.Round(lVal*10) / 10

	if err != nil {
		panic(err)
	}

	return hue, sat, lum
}

func GetPalleteOffsetArrays(step int) ([MAX_PALETTE_SIZE]int, [11]int) {
	steps := [MAX_PALETTE_SIZE]int{
		step * -4,
		step * -3,
		step * -2,
		step * -1,
		0,
		step * 1,
		step * 2,
		step * 3,
		step * 4,
	}

	stepStartIndex := [11]int{-1, -1, -1, 3, 3, 2, 2, 1, 1, 0, 0}

	return steps, stepStartIndex
}

func ClampPaletteSize(size int) int {
	return max(MIN_PALETTE_SIZE, min(MAX_PALETTE_SIZE, size))
}

func ClampAngle(angle int) int {
	return ((angle % MAX_ANGLE) + MAX_ANGLE) % MAX_ANGLE
}
