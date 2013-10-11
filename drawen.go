package tankgame

import (
//	"fmt"
	"github.com/go-gl/gl"
	"math"

)


func Hello() string{

	return "drawengine"
}


func DrawEntity(en Entity){

x := en.Xpos
y := en.Ypos
r := en.Rot
s := en.Size

enx := [3]float64{-1*s,1*s,0*s}
eny := [3]float64{1*s,1*s,-1*s}

enxr := [3]float64 {enx[0],enx[1],enx[2]}
enyr := [3]float64 {eny[0],eny[1],eny[2]}

	for i := 0; i < 3; i++{
		enx[i] = enxr[i]*math.Cos(r)-enyr[i]*math.Sin(r)
		eny[i] = enxr[i]*math.Sin(r)+enyr[i]*math.Cos(r)
		
	}

	gl.Color3d(255, 0, 0)
	gl.Begin(gl.TRIANGLES)
	gl.Vertex3d(x+enx[0], y+eny[0], 0)
	gl.Vertex3d(x+enx[1], y+eny[1], 0)
	gl.Vertex3d(x+enx[2], y+eny[2], 0)
	gl.End()
}