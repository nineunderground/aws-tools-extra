/**
 * @Author: Iñaki Rodriguez <inaki>
 * @Date:   18-Mar-2019
 * @Project: Vylm.io
 * @Filename: createClient.go
 * @Last modified by:   inaki
 * @Last modified time: 21-Mar-2019
 * @Copyright: Iñaki Rodriguez
 */
package main

import (
  "os"
  "./utils"
)

func main() {
  gui.PrintInfo()
  gui.TestProgressBar()
  gui.ShowGui()
  os.Exit(0)
}
