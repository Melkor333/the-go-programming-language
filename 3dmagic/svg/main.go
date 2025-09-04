package svg

import (
	"3dmagic/obj"
	"strings"
	"fmt"
)

type Path struct {
	C *obj.Object
	Close bool
}

func (p *Path) String() string {
	var out strings.Builder
	out.WriteString(`<path  style="fill:none;stroke:green;stroke-width:1"`)
	out.WriteString(fmt.Sprintf(` d="M%v %v`, p.C.V[0].X, p.C.V[0].Y))
	for _, i := range p.C.V[1:] { 
		out.WriteString(fmt.Sprintf("L%v %v ", i.X, i.Y))
	}
	if p.Close {
		out.WriteRune('Z')
	}
	out.WriteString("\" />")
	return out.String()
}
