package main

import (
	"fmt"
	"io"
	"os"
)

type logWriter struct {
}

func main() {
	myFile, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}

	lw := logWriter{}
	io.Copy(lw, myFile)
}

func (lw logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
