package days

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func Day13() {
	var fileName string
	if os.Getenv("MODE") == "TEST" {
		fileName = "inputfiles/Day13Sample.txt"
	} else {
		fileName = "inputfiles/Day13.txt"
	}
	f, err := os.Open(fileName)
	Check(err)
	defer f.Close()

	var packets [][]string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var number string
		if len(scanner.Text()) > 0 {
			packet := scanner.Text()
			var appendArray []string
			for _, value := range packet {
				if value == '[' {
					if len(number) > 0 {
						appendArray = append(appendArray, number)
						number = ""
					}
					appendArray = append(appendArray, "[")
				} else if value == ']' {
					if len(number) > 0 {
						appendArray = append(appendArray, number)
						number = ""
					}
					appendArray = append(appendArray, "]")
				} else if value == ',' {
					if len(number) > 0 {
						appendArray = append(appendArray, number)
						number = ""
					}
				} else {
					number += string(value)
				}
			}
			packets = append(packets, appendArray)
		}
	}
	correctIndex := runPart1(packets)
	var sum int
	for _, value := range correctIndex {
		sum += value
	}
	fmt.Println("Day 13 Puzzle Solution:")
	fmt.Printf("Part 1 - Sum of Correct Packet Indexs: %d\n", sum)
	decoder1 := "[[2]]"
	decoder2 := "[[6]]"
	decoder1Array := strings.Split(decoder1, "")
	decoder2Array := strings.Split(decoder2, "")
	packets = append(packets, decoder1Array, decoder2Array)
	correctOrderMap := runPart2(packets)
	decoderIndexProduct := correctOrderMap[decoder1] * correctOrderMap[decoder2]
	fmt.Printf("Part 2 - Product of Decoder Packet Indexes: %d\n", decoderIndexProduct)
}

func runPart1(packets [][]string) []int {
	var correctIndex []int
	for i := 0; i < len(packets); i++ {
		var correctOrder bool
		if i%2 == 0 {
			first := make([]string, len(packets[i]))
			second := make([]string, len(packets[i+1]))
			copy(first, packets[i])
			copy(second, packets[i+1])
			first, second = prepPackets(first, second)
			correctOrder = isCorrectOrder(first, second)
		}
		if correctOrder {
			correctIndex = append(correctIndex, i/2+1)
		}
	}
	return correctIndex
}

func runPart2(packets [][]string) map[string]int {
	packetMap := make(map[string]int)
	for _, packet1 := range packets {
		joinedPacket := strings.Join(packet1, "")
		order := 1
		for _, packet2 := range packets {
			if !reflect.DeepEqual(packet1, packet2) {
				first := make([]string, len(packet1))
				second := make([]string, len(packet2))
				copy(first, packet1)
				copy(second, packet2)
				first, second = prepPackets(first, second)
				correctOrder := isCorrectOrder(first, second)
				if !correctOrder {
					order++
				}
			}
		}
		packetMap[joinedPacket] = order
	}
	return packetMap
}

func prepPackets(first, second []string) ([]string, []string) {
	for j := 0; j < len(first); j++ {
		if j > len(second)-1 {
			break
		} else if first[j] == second[j] {
			continue
		} else if first[j] == "[" && second[j] != "]" {
			_, err := strconv.Atoi(string(second[j]))
			if err == nil {
				additional := make([]string, len(second[j+1:]))
				copy(additional, second[j+1:])
				second = append(second[0:j], "[", second[j], "]")
				second = append(second, additional...)
			} else {
				additional := make([]string, len(second[j:]))
				copy(additional, second[j:])
				second = append(second[0:j], "[", "]")
				second = append(second, additional...)
			}
			continue
		} else if second[j] == "[" && first[j] != "]" {
			_, err := strconv.Atoi(string(first[j]))
			if err == nil {
				additional := make([]string, len(first[j+1:]))
				copy(additional, first[j+1:])
				first = append(first[0:j], "[", first[j], "]")
				first = append(first, additional...)
			} else {
				additional := make([]string, len(first[j:]))
				copy(additional, first[j:])
				first = append(first[0:j], "[", "]")
				first = append(first, additional...)
			}
		} else if first[j] == "]" || second[j] == "]" {
			break
		}
	}
	loops := len(first) - len(second)
	for j := 0; j < loops; j++ {
		second = append(second, " ")
	}
	return first, second
}
func isCorrectOrder(first, second []string) bool {
	var correctOrder bool
	for j, value := range first {
		value1, err1 := strconv.Atoi(string(value))
		value2, err2 := strconv.Atoi(string(second[j]))
		if err1 == nil && err2 == nil && value1 < value2 {
			correctOrder = true
			break
		} else if err1 == nil && err2 == nil && value1 > value2 {
			break
		} else if value == second[j] {
			continue
		} else if err1 == nil && err2 == nil && value1 > value2 {
			break
		} else if value == first[len(first)-1] && len(first) < len(second) {
			correctOrder = true
			break
		} else if value == "]" {
			correctOrder = true
			break
		} else if second[j] == "]" {
			break
		}
	}
	return correctOrder
}
