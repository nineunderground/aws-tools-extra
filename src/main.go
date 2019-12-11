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
	"fmt"
	"os"
	"time"

	utils "./utils"
	awsservices "./utils/services"
	"github.com/schollz/progressbar"
)

func main() {
	testProgressBar()
	utils.PrintInfo()
	option := utils.ReadOption()
	utils.Clear()

	switch option {
	case utils.AWS_PROFILE:
		fmt.Println("Specify your profile:")
	case utils.SSL_CERTIFICATE_VALIDITY:
		awsservices.DoAction()
	case utils.QUIT:
		os.Exit(0)
	}

	os.Exit(0)
}

func testProgressBar() {
	bar := progressbar.New(100)
	for i := 0; i < 100; i++ {
		bar.Add(1)
		time.Sleep(5 * time.Millisecond)
	}
	fmt.Println("")
}
