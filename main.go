package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		panic(fmt.Sprintf("Wrong usage, expected 3 arguments, got %d", len(os.Args)))
	}

	file, err := os.Open(os.Args[1])
	defer file.Close()
	if err != nil {
		panic(err.Error())
	}

	output, err := os.Create(os.Args[2] + ".hpp")
	defer output.Close()
	if err != nil {
		panic(err.Error())
	}

	readBytes, err := ioutil.ReadFile(os.Args[2])
	if err != nil {
		panic(err.Error())
	}

	output.WriteString("#include <windows.h>\n\n inline BYTE " + os.Args[2] + "[] = {\n\t")
	for _, singleByte := range readBytes {
		output.WriteString(fmt.Sprintf("0x%02X, ", singleByte))
	}
	output.WriteString("\n}\n")
}
