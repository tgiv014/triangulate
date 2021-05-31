package behavior

import (
	"math/rand"

	"github.com/tgiv014/triangulate/geometry"
)

type Behavior interface {
	Init(t geometry.Triangle)
	Cycle(t geometry.Triangle) (vertex int, percent float64, ta, tb bool)
	ShouldDie(t geometry.Triangle) bool
	Inherit(t geometry.Triangle) Behavior
}

type UndecidedBehavior struct {
	initialized        bool
	decisiondepth      int
	underLyingBehavior Behavior
	decision           Behavior
}

func NewUndecidedBehavior(decisionDepth int) Behavior {
	return &UndecidedBehavior{
		false,
		decisionDepth,
		NewBalancedBehavior(),
		nil,
	}
}

func (b *UndecidedBehavior) Init(t geometry.Triangle) {
	if !b.initialized {
		b.underLyingBehavior.Init(t)
		b.decision = nil
		b.initialized = true
	}
}

func (b *UndecidedBehavior) Cycle(t geometry.Triangle) (vertex int, percent float64, ta, tb bool) {
	if !b.initialized {
		panic("A behavior was cycled prior to initialization")
	}

	return b.underLyingBehavior.Cycle(t)
}

func (b *UndecidedBehavior) ShouldDie(t geometry.Triangle) bool {
	return b.underLyingBehavior.ShouldDie(t)
}

func (b *UndecidedBehavior) Inherit(t geometry.Triangle) Behavior {
	// This will be overridden to become an exciting new thing
	if t.Depth >= b.decisiondepth {
		// It's decision time!
		if b.decision == nil {
			b.decision = RandomBehavior()
		}
		return b.decision
	}
	return b
}

func uniformNorm(std float64) float64 {
	return rand.NormFloat64()*std + 0.5
}

func clamp(n float64) float64 {
	if n >= 1 {
		return 1
	}
	if n <= 0 {
		return 0
	}
	return n
}

func RandomBehavior() Behavior {
	var newBehavior Behavior

	switch rand.Intn(4) {
	case 0:
		newBehavior = NewBalancedBehavior()
		break
	case 1:
		newBehavior = NewDepthChargeBehavior(4)
		break
	case 2:
		newBehavior = NewSpinBehavior(rand.Float64()*0.4 + 0.1)
		break
	case 3:
		newBehavior = NewUniformBehavior()
		break
	}
	return newBehavior
}
