package clevertriangle

import (
	"github.com/tgiv014/triangulate/behavior"
	"github.com/tgiv014/triangulate/geometry"
)

type CleverTriangle struct {
	geometry.Triangle
	dead     bool
	Behavior behavior.Behavior
}

func WrapTriangle(triangle geometry.Triangle, behavior behavior.Behavior) CleverTriangle {
	return CleverTriangle{
		triangle,
		false,
		behavior,
	}
}

func NewCleverTriangle(A, B, C geometry.Point, behavior behavior.Behavior) CleverTriangle {
	t := geometry.Triangle{A: A, B: B, C: C}
	behavior.Init(t)
	return CleverTriangle{
		t,
		false,
		behavior,
	}
}

func (t *CleverTriangle) Cycle() []CleverTriangle {
	if t.dead {
		return []CleverTriangle{*t}
	}
	if t.Behavior.ShouldDie(t.Triangle) {
		t.dead = true
		return []CleverTriangle{*t}
	}
	pt, percent, da, db := t.Behavior.Cycle(t.Triangle)
	triangles := t.Subdivide(pt, percent)
	ta := triangles[0]
	tb := triangles[1]
	return []CleverTriangle{
		{ta, t.dead || da, t.Behavior.Inherit(ta)},
		{tb, t.dead || db, t.Behavior.Inherit(tb)},
	}
}
