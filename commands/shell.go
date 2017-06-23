package commands

import (
	"bufio"
	"fmt"
	"os/exec"

	"github.com/fatih/color"
)

//ShellCommand ...
func ShellCommand(command string, path string) {
	cmd := exec.Command("bash", "-c", command)
	if path != "" {
		cmd.Dir = path
	}
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()

	scanner := bufio.NewScanner(stdout)
	errScanner := bufio.NewScanner(stderr)

	go func() {
		for scanner.Scan() {
			fmt.Printf("%s\n", scanner.Text())
		}
	}()
	go func() {
		for errScanner.Scan() {
			color.Red("[%s]%s\n", path, errScanner.Text())
		}
	}()
	err := cmd.Start()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

}