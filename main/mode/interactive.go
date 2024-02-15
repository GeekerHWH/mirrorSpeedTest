package mode

import (
	"fmt"
	"mirrorsTestTools/main/task"
)

const initNumURLs = 10

func Interactive() {
	// Use cfginput.ReadFile() to read the configuration file and convert it to a slice.
	// 调用cfginput.ReadFile()函数读取配置文件，并将其转换为切片
	mirrorNames, mirrorURLs := task.ReadFile("urls.json")
	// fmt.Println(mirrorNames, mirrorURLs)

	fmt.Println("Welcome to use MirrorSpeedTesting Tool")
	fmt.Println("1. Batch custom mirror site URLs. 2. Batch selection of mirror sites in the library. 3. Exit")

	var choice int
	fmt.Scanln(&choice)

	switch choice {

	case 1:
		fmt.Print("Please enter the URL of the mirror site to be tested: (Default:Tsinghua)\n")

		task.InputURL(&mirrorNames, &mirrorURLs)

		task.Test(mirrorNames, mirrorURLs, initNumURLs)

	case 2:
		fmt.Println("Please select the mirror sites to be tested (separated by spaces), 0 means select all:")
		for i, name := range mirrorNames {
			fmt.Printf("%d. %s\n", i+1, name)
		}

		task.SelectMirror(&mirrorNames, &mirrorURLs)

		task.Test(mirrorNames, mirrorURLs, initNumURLs)

	case 3:
		return
	}
}
