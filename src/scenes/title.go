/**
 * @Author: Iñaki Rodriguez <inaki>
 * @Date:   20-Mar-2019
 * @Project: Vylm.io
 * @Filename: title.go
 * @Last modified by:   inaki
 * @Last modified time: 21-Mar-2019
 * @Copyright: Iñaki Rodriguez
 */
package scn

import (
  "fmt"
  "os"
  "github.com/gen2brain/raylib-go/raylib"
)

var abtRect rl.Rectangle
var quitRect rl.Rectangle
var boardRect rl.Rectangle

/* Show TITLE Scene */
func ShowTitleScene(width int32, height int32) {

  // Score board: 800x50
  rl.DrawRectangle(0, 0, width, 50, rl.DarkGray)
  rl.DrawText(" SCORE: 0", 0, 10, 15, rl.White)
  rl.DrawText("LEVEL: 0", width-80, 10, 15, rl.White)

  // Draw buttons
  btnw := int32(200)
  btnh := int32(80)
  hmargin := int32(60)
  currX := (width/2)-(btnw/2)
  currY := int32(height - 550)

  // PLAY
  // rl.Fade(rl.Blue, 0.5)
  boardRect = rl.NewRectangle(float32(currX), float32(currY), float32(btnw), float32(btnh))
  rl.DrawRectangleRec(boardRect, rl.Blue)
  rl.DrawText("         PLAY", currX, currY + (btnh/2) - 10, 20, rl.Yellow)

  // ABOUT
  currY = currY + btnh + hmargin
  abtRect = rl.NewRectangle(float32(currX), float32(currY), float32(btnw), float32(btnh))
  rl.DrawRectangleRec(abtRect, rl.Blue)
  rl.DrawText("        ABOUT", currX, currY + (btnh/2) - 10, 20, rl.Yellow)

  // QUIT
  currY = currY + btnh + hmargin
  quitRect = rl.NewRectangle(float32(currX), float32(currY), float32(btnw), float32(btnh))
  rl.DrawRectangleRec(quitRect, rl.Blue)
  rl.DrawText("         QUIT", currX, currY + (btnh/2) - 10, 20, rl.Yellow)
}

func ListenTitleActions(currentScene int) int {
  if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
    mousePos := rl.GetMousePosition()
    if rl.CheckCollisionPointRec(mousePos, quitRect) {
      fmt.Println("Quiting")
      os.Exit(0)
    } else if rl.CheckCollisionPointRec(mousePos, abtRect) {
      return SCN_ABOUT
    } else if rl.CheckCollisionPointRec(mousePos, boardRect) {
      return SCN_BOARD
    }
  }
  return 0;
}
