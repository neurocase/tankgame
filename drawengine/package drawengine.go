package drawengine


import (
	"fmt"
	gameloop "github.com/GlenKelley/go-glutil/gameloop"
	"github.com/go-gl/gl"
	glfw "github.com/go-gl/glfw3"
	"math"
)

func DrawRTri(Red float64){
	gl.ClearColor(0.2, 0.2, 0.2, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	gl.Color3d(game.Red, 0, 0)
	gl.Begin(gl.TRIANGLES)
	gl.Vertex3d(0, 0, 0)
	gl.Vertex3d(0, 1, 0)
	gl.Vertex3d(1, 1, 0)
	gl.Vertex3d(1, 0, 0)
	gl.End()
}