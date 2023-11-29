package main

// v1.0.0	(20191119/1/d)	-- ORIGINAL

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	myos := runtime.GOOS
	fmt.Printf("Os: %s\n", myos)
	fmt.Printf("Number of CPU cores: %d\n", runtime.NumCPU())
	fmt.Printf("The root of GO: %s\n", runtime.GOROOT())
	fmt.Printf("The GO version: %s\n", runtime.Version())

	execFilePath, _ := os.Executable()
	// execdir, _ := filepath.Dir(os.Executable())
	fmt.Printf("Executable file name: %s\n", path.Base(execFilePath))
	fmt.Printf("Executable dir: %s\n", filepath.Dir(execFilePath))

	fmt.Println(strings.Repeat("-", 80))
	var cmd *exec.Cmd
	if myos == "windows" {
		fmt.Printf("Executing: %s\n", `exec.Command("DIR")`)
		cmd = exec.Command("DIR")
	} else {
		fmt.Printf("Executing: %s\n", `exec.Command("ls", "-lFh")`)
		cmd = exec.Command("ls", "-lFh")

	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

}
