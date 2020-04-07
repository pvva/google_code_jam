package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	cases := readNestingDepthInput()
	for _, cs := range cases {
		fmt.Println(fmt.Sprintf("Case #%d: %s", cs.num, nestingDepth(cs.input)))
	}
}

/*
Given a string of digits S, insert a minimum number of opening and closing parentheses into it such that the resulting
string is balanced and each digit d is inside exactly d pairs of matching parentheses.

For example, in the following strings, all digits match their nesting depth: 0((2)1), (((3))1(2)), ((((4)))),
((2))((2))(1). The first three strings have minimum length among those that have the same digits in the same order,
but the last one does not since ((22)1) also has the digits 221 and is shorter.
*/
type testCaseNestingDepth struct {
	num   int
	input string
}

func readNestingDepthInput() []testCaseNestingDepth {
	cases := 0
	var err error

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		if cases, err = strconv.Atoi(scanner.Text()); err != nil {
			return nil
		}
	}

	output := make([]testCaseNestingDepth, cases)
	for i := 0; i < cases; i++ {
		if scanner.Scan() {
			output[i] = testCaseNestingDepth{
				num:   i + 1,
				input: scanner.Text(),
			}
		} else {
			return nil
		}
	}

	return output
}

func nestingDepth(s string) string {
	sb := bytes.Buffer{}
	opened := 0

	for i := 0; i < len(s); i++ {
		reqBraces := int(s[i] - '0')
		for opened < reqBraces {
			sb.WriteString("(")
			opened++
		}
		for opened > reqBraces {
			sb.WriteString(")")
			opened--
		}
		sb.WriteString(s[i : i+1])
	}
	for opened > 0 {
		sb.WriteString(")")
		opened--
	}

	return sb.String()
}
