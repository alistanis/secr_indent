package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var (
	files string
	o     string
)

func init() {
	flag.StringVar(&files, "f", "", "Comma separated list of files to indent (will print them to stdout)")
	flag.StringVar(&o, "o", "", "Out file to write to.")
	flag.Parse()
}

func main() {
	if files == "" {
		fmt.Fprintln(os.Stderr, "Must provide -f")
		flag.Usage()
		os.Exit(-1)
	}

	paths := strings.Split(files, ",")
	buf := bytes.NewBuffer([]byte{})

	for _, p := range paths {

		buf.WriteString("*" + p + "*" + "\n\n")
		buf.WriteString("    ")
		data, err := ioutil.ReadFile(p)
		if err != nil {
			exitErr(err)
		}

		for _, b := range data {
			buf.WriteByte(b)
			if b == '\n' {
				buf.WriteString("    ")
			}
		}
		buf.WriteRune('\n')
	}
	if o != "" {
		err := ioutil.WriteFile(o, buf.Bytes(), 0664)
		if err != nil {
			exitErr(err)
		}
	} else {
		fmt.Println(buf.String())
	}

}

// Insert inserts insertData into data at the given start index
func insert(data, insertData []byte, start int) []byte {
	return append(data[:start], append(insertData, data[start:]...)...)
}

func exitErr(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(-1)
}
