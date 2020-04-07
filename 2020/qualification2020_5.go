package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	cases := readIndiciumInput()
	for _, cs := range cases {
		solution := produceValidLatinMatrix(cs.size, cs.trace)
		if solution == nil {
			fmt.Printf("Case #%d: IMPOSSIBLE\n", cs.num)
		} else {
			fmt.Printf("Case #%d: POSSIBLE\n", cs.num)
			printMatrix(solution)
		}
	}
}

type testCaseIndicium struct {
	num   int
	size  int
	trace int
}

func readIndiciumInput() []testCaseIndicium {
	cases := 0
	var err error

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		if cases, err = strconv.Atoi(scanner.Text()); err != nil {
			return nil
		}
	}

	output := make([]testCaseIndicium, cases)
	for i := 0; i < cases; i++ {
		cs := testCaseIndicium{
			num: i + 1,
		}
		if scanner.Scan() {
			numbers := strings.Split(scanner.Text(), " ")

			cs.size, _ = strconv.Atoi(numbers[0])
			cs.trace, _ = strconv.Atoi(numbers[1])
		} else {
			return nil
		}

		output[i] = cs
	}

	return output
}

func produceValidLatinMatrix(size, trace int) [][]int {
	diag := createTrace(size, trace)
	if diag == nil {
		return nil
	}

	m := make([][]int, size)
	for i := 0; i < size; i++ {
		m[i] = make([]int, size)
	}

	for i := 0; i < size; i++ {
		m[i][i] = diag[i]
	}

	possibleValues := make([][]map[int]bool, size)
	for i := 0; i < size; i++ {
		possibleValues[i] = make([]map[int]bool, size)
		for j := 0; j < size; j++ {
			if i != j {
				possibleValues[i][j] = make(map[int]bool)

				for k := 1; k <= size; k++ {
					if k != m[i][i] && k != m[j][j] {
						possibleValues[i][j][k] = true
					}
				}
			}
		}
	}

	if trySolve(m, possibleValues, size, 0, 1) {
		return m
	}

	return nil
}

func trySolve(m [][]int, possibleValues [][]map[int]bool, size, row, col int) bool {
	if row == size-1 && col == size-1 {
		return true
	}

	if m[row][col] != 0 {
		col++
		if col == size {
			row++
			col = 0
		}
	}

	if len(possibleValues[row][col]) == 0 {
		if m[row][col] == 0 {
			return false
		}
	}

	nCol := col + 1
	nRow := row
	if nCol == size {
		nRow++
		nCol = 0
	}

	for v := range possibleValues[row][col] {
		if isValidValue(m, row, col, v) {
			delete(possibleValues[row][col], v)
			m[row][col] = v

			if trySolve(m, possibleValues, size, nRow, nCol) {
				return true
			}

			possibleValues[row][col][v] = true
			m[row][col] = 0
		}
	}

	return false
}

func isValidValue(m [][]int, row, col int, value int) bool {
	for i := 0; i < len(m); i++ {
		if m[row][i] == value || m[i][col] == value {
			return false
		}
	}

	return true
}

func createTrace(size, trace int) []int {
	if size > trace || size*size < trace {
		return nil
	}
	d := make([]int, size)
	for i := 0; i < size; i++ {
		d[i] = 1
	}
	trace -= size

	for trace > 0 {
		d[nextTraceIdx(d, size)]++
		trace--
	}

	return d
}

func nextTraceIdx(d []int, limit int) int {
	idx := len(d) - 1

	for idx > 0 {
		if d[idx] == limit {
			idx--
			continue
		}
		if d[idx]-d[idx-1] == 1 {
			idx--
		}
		break
	}

	return idx
}

func printMatrix(matrix [][]int) {
	for _, sm := range matrix {
		d := ""
		for _, v := range sm {
			fmt.Printf("%s%d", d, v)
			d = " "
		}
		fmt.Println()
	}
}
