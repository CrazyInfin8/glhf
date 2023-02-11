package color

import (
	"fmt"
	"testing"
)

type colorTest struct {
	color Color

	hue, saturation, brightness, lightness float64
}

var cases = []colorTest{
	{0x000000, 0, 0, 0, 0},     // Black
	{0xFFFFFF, 0, 0, 1, 0},     // White
	{0xFF0000, 0, 1, 1, 0},     // Red
	{0x00FF00, 120, 1, 1, 0},   // Lime
	{0x0000FF, 240, 1, 1, 0},   // Blue
	{0xFFFF00, 60, 1, 1, 0},    // Yellow
	{0x00FFFF, 180, 1, 1, 0},   // Cyan
	{0xFF00FF, 300, 1, 1, 0},   // Magenta
	{0xBFBFBF, 0, 0, 0.75, 0},  // Silver
	{0x808080, 0, 0, 0.5, 0},   // Gray
	{0x800000, 0, 1, 0.5, 0},   // Maroon
	{0x808000, 60, 1, 0.5, 0},  // Olive
	{0x008000, 120, 1, 0.5, 0}, // Green
	{0x800080, 300, 1, 0.5, 0}, // Purple
	{0x008080, 180, 1, 0.5, 0}, // Teal
	{0x000080, 240, 1, 0.5, 0}, // Navy
}

func Test(t *testing.T) {
	var c Color
	for _, test := range cases {
		c.SetHSB(test.hue, test.saturation, test.brightness)
		if test.color != c {
			fmt.Printf("Values don't match: ( #%08X  |  #%08X )\n", test.color, c)
		}
		if !testHSBAboutMatches(c, test.color) {
			fmt.Printf("HSBs don't match: (h: %g, s: %g, b: %g)  |  (h: %g, s: %g, b: %g)\n", test.hue, test.saturation, test.brightness, c.Hue(), c.Saturation(), c.Brightness())
		}
	}
}

func floatAboutEqual(a, b, threshold float64) bool {
	if a > b {
		a -= b
	} else {
		a = b - a
	}
	return a < threshold
}

func testHSBAboutMatches(c1, c2 Color) bool {
	return floatAboutEqual(c1.Hue(), c2.Hue(), 0.1) &&
		floatAboutEqual(c1.Saturation(), c2.Saturation(), 0.0001) &&
		floatAboutEqual(c1.Brightness(), c2.Brightness(), 0.0001)
}
