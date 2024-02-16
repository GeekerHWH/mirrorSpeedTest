package main

import (
	"mirrorsTestTools/main/mode"
	"mirrorsTestTools/main/task"
)

func main() {

	// Check if OS is Debian 12 or Ubuntu 22.04
	if task.CheckOS() == 0 {
		return
	}

	mode.ParametersRun()
}
