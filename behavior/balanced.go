package behavior

import (
	"math/rand"

	"github.com/tgiv014/triangulate/geometry"
)

type BalancedBehavior struct {
	BaseBehavior
}

func NewBalancedBehavior() Behavior {
	return &BalancedBehavior{}
}

func (b *BalancedBehavior) Cycle(t geometry.Triangle) (vertex int, percent float64, ta, tb bool) {
	v := t.PtOppositeLongestSide()
	p := clamp(uniformNorm(0.2))
	return v, p, false, false
}

func (b *BalancedBehavior) ShouldDie(t geometry.Triangle) bool {
	if t.Depth > 0 && rand.Float64() < 0.05 {
		return true
	}
	return b.BaseBehavior.ShouldDie(t)
}

func (b *BalancedBehavior) Inherit(t geometry.Triangle) Behavior {
	return b
}
