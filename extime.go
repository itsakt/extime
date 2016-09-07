package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {

	args := os.Args[1:]
	var mArgs string
	if len(args) == 0 {
		return
	}

	for i := 1; i < len(args); i++ {
		mArgs = mArgs + string(args[i]) + " "
	}

	tStart := time.Now()
	out, err := exec.Command(args[0], strings.TrimSpace(mArgs)).Output()
	tEnd := time.Now()

	if err == nil {
		tDiff := tEnd.Sub(tStart)
		fmt.Println(string(out))
		fmt.Println()
		fmt.Println("Execution time: ", tDiff)
	} else {
		fmt.Println("Error executing command")
		fmt.Println(err)
	}

}
