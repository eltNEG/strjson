package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/atotto/clipboard"
)

func main() {
	args := os.Args[1]
	if !strings.HasSuffix(args, ".json") {
		fmt.Println("Arg is not a json file")
		return
	}
	jsonFile, err := os.Open(args)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer jsonFile.Close()
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	buffer := new(bytes.Buffer)
	err = json.Compact(buffer, byteValue)
	if err != nil {
		fmt.Println(err)
		return
	}
	clipboard.WriteAll(buffer.String())
}
