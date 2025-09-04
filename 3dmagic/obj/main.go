// function to render all coordinates in a raster 0-1
package obj

import (
	"errors"
	"math"
	"slices"
)

type Vertex struct{
	X, Y, Z float64
	// the index in the object
	O *Object
}
type Edge [2]int
// TODO: we can't just delete a vertex from the array or it breaks the edges

type Object struct{
	V []*Vertex
	Edges []Edge
}

func NewObject() *Object{
	return NewObjectSized(0)
}
func NewObjectSized(i int) *Object{
	o := &Object{
		V: make([]*Vertex, i),
		Edges: make([]Edge, 0),
	}
	return o
}


func (v *Vertex) Connect(v2 *Vertex) error {
	// Never merge two separate objects. That doesn't work.
	//assert &(v.O) == &(v2.O)

	// generate the edge, which is just 2 slice indexes
	var edge [2]int
	ei := 0
	// find the index of both vertices. could optionally be stored in the vertice itself
	// TODO: move into Vertex.Index()
	// though that would create 2 loops...
	for i, o := range v.O.V {
		if o == v || o == v2 {
			edge[ei] = i
			ei += 1
		}
		if ei > 1 {
			break
		}
	}
	if !slices.Contains(v.O.Edges, edge) {
		v.O.Edges = append(v.O.Edges, edge)
		return nil
	}
	return  errors.New("Vertices already connected")
}

func RegularPolygon(edgeCount int, rad, center float64) *Object {
	o := NewObjectSized(edgeCount)
	for i := range edgeCount {
		v := &Vertex{}
		o.V[i] = v
		// get degree and convert to rad
		degree := 360. / float64(edgeCount) * float64(i) * math.Pi / 180
		v.O = o
		v.X = math.Cos(degree)
		v.Y = math.Sin(degree)
		// resize and shift around center
		v.X = (v.X * rad) + center
		v.Y = (v.Y * rad) + center
		if i > 0 {
			v.Connect(o.V[i-1])
		}
		if i == edgeCount {
			v.Connect(o.V[0])
		}
	}
	return o
}
