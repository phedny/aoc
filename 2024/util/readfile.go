package util

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

var rFilename = regexp.MustCompile(`day(\d+).go`)

func GetDay() int {
	for i := 1; ; i++ {
		_, f, _, ok := runtime.Caller(i)
		if !ok {
			panic("no caller")
		}
		m := rFilename.FindStringSubmatch(f)
		if len(m) == 2 {
			i, err := strconv.Atoi(m[1])
			if err != nil {
				panic(err)
			}
			return i
		}
	}
}

func ReadFile() []byte {
	file := "real"
	if len(os.Args) > 1 {
		file = os.Args[1]
	}
	bytes, err := os.ReadFile(fmt.Sprintf("day%d/%s.txt", GetDay(), file))
	if err != nil {
		panic(err)
	}
	return bytes
}

func ReadLines() Lines {
	return strings.Split(string(ReadFile()), "\n")
}

func ReadByteMatrix() ByteMatrix {
	return bytes.Split(ReadFile(), []byte{'\n'})
}
