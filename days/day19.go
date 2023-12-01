package days

import (
	"bufio"
	"os"
)

func Day19() {
	var fileName string
	if os.Getenv("MODE") == "TEST" {
		fileName = "inputfiles/Day19Sample.txt"
	} else {
		fileName = "inputfiles/Day19.txt"
	}
	f, err := os.Open(fileName)
	Check(err)
	defer f.Close()
	var lavaDropletString []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lavaDropletString = append(lavaDropletString, scanner.Text())
	}
}
