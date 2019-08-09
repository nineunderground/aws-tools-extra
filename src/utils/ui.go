/**
 * @Author: Iñaki Rodriguez <inaki>
 * @Date:   18-Mar-2019
 * @Project: Vylm.io
 * @Filename: mainUtils.go
 * @Last modified by:   inaki
 * @Last modified time: 21-Mar-2019
 * @Copyright: Iñaki Rodriguez
 */
package gui

import (
  "fmt"
  "github.com/schollz/progressbar"
  "time"
  "github.com/gen2brain/raylib-go/raylib"
  "../scenes"
)

var xPos = int32(0)
var yPos = int32(50)
var speed = int32(5)
var screenWidth int32
var screenHeight int32

func PrintInfo() {
  fmt.Println("Creating a CLI client for one time use. \nOne CLI command file is created uploaded files to an S3 bucket.\nPrinting GUI...")
}

func TestProgressBar() {
  bar := progressbar.New(100)
  for i := 0; i < 100; i++ {
      bar.Add(1)
      time.Sleep(10 * time.Millisecond)
  }
  fmt.Println("")
}

/**
More info at
  https://github.com/gen2brain/raylib-go/blob/master/raylib/raylib.go
  https://www.raylib.com/cheatsheet/cheatsheet.html
  https://godoc.org/github.com/gen2brain/raylib-go/raylib
**/
func ShowGui() {
  // Initialization
  screenWidth = int32(800)
  screenHeight = int32(700)
  rl.InitWindow(screenWidth, screenHeight, "STREGUO - Colour Memory!")
  rl.SetTargetFPS(60)
  //rl.ToggleFullscreen()
  //icon := rl.LoadImage("./res/logo.png")
  //texture := rl.LoadTextureFromImage(icon)

	for !rl.WindowShouldClose() {
    rl.BeginDrawing()
    rl.ClearBackground(rl.White)
    scn.ShowScene(screenWidth, screenHeight)
    rl.EndDrawing()
	}
	rl.CloseWindow()
}

func drawTiles() {
  // Score board
  rl.DrawRectangle(0, 0, 800, 50, rl.DarkGray)
  rl.DrawText(" SCORE: 0", 0, 10, 15, rl.White)
  rl.DrawText("LEVEL: 0", 800-80, 10, 15, rl.White)

  // Grid with 16 circles
  //rl.DrawRectangle(0, 50, 50, 50, rl.Blue)
  var xmargin = int32(50)
  var ymargin = int32(100)
  var size = float32(50)
  // 1 row
  rl.DrawCircle(xmargin, ymargin, size, rl.Blue);
  xmargin = xmargin + 100
  rl.DrawCircle(xmargin, ymargin, size, rl.Black);
  xmargin = xmargin + 100
  rl.DrawCircle(xmargin, ymargin, size, rl.Yellow);
  xmargin = xmargin + 100
  rl.DrawCircle(xmargin, ymargin, size, rl.Brown);

  // 2 row
  ymargin = ymargin + 100
  xmargin = 50
  rl.DrawCircle(xmargin, ymargin, size, rl.Green);
  xmargin = xmargin + 100
  rl.DrawCircle(xmargin, ymargin, size, rl.Maroon);
  xmargin = xmargin + 100
  rl.DrawCircle(xmargin, ymargin, size, rl.Magenta);
  xmargin = xmargin + 100
  rl.DrawCircle(xmargin, ymargin, size, rl.Beige);

  // 3 row
  ymargin = ymargin + 100
  xmargin = 50
  rl.DrawCircle(xmargin, ymargin, size, rl.DarkGray);
  xmargin = xmargin + 100
  rl.DrawCircle(xmargin, ymargin, size, rl.Gray);
  xmargin = xmargin + 100
  rl.DrawCircle(xmargin, ymargin, size, rl.NewColor(205, 0, 155, 255));
  xmargin = xmargin + 100
  rl.DrawCircle(xmargin, ymargin, size, rl.NewColor(95, 30, 0, 255));

  // 4 row
  ymargin = ymargin + 100
  xmargin = 50
  rl.DrawCircle(xmargin, ymargin, size, rl.SkyBlue);
  xmargin = xmargin + 100
  rl.DrawCircle(xmargin, ymargin, size, rl.DarkPurple);
  xmargin = xmargin + 100
  rl.DrawCircle(xmargin, ymargin, size, rl.Orange);
  xmargin = xmargin + 100
  rl.DrawCircle(xmargin, ymargin, size, rl.Pink);
}

func ListenEvents() {
    if rl.IsKeyDown(rl.KeyUp) {
      yPos = yPos - speed //UP
    } else if rl.IsKeyDown(rl.KeyDown) {
      yPos = yPos + speed //DOWN
    } else if rl.IsKeyDown(rl.KeyLeft) {
      xPos = xPos - speed //LEFT
    } else if rl.IsKeyDown(rl.KeyRight) {
      xPos = xPos + speed //RIGHT
    }
}
