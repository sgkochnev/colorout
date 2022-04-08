package color

import (
	"github.com/sgkochnev/colorout/color/bg"
	"github.com/sgkochnev/colorout/color/fg"
)

const customColorArg = 2

type Color struct {
	r uint8
	g uint8
	b uint8
}

func New(r, g, b uint8) *Color {
	return &Color{
		r: r,
		g: g,
		b: b,
	}
}

func (c *Color) Back() []bg.Background {
	return []bg.Background{
		bg.Custom, customColorArg,
		c.r, c.g, c.b,
	}
}

func (c *Color) Fore() []fg.Foreground {
	return []fg.Foreground{
		fg.Custom, customColorArg,
		c.r, c.g, c.b,
	}
}
