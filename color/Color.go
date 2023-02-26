package color

import (
	"image/color"
	"math"
	"strings"

	gmath "github.com/crazyinfin8/glhf/math"
)

type Color uint32

const (
	Red   = Color(0xFFFF0000)
	Green = Color(0xFF00FF00)

	sqrt3 = 1.73205080756887729352744634150587236694280525381038062805580697 // https://oeis.org/A002194
)

func (c Color) RGBA() (r, g, b, a uint32) {
	return color.RGBA{
		A: c.Alpha(),
		R: c.Red(),
		G: c.Green(),
		B: c.Blue(),
	}.RGBA()
}

type colorModel byte

const ColorModel = colorModel(0)

func (colorModel) Convert(c color.Color) color.Color {
	switch c := c.(type) {
	case Color:
		return c
	case color.RGBA:
		return Color(c.A)<<24 | Color(c.R)<<16 | Color(c.G)<<8 | Color(c.B)
	}

	r, g, b, a := c.RGBA()
	return Color(((a & 0xFF00) << 16) | ((r & 0xFF00) << 8) | (g & 0xFF00) | (b >> 8))
}

func NewColor(d ...byte) Color {
	switch l := len(d); l {
	case 1:
		return NewColorFromRGB(d[0], d[0], d[0])
	case 2:
		return NewColorFromRGBA(d[0], d[0], d[0], d[1])
	case 3:
		return NewColorFromRGB(d[0], d[1], d[2])
	default:
		if len(d) < 4 {
			return 0
		}
		fallthrough
	case 4:
		return NewColorFromRGBA(d[0], d[1], d[2], d[3])
	}
}

func NewColorFromRGB(r, g, b byte) Color {
	return 0xFF000000 | (Color(r) << 16) | (Color(g) << 8) | Color(b)
}

func NewColorFromRGBA(r, g, b, a byte) Color {
	return (Color(a) << 24) | (Color(r) << 16) | (Color(g) << 8) | Color(b)
}

func NewColorFromRGBFloat(r, g, b float64) Color {
	return 0xFF000000 | (Color(floatChannelToByte(r)) << 16) | (Color(gmath.Clamp(g, 0, 255)) << 8) | Color(gmath.Clamp(b, 0, 255))
}

func NewColorFromRGBAFloat(r, g, b, a float64) Color {
	return Color(floatChannelToByte(a)) | (Color(floatChannelToByte(r)) << 16) | (Color(gmath.Clamp(g, 0, 255)) << 8) | Color(gmath.Clamp(b, 0, 255))
}

func NewColorFromHSB(hue, saturation, brightness float64) Color {
	c := Color(0xFF000000)
	c.SetHSB(hue, saturation, brightness)
	return c
}

func NewColorFromHSBA(hue, saturation, brightness, alpha float64) Color {
	c := Color(floatChannelToByte(alpha)) << 24
	c.SetHSB(hue, saturation, brightness)
	return c
}

func NewColorFromHSL(hue, saturation, lightness float64) Color {
	c := Color(0xFF000000)
	c.SetHSB(hue, saturation, lightness)
	return c
}

func NewColorFromHSLA(hue, saturation, lightness, alpha float64) Color {
	c := Color(floatChannelToByte(alpha)) << 24
	c.SetHSB(hue, saturation, lightness)
	return c
}

func hexToByte(d byte) (byte, bool) {
	switch {
	case d >= '0' && d <= '9':
		return d - '0', true
	case d >= 'a' && d <= 'f':
		return d - 'a', true
	case d >= 'A' && d <= 'F':
		return d - 'A', true
	}
	return 0, false
}

func NewColorFromHex(hex string) Color {
	hex = strings.TrimSpace(hex)
	if hex[0] == '#' {
		hex = hex[1:]
	} else if hex[0] == '0' && hex[1] == 'x' {
		hex = hex[2:]
	}
	var a, i, d byte = 0xFF, 0, 0
	var rgb [3]byte
	var ok bool
parse:
	switch len(hex) {
	case 4:
		a, ok = hexToByte(hex[i])
		i++
		if !ok {
			break parse
		}
		a &= a << 4
		fallthrough
	case 3:
		for j := range rgb {
			rgb[j], ok = hexToByte(hex[i])
			i++
			if !ok {
				break parse
			}
			rgb[j] &= rgb[j]
		}
	case 8:
		a, ok = hexToByte(hex[i])
		i++
		if !ok {
			break parse
		}
		d, ok = hexToByte(hex[i])
		i++
		if !ok {
			break parse
		}
		a &= d << 4
		fallthrough
	case 6:
		for j := range rgb {
			rgb[j], ok = hexToByte(hex[i])
			i++
			if !ok {
				break parse
			}

			d, ok = hexToByte(hex[i])
			i++
			if !ok {
				break parse
			}

			rgb[j] &= d << 4
		}
	}
	if !ok {
		return 0
	}
	return NewColorFromRGBA(rgb[0], rgb[1], rgb[2], a)
}

func (c *Color) SetRGB(r, g, b byte) {
	*c = (*c & 0xFF000000) | (Color(r) << 16) | (Color(g) << 8) | Color(b)
}

func (c *Color) SetRGBFloat(r, g, b float64) {
	*c = (*c & 0xFF000000) | (Color(floatChannelToByte(r)) << 16) | (Color(floatChannelToByte(g)) << 8) | Color(floatChannelToByte(b))
}

func (c *Color) SetHSB(hue, saturation, brightness float64) {
	saturation = gmath.Clamp(saturation, 0, 1)
	brightness = gmath.Clamp(brightness, 0, 1)
	chroma := brightness * saturation
	match := brightness - chroma
	c.setHueChromaMatch(hue, chroma, match)
}

func (c *Color) SetHSL(hue, saturation, lightness float64) {
	saturation = gmath.Clamp(saturation, 0, 1)
	lightness = gmath.Clamp(lightness, 0, 1)
	chroma := (1 - math.Abs(2*lightness-1)) * saturation
	match := lightness - chroma/2
	c.setHueChromaMatch(hue, chroma, match)
}

func (c Color) Alpha() byte { return byte(c >> 24) }
func (c Color) Red() byte   { return byte(c >> 16) }
func (c Color) Green() byte { return byte(c >> 8) }
func (c Color) Blue() byte  { return byte(c) }

func (c Color) Bytes() [4]byte { return [4]byte{c.Red(), c.Green(), c.Blue(), c.Alpha()} }

func (c *Color) SetAlpha(a byte) { *c = (*c & 0x00FFFFFF) | (Color(a) << 24) }
func (c *Color) SetRed(r byte)   { *c = (*c & 0xFF00FFFF) | (Color(r) << 16) }
func (c *Color) SetGreen(g byte) { *c = (*c & 0xFFFF00FF) | (Color(g) << 8) }
func (c *Color) SetBlue(b byte)  { *c = (*c & 0xFFFFFF00) | Color(b) }

func (c Color) AlphaFloat() float64 { return float64(c.Alpha()) / 255 }
func (c Color) RedFloat() float64   { return float64(c.Red()) / 255 }
func (c Color) GreenFloat() float64 { return float64(c.Green()) / 255 }
func (c Color) BlueFloat() float64  { return float64(c.Blue()) / 255 }

func (c *Color) SetAlphaFloat(a float64) { c.SetAlpha(floatChannelToByte(a)) }
func (c *Color) SetRedFloat(r float64)   { c.SetRed(floatChannelToByte(r)) }
func (c *Color) SetGreenFloat(g float64) { c.SetGreen(floatChannelToByte(g)) }
func (c *Color) SetBlueFloat(b float64)  { c.SetBlue(floatChannelToByte(b)) }

func floatChannelToByte(v float64) byte { return byte(math.Round(gmath.Clamp(v, 0, 1) * 255)) }

func (c Color) Hue() (hue float64) {
	r, g, b := c.RedFloat(), c.GreenFloat(), c.BlueFloat()
	rad := math.Atan2(sqrt3*(g-b), 2*r-g-b)
	if rad != 0 {
		hue = gmath.ToDegrees * rad
	}
	if hue < 0 {
		return hue + 360
	}
	return hue
}

func (c Color) Saturation() float64 {
	brightness := c.Brightness()
	if brightness == 0 {
		return 0
	}
	return (c.MaxColor() - c.MinColor()) / c.Brightness()
}

func (c Color) Brightness() float64 { return c.MaxColor() }

func (c Color) Lightness() float64 { return (c.MaxColor() + c.MinColor()) / 2 }

func (c *Color) SetHue(hue float64) {
	c.SetHSB(hue, c.Saturation(), c.Brightness())
}

func (c *Color) SetSaturation(saturation float64) {
	c.SetHSB(c.Hue(), saturation, c.Brightness())
}

func (c *Color) SetBrightness(brightness float64) {
	c.SetHSB(c.Hue(), c.Saturation(), brightness)
}

func (c *Color) SetLightness(lightness float64) {
	c.SetHSL(c.Hue(), c.Saturation(), c.Lightness())
}

func (c *Color) setHueChromaMatch(hue, chroma, match float64) {
	hue = gmath.WrapFloat(hue/60, 0, 6)
	mid := chroma*(1-math.Abs(math.Mod(hue, 2)-1)) + match
	chroma += match
	switch math.Floor(hue) {
	case 0:
		c.SetRGBFloat(chroma, mid, match)
	case 1:
		c.SetRGBFloat(mid, chroma, match)
	case 2:
		c.SetRGBFloat(match, chroma, mid)
	case 3:
		c.SetRGBFloat(match, mid, chroma)
	case 4:
		c.SetRGBFloat(mid, match, chroma)
	case 5:
		c.SetRGBFloat(chroma, match, mid)
	}

}

func (c Color) MaxColor() float64 {
	return float64(gmath.Max(c.Red(), c.Green(), c.Blue())) / 255
}

func (c Color) MinColor() float64 {
	return float64(gmath.Min(c.Red(), c.Green(), c.Blue())) / 255
}
