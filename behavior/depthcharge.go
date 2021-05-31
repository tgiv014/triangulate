package behavior

import "github.com/tgiv014/triangulate/geometry"

type DepthChargeBehavior struct {
	initialized        bool
	changeDepth        int
	underLyingBehavior Behavior
	changeBehavior     func() Behavior
}

func NewDepthChargeBehavior(decisionDepth int) Behavior {
	return &DepthChargeBehavior{
		false,
		decisionDepth,
		NewBalancedBehavior(),
		NewTargetDepthBehavior,
	}
}

func (b *DepthChargeBehavior) Init(t geometry.Triangle) {
	if !b.initialized {
		b.underLyingBehavior.Init(t)
		b.initialized = true
	}
}

func (b *DepthChargeBehavior) Cycle(t geometry.Triangle) (vertex int, percent float64, ta, tb bool) {
	if !b.initialized {
		panic("A behavior was cycled prior to initialization")
	}

	return b.underLyingBehavior.Cycle(t)
}

func (b *DepthChargeBehavior) ShouldDie(t geometry.Triangle) bool {
	return b.underLyingBehavior.ShouldDie(t)
}

func (b *DepthChargeBehavior) Inherit(t geometry.Triangle) Behavior {
	// This will be overridden to become an exciting new thing
	if t.Depth >= b.changeDepth {
		newBehavior := b.changeBehavior()
		newBehavior.Init(t)

		return newBehavior
	}
	return b
}
