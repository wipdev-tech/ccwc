package counters

import (
	"bufio"
	"fmt"
)

func Count(r *bufio.Reader, splitFunc bufio.SplitFunc) int {
	n := 0
	scanner := bufio.NewScanner(r)
	scanner.Split(splitFunc)
	for scanner.Scan() {
		n++
	}
	return n
}

func Bytes(r *bufio.Reader) int {
	return Count(r, bufio.ScanBytes)
}

func Lines(r *bufio.Reader) int {
	return Count(r, bufio.ScanLines)
}

func FormatOutput(fileName string, ns ...int) string {
	lMax := len(fileName)
	for _, n := range ns {
		if l := len(fmt.Sprint(n)); l > lMax {
			lMax = l
		}
	}

	fmtStr := fmt.Sprintf("%%"+"%v"+"v", lMax)
	out := ""
	for _, n := range ns {
		out += fmt.Sprintf(fmtStr, n)
		out += " "
	}

	out += fmt.Sprintf(fmtStr, fileName)
	return out
}
