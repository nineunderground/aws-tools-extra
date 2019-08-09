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
  //"fmt"
  "github.com/gen2brain/raylib-go/raylib"
)

var backRect rl.Rectangle
var aboutMsg = "Game developed by\n<nineunderground_at_gmail_dot_com>\n\n\nBasically, it consists on a implementation of the game King's Up!"

/* Show ABOUT Scene */
func ShowAboutScene(width int32, height int32) {
  //fmt.Println("Show ABOUT Scene ")

  // Score board: 800x50
  // rl.DrawRectangle(0, 0, width, 50, rl.DarkGray)
  // rl.DrawText(" SCORE: 0", 0, 10, 15, rl.White)
  // rl.DrawText("LEVEL: 0", width-80, 10, 15, rl.White)

  rl.DrawText(aboutMsg, 0, 50, 25, rl.Black)

  // BACK
  btnw := int32(200)
  btnh := int32(80)
  currX := int32((width/2) - (btnw/2))
  currY := int32(height - btnh - 40)
  backRect = rl.NewRectangle(float32(currX), float32(currY), float32(btnw), float32(btnh))
  rl.DrawRectangleRec(backRect, rl.Blue)
  rl.DrawText("         BACK", currX, currY + (80/2) - 10, 20, rl.White)
}

/* Listen About Actions */
func ListenAboutActions(currentScene int) int {
  if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
    mousePos := rl.GetMousePosition()
    if rl.CheckCollisionPointRec(mousePos, backRect) {
      return SCN_TITLE
    }
  }
  return 0;
}
