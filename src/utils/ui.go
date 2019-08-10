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
  "github.com/schollz/progressbar"
  "time"
  "fyne.io/fyne/widget"
	"fyne.io/fyne/app"
  "../res"
  "fyne.io/fyne/dialog"
)

var mainApp = app.New()
//var mainDesc = "Please select one choice"
// var FYNE_THEME=light

// Generic
func PrintInfo() {
  //fmt.Println("Testing. \nTesting more...")
  fmt.Println("************************************************")
  fmt.Println("*             AWS UTILS                        *")
  fmt.Println("************************************************")
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

  // Main application obj
  //mainApp = app.New()
  mainApp.SetIcon(icon.GetLogo())

  // New window
	mainWindow := mainApp.NewWindow("AWS Tools by 9Underground!")

  // Set window content
  // LABELS
  selectLbl := "**********************\nPlease, select your action:\n**********************"
  labelDesc := widget.NewLabel(selectLbl)

  quitBtn := widget.NewButton("Quit", func() {
    //QuitApp(mainWindow)
    quitDialog := dialog.NewConfirm("Exit","Do you quit?", exit, mainWindow)
    quitDialog.Show()
  })

  // Set form content
  form := &widget.Form{}
  //form.Append("formItemA", widget.NewLabel("formLabelA"))
  entry := widget.NewEntry()
  form.Append("S3Bucket", entry)

  sepADesc := widget.NewLabel("_______________")
  sepBDesc := widget.NewLabel("_______________")
  quitLayout := widget.NewHBox(sepADesc, quitBtn, sepBDesc)

  // Set main content
  mainLayout := widget.NewVBox(labelDesc, form, quitLayout)
  mainWindow.SetContent(mainLayout)
  mainWindow.CenterOnScreen()
	mainWindow.ShowAndRun()
}

func exit(isQuit bool){
  if isQuit {
    mainApp.Quit()
  }
}
