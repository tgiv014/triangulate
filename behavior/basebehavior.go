package behavior

import (
	"github.com/tgiv014/triangulate/geometry"
)

type BaseBehavior struct {
}

func NewBaseBehavior() Behavior {
	return &BaseBehavior{}
}

func (b *BaseBehavior) Init(t geometry.Triangle) {
}

func (b *BaseBehavior) Cycle(t geometry.Triangle) (vertex int, percent float64, ta, tb bool) {
	v := 0
	p := 0.5
	return v, p, false, false
}

func (b *BaseBehavior) ShouldDie(t geometry.Triangle) bool {
	return t.Area() < 3000
}

func (b *BaseBehavior) Inherit(t geometry.Triangle) Behavior {
	return b
}
