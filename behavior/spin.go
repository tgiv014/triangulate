package behavior

import (
	"github.com/tgiv014/triangulate/geometry"
)

type SpinBehavior struct {
	BaseBehavior
	initialized bool
	spin        float64
}

func NewSpinBehavior(spin float64) Behavior {
	return &SpinBehavior{
		BaseBehavior{},
		false,
		spin,
	}
}

func (b *SpinBehavior) Cycle(t geometry.Triangle) (vertex int, percent float64, ta, tb bool) {
	v := 1
	p := b.spin
	return v, p, true, false
}

func (b *SpinBehavior) Inherit(t geometry.Triangle) Behavior {
	return b
}
