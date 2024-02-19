package mode

import (
	"fmt"
	"mirrorsTestTools/main/task"
)

func Interactive(interactiveCountry string) {
	// Use cfginput.ReadFile() to read the configuration file and convert it to a slice.
	mirrorNames, mirrorURLs, err := task.ReadFile("urls.json", interactiveCountry)
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(mirrorNames, mirrorURLs)

	fmt.Println("Welcome to use MirrorSpeedTesting Tool")
	fmt.Println("1. Batch selection of mirror sites in the library. 2. Batch custom mirror site URLs. 3. Exit")

	var choice int
	fmt.Scanln(&choice)

	switch choice {

	case 1:
		fmt.Println("Please select the mirror sites to be tested (separated by spaces), 0 means select all:")
		for i, name := range mirrorNames {
			fmt.Printf("%d. %s\n", i+1, name)
		}

		selectNames, selectURLs := task.SelectMirror(&mirrorNames, &mirrorURLs)

		task.Test(selectNames, selectURLs)

	case 2:
		fmt.Print("Please enter the URL of the mirror site to be tested: (Default:Tsinghua)\n")

		URLs := task.InputURL()

		task.Test(URLs, URLs)

	case 3:
		return
	}
}
