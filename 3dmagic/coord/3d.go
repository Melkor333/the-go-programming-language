package coord

import (
	"3dmagic/obj"
	"math"
)
const (
	// 45degrees to radians
	slope = 45. * math.Pi / 180
)

type Point struct{
	X, Y float64
}
type Line []*Point

func perspective(cloud *obj.Object, angle int) *obj.Object {
	o := obj.NewObject()
	o.Edges = make(obj.Edges, len(cloud.Edges))
	copy(o.Edges, cloud.Edges)
	for i, v := range cloud.V {
		c := o.V[i]
		c.X = v.X + math.Cos(slope) * v.Z
		c.Y = v.Y + math.Sin(slope) * v.Z
		c.Z = 0
	}
	return o
}
