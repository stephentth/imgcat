package main

import (
	"os/exec"
	"strconv"
	"strings"
)

func GetTerminalSize() (int, int, error) {
	var col, row int
	cmd := exec.Command("sh", "-c", "tput cols; tput lines")
	stdout, err := cmd.CombinedOutput()
	if err != nil {
		return 0, 0, err
	}

	stdoutString := strings.Replace(string(stdout), "\n", " ", -1)
	outputs := strings.Split(stdoutString, " ")

	col, err = strconv.Atoi(outputs[0])
	if err != nil {
		return 0, 0, err
	}

	row, err = strconv.Atoi(outputs[1])
	if err != nil {
		return 0, 0, err
	}
	return col, row, nil
}
