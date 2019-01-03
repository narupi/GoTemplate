package template

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func SetScSplit() { sc.Split(bufio.ScanWords) }

func ReadStr() string {
	sc.Scan()
	return sc.Text()
}

func ReadInt() int {
	i, e := strconv.Atoi(ReadStr())
	if e != nil {
		panic(e)
	}
	return i
}

func ReadFloat() float64 {
	f, e := strconv.ParseFloat(ReadStr(), 64)
	if e != nil {
		panic(e)
	}
	return f
}

func SplitStr(str string, s string) {
	return strings.Split(str, s)
}

func DeleteElement(data []int, idx int) []int {
	data = append(data[:idx], a[idx+1:]...)
	return data
}
