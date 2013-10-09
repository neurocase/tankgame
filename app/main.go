package main

import (
	"fmt"
	gameloop "github.com/GlenKelley/go-glutil/gameloop"
	"github.com/go-gl/gl"
	glfw "github.com/go-gl/glfw3"
	"math"
	"github.com/neurocase/tankgame"
)

func main() {
	game := &Game{}
	err := gameloop.CreateWindow(640, 480, "daz gl test", false, game)
	fmt.Println(err)
}

type Game struct {
	Red float64
}

func (game *Game) Init(window *glfw.Window) {

}

func (game *Game) Draw(window *glfw.Window) {
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

func (game *Game) Reshape(window *glfw.Window, width, height int) {

}

func (game *Game) MouseClick(window *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {

if action == glfw.Press{

	switch button{

	case glfw.MouseButtonLeft:
		fmt.Println(tankgame.Hello())
	}
}

}

func (game *Game) MouseMove(window *glfw.Window, xpos float64, ypos float64) {

 fmt.Println("mouse move", xpos, ypos)

}

func (game *Game) KeyPress(window *glfw.Window, k glfw.Key, s int, action glfw.Action, mods glfw.ModifierKey) {
	fmt.Println("keypress", k)
	if action == glfw.Release {
		switch k {

		case glfw.KeyEscape:
			window.SetShouldClose(true)

		}
	}
}

func (game *Game) Scroll(window *glfw.Window, xoff float64, yoff float64) {

}

func (game *Game) Simulate(time gameloop.GameTime) {
	game.Red = math.Sin(time.Elapsed.Seconds())
}

func (game *Game) OnClose(window *glfw.Window) {

}

func (game *Game) IsIdle() bool {
	//if idle is true, the gameloop will
	//wait for user input before drawing
	return false
}

func (game *Game) NeedsRender() bool {
	//if render is false the game will not redraw the screen
	return true
}
