package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	cases := readVestigiumInput()
	for _, cs := range cases {
		t, r, c := analyseLatinMatrix(cs.matrix)
		fmt.Println(fmt.Sprintf("Case #%d: %d %d %d", cs.num, t, r, c))
	}
}

/*
Given a matrix that contains only integers between 1 and N, we want to compute its trace and check whether it is a natural
Latin square. To give some additional information, instead of simply telling us whether the matrix is a natural Latin
square or not, please compute the number of rows and the number of columns that contain repeated values.
*/

type testCaseVestigium struct {
	num    int
	matrix [][]int
}

func readVestigiumInput() []testCaseVestigium {
	cases := 0
	var err error

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		if cases, err = strconv.Atoi(scanner.Text()); err != nil {
			return nil
		}
	}

	output := make([]testCaseVestigium, cases)
	idx := 0
	size := 0
	for scanner.Scan() {
		if size, err = strconv.Atoi(scanner.Text()); err != nil {
			return nil
		}

		cs := testCaseVestigium{
			num:    idx + 1,
			matrix: make([][]int, size),
		}

		for i := 0; i < size; i++ {
			if scanner.Scan() {
				numbers := strings.Split(scanner.Text(), " ")
				cs.matrix[i] = make([]int, size)

				for j, sn := range numbers {
					if n, e := strconv.Atoi(sn); e == nil {
						cs.matrix[i][j] = n
					}
				}
			} else {
				return nil
			}
		}

		output[idx] = cs
		idx++
	}

	return output
}

func analyseLatinMatrix(matrix [][]int) (int, int, int) {
	l := len(matrix)
	if l == 1 {
		return matrix[0][0], 0, 0
	}

	trace := 0
	dupRows := 0
	dupColumns := 0
	for i := 0; i < l; i++ {
		trace += matrix[i][i]

		rowSeen := make(map[int]bool)
		colSeen := make(map[int]bool)
		sr := false
		sc := false
		for k := 0; k < l; k++ {
			if !sr && rowSeen[matrix[i][k]] {
				sr = true
			}
			rowSeen[matrix[i][k]] = true
			if !sc && colSeen[matrix[k][i]] {
				sc = true
			}
			colSeen[matrix[k][i]] = true

			if sr && sc {
				break
			}
		}
		if sr {
			dupRows++
		}
		if sc {
			dupColumns++
		}
	}

	return trace, dupRows, dupColumns
}
