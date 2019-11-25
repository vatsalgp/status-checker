package main 

import (
	"time"
	"os"
	"strconv"
)

func main () {

	links := readFromDrive("links.txt")
	var intv int64 = 3
	var cont bool = false
	var errBool error = nil
	var errInt error = nil

	if len(os.Args) >=2{
		cont,errBool = strconv.ParseBool(os.Args[1])
		if errBool!=nil {
			cont = false
		}
	}	
	if len(os.Args) >= 3{
		intv,errInt = strconv.ParseInt(os.Args[2], 10, 64)
		if errInt!=nil || intv < 1 {
			intv = 3
		}
	}
	c := make(chan string)
	
	for _,link := range links {
		go checkLink(link, c)
	}

	if (cont){
		for link := range c{
			go func(link string) {
				time.Sleep(time.Duration(intv) * time.Second)
				checkLink(link, c)
			}(link)
		}
	} else {
		for i:=0; i<len(links); i++ {
			<-c	
		}
	}
}
