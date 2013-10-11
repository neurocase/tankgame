package main





/*
 ---------------- SUPER BASIC RTS

STAGE 1:

Click to place red or green triangles, green triangles will automaticly fly directly into nearest red triangle and self destruct, 
both triangles will disappear. (press key to switch)

STAGE 2:

Yellow triangles will shoot at green triangles and try to shoot them before they hit red triangle.


STAGE 3:

Impliment some sort of path finding, generate paths with trace if possible, revert to waypoints if too difficult. 

STAGE 4:

Review possibilities of creating basic tower defence game or RTS



*/



import (
	"fmt"
	gameloop "github.com/GlenKelley/go-glutil/gameloop"
	"github.com/go-gl/gl"
	glfw "github.com/go-gl/glfw3"
//	"math"
	"github.com/neurocase/tankgame"

)




var en = tankgame.Entity{0,-1,-1, 1,"blue"}



func main() {
	game := &Game{}
	err := gameloop.CreateWindow(640, 480, "daz gl test", false, game)
	fmt.Println(err)
}

type Game struct {
//	Red float64
}

func (game *Game) Init(window *glfw.Window) {
	//Select the 'projection matrix'
	gl.MatrixMode(gl.PROJECTION)
	//Reset
	gl.LoadIdentity()
	//Scale everything down, to 1/10 scale
	gl.Scaled(0.1,0.1,0.1)

}

func (game *Game) Draw(window *glfw.Window) {
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
	gl.ClearColor(0.2, 0.2, 0.2, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	

	tankgame.DrawEntity(en)

}

func (game *Game) Reshape(window *glfw.Window, width, height int) {

}

func (game *Game) MouseClick(window *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {

if action == glfw.Press{

	switch button{

	case glfw.MouseButtonLeft:
		fmt.Println(tankgame.Hello())
		en.Rot += 0.1
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
	//game.Red = math.Sin(time.Elapsed.Seconds())
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
