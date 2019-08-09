/**
 * @Author: Iñaki Rodriguez <inaki>
 * @Date:   20-Mar-2019
 * @Project: Vylm.io
 * @Filename: main.go
 * @Last modified by:   inaki
 * @Last modified time: 21-Mar-2019
 * @Copyright: Iñaki Rodriguez
 */
package scn

import (
  // "fmt"
  //"github.com/gen2brain/raylib-go/raylib"
)

// Global scenes
const SCN_TITLE = 1
const SCN_OPTIONS = 2
const SCN_BOARD = 3
const SCN_ABOUT = 4

// Current scenes
var currScene = SCN_TITLE
var newScene = 0;

/* Show Scene */
func ShowScene(width int32, height int32) {
  if currScene == SCN_TITLE {
    ShowTitleScene(width, height)
    newScene = ListenTitleActions(currScene)
  } else if currScene == SCN_OPTIONS {
    // TODO ...
  } else if currScene == SCN_BOARD {
    ShowBoardScene(width, height)
    newScene = ListenBoardActions(currScene)
  } else if currScene == SCN_ABOUT {
    ShowAboutScene(width, height)
    newScene = ListenAboutActions(currScene)
  }
  // Refresh the scene to show if needed
  if newScene != 0 {
    currScene = newScene
  }
}
