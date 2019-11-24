/**
 * @Author: Iñaki Rodriguez <inaki>
 * @Date:   18-Mar-2019
 * @Project: Vylm.io
 * @Filename: createClient.go
 * @Last modified by:   inaki
 * @Last modified time: 11-Aug-2019
 * @Copyright: Iñaki Rodriguez
 */
package main

import (
  //"os"
  "github.com/schollz/progressbar"
  "fmt"
  "time"
  "./utils"
)

func main() {
  ui.PrintInfo()
  testProgressBar()
  ui.ShowMainGUI()
  //os.Exit(0)
}

func testProgressBar() {
  bar := progressbar.New(100)
  for i := 0; i < 100; i++ {
      bar.Add(1)
      time.Sleep(5 * time.Millisecond)
  }
  fmt.Println("")
}
