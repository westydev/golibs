package utils

import (
	"os"
	"os/exec"
	"runtime"
)

func SetConsoleTitle(newTitle string) {
	cmd := exec.Command("cmd", "/c", "title", newTitle)
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Clear() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}
