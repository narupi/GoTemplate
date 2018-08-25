package tools

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

func ReadLine() string {
	sc.Scan()
	return sc.Text()
}

func ReadInt() int {
	i, e := strconv.Atoi(readStr())
	if e != nil {
		panic(e)
	}
	return i
}

func ReadFloat() float64 {
	f, e := strconv.ParseFloat(readStr(), 64)
	if e != nil {
		panic(e)
	}
	return f
}
