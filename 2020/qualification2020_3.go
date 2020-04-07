package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const solutionImpossible = "IMPOSSIBLE"

func main() {
	cases := readParentingPartneringInput()
	for _, cs := range cases {
		fmt.Println(fmt.Sprintf("Case #%d: %s", cs.num, assignParentingPartners(cs)))
	}
}

type testCaseParentingPartnering struct {
	num       int
	intervals [][2]int
}

func readParentingPartneringInput() []testCaseParentingPartnering {
	cases := 0
	var err error

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		if cases, err = strconv.Atoi(scanner.Text()); err != nil {
			return nil
		}
	}

	output := make([]testCaseParentingPartnering, cases)
	idx := 0
	size := 0
	for scanner.Scan() {
		if size, err = strconv.Atoi(scanner.Text()); err != nil {
			return nil
		}

		cs := testCaseParentingPartnering{
			num:       idx + 1,
			intervals: make([][2]int, size),
		}

		for i := 0; i < size; i++ {
			if scanner.Scan() {
				numbers := strings.Split(scanner.Text(), " ")

				cs.intervals[i][0], _ = strconv.Atoi(numbers[0])
				cs.intervals[i][1], _ = strconv.Atoi(numbers[1])
			} else {
				return nil
			}
		}

		output[idx] = cs
		idx++
	}

	return output
}

func assignParentingPartners(tc testCaseParentingPartnering) string {
	result := make([]byte, len(tc.intervals))

	// id => set of connected nodes
	graph := make([]map[int]bool, len(tc.intervals))
	for i, interval := range tc.intervals {
		graph[i] = make(map[int]bool)
		for j, tInterval := range tc.intervals {
			if interval[1] <= tInterval[0] || interval[0] >= tInterval[1] || i == j {
				continue
			}
			graph[i][j] = true
		}
	}

	var currentOcc byte = 'C'
	nodes := []int{}

	for nodeId := range graph {
		if result[nodeId] == 0 && len(nodes) == 0 {
			nodes = append(nodes, nodeId)
			result[nodeId] = currentOcc

			if currentOcc == 'C' {
				currentOcc = 'J'
			} else {
				currentOcc = 'C'
			}
		}

		for len(nodes) > 0 {
			nNodes := []int{}

			for _, id := range nodes {
				for n := range graph[id] {
					if result[n] != 0 {
						if result[n] != currentOcc {
							return solutionImpossible
						}
					} else {
						result[n] = currentOcc
						nNodes = append(nNodes, n)
					}
				}
			}

			if currentOcc == 'C' {
				currentOcc = 'J'
			} else {
				currentOcc = 'C'
			}

			nodes = nNodes
		}
	}

	return string(result)
}
