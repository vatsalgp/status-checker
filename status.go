package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"strings"
	"os"
)

func checkLink(link string, c chan string) {
	_,err := http.Get(link)

	if err!=nil {
		fmt.Println(link, "is not working")
	}else {
		fmt.Println(link, "is working")
	}

	c <- link
}

func readFromDrive(filename string) []string{
	bs,err := ioutil.ReadFile(filename)
	if err!=nil {
		fmt.Println("Error: ", err);
		os.Exit(1);
	}

	lines := strings.Split(string(bs), "\n")
	for i,line := range lines {
		if line[0]=='#' {
			lines = append(lines[:i],lines[i+1:]...)
		}
	}
	return lines
}