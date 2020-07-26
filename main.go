package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

var limit, offset int
var from, to string

func init() {
	flag.StringVar(&from, "from", "", "path/to/source")
	flag.StringVar(&to, "to", "", "path/to/destination")
	flag.IntVar(&limit, "limit", 0, "The number of bytes to be written")
	flag.IntVar(&offset, "offset",0, "The offset of bytes to start read src file from")
}

func checkFlags() {
	if from == "" {
		fmt.Printf("The source path has not been set \n")
		os.Exit(1)
	}

	if to == "" {
		fmt.Printf("The destination path has not been set \n")
		os.Exit(1)
	}

	if limit < 1 {
		fmt.Printf("The limit should be positive number \n")
		os.Exit(1)
	}

	if offset < 0 {
		fmt.Printf("The offset should be equal or more than 0 \n")
		os.Exit(1)
	}
}

func main()  {
	flag.Parse()
	checkFlags()

	src, err := os.Open(from)
	if err != nil {
		fmt.Printf("Can not read from file: %v \n", from)
		os.Exit(1)
	}


	dst, err := os.Create(to)
	if err != nil {
		fmt.Printf("Can not write to file: %v \n", to)
		os.Exit(1)
	}


	if _, err := src.Seek(int64(offset), 0); err != nil {
		fmt.Printf("Can not copy from file %v to file: %v \n", from, to)
		os.Exit(1)
	}

	if _, err:= io.CopyN(dst, src, int64(limit)); err != nil {
		fmt.Printf("Can not copy from file %v to file: %v \n", from, to)
		os.Exit(1)
	}

}
