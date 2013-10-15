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

	"math"

)



/*
x := en.Xpos
y := en.Ypos
r := en.Rot
s := en.Size
col := en.Colour
*/
var en = tankgame.Entity{0,-1,-1, 1,"blue"}
var grd = tankgame.Entity{-10,-10,0,0.1,"grey"}
var active = tankgame.Entity{0, 0,1, 1,"blue"}
var grnent = tankgame.Entity{0,float64(mintx),float64(minty), 0.3,"green"}


var entary = []tankgame.Entity{}
var entaryg = []tankgame.Entity{}



var mousex = 0.0
var mousey = 0.0
var winHeight = 480.0
var winWidth = 640.0
var mode = 0
var mintx int
var minty int
var defaultrot = 0.0


var grnx = 0.0
var grny = 0.0

var vx = 0.0
var vy = 0.0

func switchMode(){		
	if mode < 1{
		mode++
	}else{
		mode =0
	}
}

func main() {
	game := &Game{}
	err := gameloop.CreateWindow(int(winWidth), int(winHeight), "daz gl test", false, game)
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




















	/*a := int(mousex)
	b := int(mousey)*/


	vx = (mousex * 2 / winWidth - 1)*10
	//fmt.Println(mousex, "* 2 /", winWidth, "-1")
	vy = (1 - mousey * 2 / winHeight)*10
	//fmt.Println(mousey, "* 2 /", winHeight, "-1")

	mintx = int(vx)
	minty = int(vy)

	for j := -10; j < 11; j++{




		grd.Xpos = float64(j)
		for y := -10; y < 11; y++{


			grd.Ypos = float64(y)
			tankgame.DrawEntity(grd)
		}
	}

	
	for i := 0; i < len(entary); i++{
		//entary[i].Rot = 

		hx := grnx - entary[i].Xpos 
		hy := grny - entary[i].Ypos
		hhyp := (hx * hx) + (hy * hy)
		hhyp = math.Sqrt(hhyp)
		hsin := hy / hhyp
		radian := math.Asin(hsin)
		degree := radian*180/math.Pi

		if grnx < entary[i].Xpos{
		degree = 0 - degree -90
		}else{
		degree = degree + 90	
		}
		if degree > 360{
			degree -= 360
		}else if degree < 0{
			degree += 360
		}

		entary[i].Rot = degree



		fmt.Println(entary[i].Rot)


		tankgame.DrawEntity(entary[i])
	}
		tankgame.DrawEntity(grnent)


	tankgame.DrawEntity(active)



	//tankgame.DrawEntity(en)

	//var ten = tankgame.Entity{0,-0,0, 10,"blue"}
	//tankgame.DrawEntity(ten)
}

func (game *Game) Reshape(window *glfw.Window, width, height int) {
winHeight = float64(height)
winWidth = float64(width)
fmt.Println("reshape does not work", float64(width), float64(height))
}

func (game *Game) MouseClick(window *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {
	if action == glfw.Press{

		switch button{

		case glfw.MouseButtonLeft:


			if mode == 0{

				ena := tankgame.Entity{0,float64(mintx),float64(minty), 0.3,"red"}
				fmt.Println("mouse move", vx, vy)
				entary = append(entary,ena)
			}

			if mode == 1{
				grnent = tankgame.Entity{0,float64(mintx),float64(minty), 0.3,"green"}
				grnx = float64(mintx)
				grny = float64(minty)
			}


		//	en.Rot += 0.1
		}
	}
}



func (game *Game) MouseMove(window *glfw.Window, xpos float64, ypos float64) {
	mousex = xpos
	mousey = ypos

	active = tankgame.Entity{0, float64(mintx),float64(minty), 0.3,"blue"}
	fmt.Println(mintx,minty)


}

func (game *Game) KeyPress(window *glfw.Window, k glfw.Key, s int, action glfw.Action, mods glfw.ModifierKey) {
	fmt.Println("keypress", k)
	if action == glfw.Release {
		switch k {

		case 77:			
			switchMode()
			fmt.Println(mode, "mode")
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
