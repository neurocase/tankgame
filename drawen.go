package tankgame

import (
//	"fmt"
	"github.com/go-gl/gl"
	//"math"

)


func Hello() string{

	return "drawengine"
}

func DrawEntity(en Entity){
x := en.Xpos
y := en.Ypos

	gl.Color3d(255, 0, 0)
	gl.Begin(gl.TRIANGLES)
	gl.Vertex3d(x, y, 0)
	gl.Vertex3d(x, y+1, 0)
	gl.Vertex3d(x+1, y+1, 0)
	gl.Vertex3d(x+1, y, 0)
	gl.End()
}