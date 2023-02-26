package glhf

type (
	FrameCollection struct {
		frames   []*Frame
		frameMap map[string]*Frame
	}
)

func NewFrameCollection(frames ...*Frame) *FrameCollection {
	c := new(FrameCollection)

	c.frames = frames

	return c
}

func (c *FrameCollection) NumFrames() int { return len(c.frames) }

func (c *FrameCollection) GetFrame(index int) (*Frame, bool) {
	frame := c.frames[index]
	return frame, frame != nil
}

func (c *FrameCollection) AddFrame(frame *Frame) {
	c.frames = append(c.frames, frame)

}
