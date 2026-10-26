package utils

import (
	"bytes"
	"os"
	"os/exec"
	"strconv"
)

func pipeCommands(getVal bool, cmd1, cmd2 *exec.Cmd) bytes.Buffer {
	pipe, err := cmd1.StdoutPipe()

	if err != nil {
		panic(err)
	}

	cmd2.Stdin = pipe

	var buf bytes.Buffer

	if getVal {
		cmd2.Stdout = &buf
	} else {
		cmd2.Stdout = os.Stdout
	}

	if err := cmd1.Start(); err != nil {
		panic(err)
	}

	if err := cmd2.Start(); err != nil {
		panic(err)
	}

	if err := cmd1.Wait(); err != nil {
		panic(err)
	}

	if err := cmd2.Wait(); err != nil {
		panic(err)
	}

	return buf
}

func GetHSLOutput(input string) string {
	cmd1 := exec.Command("pastel", "color", input)
	cmd2 := exec.Command("pastel", "format", "hsl")

	buf := pipeCommands(true, cmd1, cmd2)

	return buf.String()
}

func buildHSLString(color HSLColor) string {
	hue := strconv.Itoa(color.Hue)
	sat := strconv.FormatFloat(color.Sat, 'f', 1, 64)
	lum := strconv.FormatFloat(color.Lum, 'f', 1, 64)

	return "hsl(" + hue + ", " + sat + "%, " + lum + "%)"
}

func PrintHSLColor(color HSLColor, format ColorFormat) {
	input := buildHSLString(color)

	cmd1 := exec.Command("pastel", "color", input)
	cmd2 := exec.Command("pastel", "format", format.String())

	pipeCommands(false, cmd1, cmd2)
}

func PrintPalette(color []HSLColor, format ColorFormat) {
	for i := range len(color) {
		PrintHSLColor(color[i], format)
	}
}
