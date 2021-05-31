package behavior

import (
	"math/rand"

	"github.com/tgiv014/triangulate/geometry"
)

type UniformBehavior struct {
	BaseBehavior
}

func NewUniformBehavior() Behavior {
	return &UniformBehavior{}
}

func (b *UniformBehavior) Cycle(t geometry.Triangle) (vertex int, percent float64, ta, tb bool) {
	v := t.PtOppositeLongestSide()
	p := 0.5
	return v, p, false, false
}

func (b *UniformBehavior) ShouldDie(t geometry.Triangle) bool {
	if t.Depth > 0 && rand.Float64() < 0.33 {
		return true
	}

	return b.BaseBehavior.ShouldDie(t)
}

func (b *UniformBehavior) Inherit(t geometry.Triangle) Behavior {
	return b
}
