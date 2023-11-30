package main

// v1.0.0	(20191120/1/en)		-- ORIGINAL

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// MAIN
func main() {

	hostname, _ := os.Hostname()
	wdir, _ := os.Getwd()
	thisf, _ := os.Executable()
	thisp := filepath.Dir(thisf)

	fmt.Printf(
		"Hostname: %s\n"+
			"Working dir: %s\n"+
			"This program full path: %s\n"+
			"Just path: %s\n",
		hostname, wdir, thisf, thisp,
	)

	crt(80)
	cmd := exec.Command("ls", "-lHh")
	cmd.Stdout = os.Stdout // cmd.Stdout -> stdout
	cmd.Stderr = os.Stderr // cmd.Stderr -> stderr
	cmd.Run()

	crt(80)
	sysInfo()

} // end MAIN

func crt(n int) {
	line := fmt.Sprintf(strings.Repeat("-", n))
	fmt.Println(line)
}

func sysInfo() {
	infoLines := os.Environ()

	for _, line := range infoLines {
		if strings.Contains(line, "USER") {
			fmt.Println(line)
		}
	}

}
