/**
 * @Author: Iñaki Rodriguez <inaki>
 * @Date:   10-Aug-2019
 * @Project: Vylm.io
 * @Filename: icon.go
 * @Last modified by:   inaki
 * @Last modified time: 11-Aug-2019
 * @Copyright: Iñaki Rodriguez
 */
package icon

import (
  "fyne.io/fyne"
  "os"
  "log"
)

func GetLogo() *fyne.StaticResource {
  readerFile, err := os.Open("src/res/logo.png")
  if err != nil {
    log.Fatal(err)
  }
  data := make([]byte, 100)
  //count, err := readerFile.Read(data)
  readerFile.Read(data)

  return fyne.NewStaticResource("test", data)
}
