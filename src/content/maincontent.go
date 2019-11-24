/**
 * @Author: Iñaki Rodriguez <inaki>
 * @Date:   11-Aug-2019
 * @Project: Vylm.io
 * @Filename: maincontent.go
 * @Last modified by:   inaki
 * @Last modified time: 11-Aug-2019
 * @Copyright: Iñaki Rodriguez
 */
package content

import (
 "fmt"
 "fyne.io/fyne"
 "fyne.io/fyne/widget"
 "fyne.io/fyne/dialog"
)

var mainApp fyne.App
var mainWindow fyne.Window

func CreateMainLayout(uiapp fyne.App, uiWindow fyne.Window) *widget.Box {
  fmt.Println("************************************************")
  fmt.Println("*             MAIN CONTENT                     *")
  fmt.Println("************************************************")
  mainApp = uiapp
  mainWindow = uiWindow

  // MAIN LAYOUT
  selectLbl := "**********************\nPlease, select your action:\n**********************"
  labelDesc := widget.NewLabel(selectLbl)
  quitBtn := widget.NewButton("Quit", func() {
    quitDialog := dialog.NewConfirm("Exit","Do you quit?", exit, mainWindow)
    quitDialog.Show()
  })

  schedulerLayout := CreateSchedulerLayout()
  schedulerBtn := widget.NewButton("Scheduler", func() {
    setContent(schedulerLayout)
  })
  form := &widget.Form{}
  entry := widget.NewEntry()
  form.Append("S3Bucket", entry)
  sepADesc := widget.NewLabel("_______________")
  sepBDesc := widget.NewLabel("_______________")
  quitLayout := widget.NewHBox(sepADesc, quitBtn, sepBDesc)
  return widget.NewVBox(labelDesc, form, schedulerBtn, quitLayout)
}

func setContent(layout *widget.Box){
  mainWindow.SetContent(layout)
}

// EVENTS
func exit(isQuit bool){
  if isQuit {
    mainApp.Quit()
  }
}
