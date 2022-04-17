package log

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

var re = `\d{4}-\d{1,2}-\d{1,2}[ ]\d{2}[:]\d{2}:\d{2}`
var re1 = `ERROR`

func LogAlyze() {

	file, err := os.Open("")
	//handle errors while opening
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}

	fileScanner := bufio.NewScanner(file)
	//var logmap = map[bool]string{}
	// read line by line
	var signtext string
	for fileScanner.Scan() {
		text := fileScanner.Text()

		match, _ := regexp.MatchString(re, text)
		//var i = 0
		//var tempmatchstr string
		if match {
			text = "==--" + text
		}
		signtext = signtext + text + "\n"
		//fmt.Println(text)
		//logmap[match] = signtext
	}
	//fmt.Println(logmap[true])
	spilttext := strings.Split(signtext, "==--")
	fmt.Println(len(spilttext))
	//form
	for i := 0; i < len(spilttext); i++ {
		if matchError, _ := regexp.MatchString(re1, spilttext[i]); matchError {
			fmt.Println(spilttext[i])
		}
		//fmt.Println(spilttext[i])
	}

	// handle first encountered error while reading
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}

	file.Close()
}
