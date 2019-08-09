/**
 * @Author: Iñaki Rodriguez <inaki>
 * @Date:   21-Mar-2019
 * @Project: Vylm.io
 * @Filename: board.go
 * @Last modified by:   inaki
 * @Last modified time: 22-Mar-2019
 * @Copyright: Iñaki Rodriguez
 */
package scn

import (
 //"../res"
 "github.com/gen2brain/raylib-go/raylib"
)

var floorZero rl.Rectangle
var floorOne rl.Rectangle
var floorTwo rl.Rectangle
var floorThree rl.Rectangle
var floorFour rl.Rectangle
var floorFive rl.Rectangle
var floorThrone rl.Rectangle

// var areImgLoaded bool = false
// var backgroundTex rl.Texture2D

 /* Show BOARD Scene */
func ShowBoardScene(width int32, height int32) {
  // if !areImgLoaded {
  //   backgroundTex = res.LoadImageLogo()
  //   areImgLoaded = true
  // }
 // Board is 800x700
 // Score board: 800x50
 rl.DrawRectangle(0, 0, width, 50, rl.DarkGray)
 rl.DrawText(" SCORE: 0", 0, 10, 15, rl.White)
 rl.DrawText("LEVEL: 0", width-80, 10, 15, rl.White)

 btnw := int32(550)
 btnh := int32(70)
 hmargin := int32(50)
 currX := (width/2)-(btnw/2)
 currY := int32(height - btnh - hmargin)

 floorZero = rl.NewRectangle(float32(currX), float32(currY), float32(btnw), float32(btnh))
 rl.DrawRectangleRec(floorZero, rl.Brown)

 currY = currY - btnh - 5
 floorOne = rl.NewRectangle(float32(currX), float32(currY), float32(btnw), float32(btnh))
 rl.DrawRectangleRec(floorOne, rl.Red)

 currY = currY - btnh - 5
 floorTwo = rl.NewRectangle(float32(currX), float32(currY), float32(btnw), float32(btnh))
 rl.DrawRectangleRec(floorTwo, rl.Yellow)

 btnw = btnw - 110
 currX = (width/2)-(btnw/2)
 currY = currY - btnh - 5
 floorThree = rl.NewRectangle(float32(currX), float32(currY), float32(btnw), float32(btnh))
 rl.DrawRectangleRec(floorThree, rl.Green)

 currY = currY - btnh - 5
 floorFour = rl.NewRectangle(float32(currX), float32(currY), float32(btnw), float32(btnh))
 rl.DrawRectangleRec(floorFour, rl.Orange)

 btnw = btnw - 120
 currX = (width/2)-(btnw/2)
 currY = currY - btnh - 5
 floorFive = rl.NewRectangle(float32(currX), float32(currY), float32(btnw), float32(btnh))
 rl.DrawRectangleRec(floorFive, rl.Blue)

 btnw = btnw - 170
 currX = (width/2)-(btnw/2)
 currY = currY - btnh - 5
 floorThrone = rl.NewRectangle(float32(currX), float32(currY), float32(btnw), float32(btnh))
 rl.DrawRectangleRec(floorThrone, rl.Gold)

 // Show scoreboard
 circleRadio := 20
 radius := float32(circleRadio)
 centerX := int32(20)
 centerY := int32(20 + 50)

 // First row with tiles
 totalScoreTiles := 1
 for totalScoreTiles < 21 {
   rl.DrawCircle(centerX, centerY, radius, rl.Black)
   centerX = centerX + int32(radius * 2)
   totalScoreTiles++
 }
 // n rows with two tiles
 totalScoreTiles = 1
 for totalScoreTiles < 15 {
   centerX = int32(circleRadio)
   centerY = centerY + int32(radius * 2)
   rl.DrawCircle(centerX, centerY, radius, rl.Black)
   centerX = width - int32(circleRadio)
   rl.DrawCircle(centerX, centerY, radius, rl.Black)
   totalScoreTiles++
 }
 // Last row with tiles
 centerX = int32(circleRadio)
 centerY = centerY + int32(radius * 2)
 totalScoreTiles = 1
 for totalScoreTiles < 21 {
   rl.DrawCircle(centerX, centerY, radius, rl.Black)
   centerX = centerX + int32(radius * 2)
   totalScoreTiles++
 }

 //rl.DrawTexture(backgroundTex, int32(0), int32(50), rl.White)

 // BACK
 // btnw := int32(200)
 // btnh := int32(80)
 // currX := int32((width/2) - (btnw/2))
 // currY := int32(height - btnh - 40)
 // backRect = rl.NewRectangle(float32(currX), float32(currY), float32(btnw), float32(btnh))
 // rl.DrawRectangleRec(backRect, rl.Blue)
 // rl.DrawText("         BACK", currX, currY + (80/2) - 10, 20, rl.White)
}

 /* Listen About Actions */
func ListenBoardActions(currentScene int) int {
 //   if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
 //     mousePos := rl.GetMousePosition()
 //     if rl.CheckCollisionPointRec(mousePos, backRect) {
 //       return SCN_TITLE
 //     }
 //   }
 return 0;
}
