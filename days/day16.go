package days

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Day16() {
	var fileName string
	if os.Getenv("MODE") == "TEST" {
		fileName = "inputfiles/Day16Sample.txt"
	} else {
		fileName = "inputfiles/Day16.txt"
	}
	f, err := os.Open(fileName)
	Check(err)
	defer f.Close()

	valveMap := make(map[string]*Valve)
	scanner := bufio.NewScanner(f)
	var valveRates []int
	for scanner.Scan() {
		valves := strings.Split(scanner.Text(), " ")
		rate, err := strconv.Atoi(strings.FieldsFunc(valves[4], SplitOnNumber)[0])
		Check(err)
		var open bool
		if rate == 0 {
			open = true
		}
		vToAdd := Valve{Name: valves[1], Rate: rate, Open: open}
		for _, v := range valves[9:] {
			vToAdd.Tunnels = append(vToAdd.Tunnels, strings.ReplaceAll(v, ",", ""))
		}
		valveMap[vToAdd.Name] = &vToAdd
		valveRates = append(valveRates, rate)
	}
	sort.Ints(valveRates)
	for _, valve := range valveMap {
		for _, name := range valve.Tunnels {
			valve.ConnectedValves = append(valve.ConnectedValves, valveMap[name])
		}
	}
	fmt.Println("Day 16 Puzzle Solution:")
}
