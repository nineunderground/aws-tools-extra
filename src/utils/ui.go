/**
 * @Author: Iñaki Rodriguez <inaki>
 * @Date:   18-Mar-2019
 * @Project: Vylm.io
 * @Filename: mainUtils.go
 * @Last modified by:   inaki
 * @Last modified time: 10-Aug-2019
 * @Copyright: Iñaki Rodriguez
 */
package ui

import (
  "fmt"
  "github.com/schollz/progressbar"
  "time"
  "fyne.io/fyne/widget"
	"fyne.io/fyne/app"
)

var mainDec = "Please select one choice"

// Generic
func PrintInfo() {
  fmt.Println("Testing. \nTesting more...")
}

func TestProgressBar() {
  bar := progressbar.New(100)
  for i := 0; i < 100; i++ {
      bar.Add(1)
      time.Sleep(5 * time.Millisecond)
  }
  fmt.Println("")
}

// GUI
// More info at: https://github.com/fyne-io/fyne
// https://github.com/fyne-io/examples/
func ShowMainGUI() {
  fmt.Println("Showing GUI")
  app := app.New()
	w := app.NewWindow("AWS Tools by 9Underground!")
	w.SetContent(widget.NewVBox(
		widget.NewLabel(mainDec),
		widget.NewButton("Quit", func() {
			app.Quit()
		}),
	))
	w.ShowAndRun()
}
