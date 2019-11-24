/**
 * @Author: Iñaki Rodriguez <inaki>
 * @Date:   18-Mar-2019
 * @Project: Vylm.io
 * @Filename: mainUtils.go
 * @Last modified by:   inaki
 * @Last modified time: 11-Aug-2019
 * @Copyright: Iñaki Rodriguez
 */
package ui

import (
  "fmt"
  "fyne.io/fyne/app"
  "../res"
  "../content"
)

var mainApp = app.New()
var mainWindow = mainApp.NewWindow("AWS Tools by 9Underground!")
//var mainDesc = "Please select one choice"
// var FYNE_THEME=light

// Generic
func PrintInfo() {
  //fmt.Println("Testing. \nTesting more...")
  fmt.Println("************************************************")
  fmt.Println("*             AWS UTILS                        *")
  fmt.Println("************************************************")
}

// GUI
// More info at: https://github.com/fyne-io/fyne
// https://github.com/fyne-io/examples/
func ShowMainGUI() {
  fmt.Println("Showing GUI")
  mainApp.SetIcon(icon.GetLogo())

  // LAYOUTS:
  mainLayout := content.CreateMainLayout(mainApp, mainWindow)

  // Set MAIN LAYOUT
  mainWindow.SetContent(mainLayout)
  mainWindow.CenterOnScreen()
	mainWindow.ShowAndRun()
}
