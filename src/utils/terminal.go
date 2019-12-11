/**
 * @Author: Iñaki Rodriguez <inaki>
 * @Date:   18-Mar-2019
 * @Project: Vylm.io
 * @Filename: terminal.go
 * @Last modified by:   inaki
 * @Last modified time: 11-Aug-2019
 * @Copyright: Iñaki Rodriguez
 */
package utils

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
)

const QUIT = 0
const SSL_CERTIFICATE_VALIDITY = 1
const AWS_PROFILE = 2

/**
 * Print all available options
 */
func PrintInfo() {
	fmt.Println("************************************************")
	fmt.Println("*             AWS UTILS                        *")
	fmt.Println("************************************************")
	fmt.Println("*                                              *")
	fmt.Println("*  OPTIONS:                                    *")
	fmt.Println("*                                              *")
	fmt.Println("*  (1) CHECK ALL SSL CERTIFICATE VALIDITY      *")
	fmt.Println("*  (2) AWS PROFILE                             *")
	fmt.Println("*  (3) OTHER                                   *")
	fmt.Println("*                                              *")
	fmt.Println("*  (0) EXIT                                    *")
	fmt.Println("************************************************")
}

/**
 * Read the user input until a valid option is get
 */
func ReadOption() int {
	var optionInt int

	for true {
		_, err := fmt.Scanf("%d", &optionInt)
		if err != nil {
			log.Fatal(err)
		}
		if optionInt < 0 || optionInt > 3 {
			fmt.Print("NOT Valid option \n")
		} else {
			return optionInt
		}
	}
	return 0
}

/**
 * Clear bash console screen
 */
func Clear() {
	if runtime.GOOS == "darwin" { // OSX
		fmt.Print("Running linux comand... \n")
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			log.Fatalf("cmd.Run() failed with %s\n", err)
		}
	} else if runtime.GOOS == "windows" { // Windows
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
