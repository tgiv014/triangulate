package behavior

import (
	"math/rand"

	"github.com/tgiv014/triangulate/geometry"
)

type TargetDepthBehavior struct {
	BaseBehavior
	targetDepth int
}

func NewTargetDepthBehavior() Behavior {
	return &TargetDepthBehavior{}
}

func (b *TargetDepthBehavior) Init(t geometry.Triangle) {
	b.targetDepth = t.Depth + rand.Intn(10) + 2
}

func (b *TargetDepthBehavior) Cycle(t geometry.Triangle) (vertex int, percent float64, ta, tb bool) {
	v := t.PtOppositeLongestSide()
	p := clamp(uniformNorm(0.5))

	return v, p, false, false
}

func (b *TargetDepthBehavior) ShouldDie(t geometry.Triangle) bool {
	if t.Depth >= b.targetDepth {
		return true
	}

	return b.BaseBehavior.ShouldDie(t)
}

func (b *TargetDepthBehavior) Inherit(t geometry.Triangle) Behavior {
	return b
}
