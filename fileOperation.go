package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main_fileOperation() {
	f, err := os.Open("../5-type-converison/read-file.txt")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	f, err = os.Open("read-file.txt")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make([]byte, 100)
	n, err := f.Read(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Read %d bytes: %s\n", n, string(data[:n]))
	ioutilData, err := ioutil.ReadFile("read-file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Read %d bytes: %s\n", len(ioutilData), string(ioutilData))

}
