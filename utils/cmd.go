package utils

import (
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func SetConsoleTitle(newTitle string) {
	newTitle = strings.Trim(newTitle, "\"")

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
