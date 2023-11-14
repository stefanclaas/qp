package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	qp "mime/quotedprintable"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: qp [-d|-e < infile > outfile]")
		os.Exit(1)
	}

	opt := os.Args[1]

	reader := bufio.NewReader(os.Stdin)

	switch opt {
	case "-d":
		for {
			input, err := reader.ReadString('\n')
			if err == io.EOF {
				break
			}
			decoded := qp.NewReader(strings.NewReader(input))
			buf := new(strings.Builder)
			io.Copy(buf, decoded)
			fmt.Print(buf.String())
		}
	case "-e":
		for {
			input, err := reader.ReadString('\n')
			if err == io.EOF {
				break
			}
			encoded := qp.NewWriter(os.Stdout)
			encoded.Write([]byte(input))
			encoded.Close()
		}
	default:
		fmt.Fprintln(os.Stderr, "Usage: qp [-d|-e < infile > outfile]")
		os.Exit(1)
	}
}

