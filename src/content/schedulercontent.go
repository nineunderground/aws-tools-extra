/**
 * @Author: Iñaki Rodriguez <inaki>
 * @Date:   11-Aug-2019
 * @Project: Vylm.io
 * @Filename: schedulercontent.go
 * @Last modified by:   inaki
 * @Last modified time: 11-Aug-2019
 * @Copyright: Iñaki Rodriguez
 */
package content

import (
  "fmt"
  "fyne.io/fyne/widget"
)

func CreateSchedulerLayout(uiapp fyne.App, uiWindow fyne.Window, mainLayout String) *widget.Box {
  fmt.Println("************************************************")
  fmt.Println("*             SCHEDULER CONTENT                *")
  fmt.Println("************************************************")
  mainApp = uiapp
  mainWindow = uiWindow

  // LAYOUT
  labelDesc := widget.NewLabel("Scheduler view")
  backToMainBtn := widget.NewButton("Back", func() {
    setContent(mainLayout)
  })

  return widget.NewVBox(labelDesc, backToMainBtn)
}
