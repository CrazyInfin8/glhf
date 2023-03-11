package glhf

import (
	"github.com/crazyinfin8/glhf/math"
)

type (
	FrameCollection struct {
		parent   *Frame
		frames   []*Frame
		frameMap map[string]*Frame
	}
)

func NewFrameCollection(parent *Frame) *FrameCollection {
	c := new(FrameCollection)

	c.parent = parent

	return c
}

func NewFrameCollectionFromAnimation(parent *Frame, width, height int) *FrameCollection {
	assetWidth, assetHeight := parent.Width(), parent.Height()
	if width == 0 && height == 0 {
		width = math.Min(assetWidth, assetHeight)
		height = width
	} else if width == 0 {
		width = height
	} else if height == 0 {
		height = width
	}

	c := new(FrameCollection)

	c.parent = parent

	for y := 0; y+height <= assetHeight; y += height {
		for x := 0; x+width <= assetWidth; x += width {
			println("(", x, ":", y, ")", len(c.frames))
			c.frames = append(c.frames, parent.SubFrame(x, y, width, height))
		}
	}

	return c
}

func (c *FrameCollection) NumFrames() int { return len(c.frames) }

func (c *FrameCollection) GetFrame(index int) (*Frame, bool) {
	if len(c.frames) == 0 && index == 0 {
		return c.parent, true
	}
	if index >= c.NumFrames() {
		return nil, false
	}
	frame := c.frames[index]
	return frame, true
}

func (c *FrameCollection) AddFrame(frame *Frame) {
	c.frames = append(c.frames, frame)
}

func (c *FrameCollection) AddSubFrame(name string, x, y, width, height int) *Frame {
	subframe := c.parent.SubFrame(x, y, width, height)
	c.frames = append(c.frames, subframe)
	if name != "" {
		c.frameMap[name] = subframe
	}
	return subframe
}
