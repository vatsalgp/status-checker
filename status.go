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

	doc := strings.TrimSpace(string(bs))
	lines := strings.Split(doc, "\n")
	res := []string{}

	for _,line := range lines {
		line = strings.TrimSpace(line)
		if (line[0]!='#' && strings.Contains(line, "http://")){
			res = append(res, line)
		}
	}
	return res
}
