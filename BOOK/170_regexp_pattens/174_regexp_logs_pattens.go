package main

import (
	"bufio"
	"fmt"
	"github.com/sqweek/dialog"
	"io"
	"os"
	"regexp"
	"strings"
	"time"
)

func main() {
	filename, err := dialog.File().Load()
	if err != nil {
		fmt.Println("Please choose a correct file! ", err)
		os.Exit(1)
	}
	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("error opening file %s", err)
		os.Exit(1)
	}
	defer f.Close()

	notAMatch := 0
	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading file %s", err)
		}
		r1 := regexp.MustCompile(
			`.*\[(\d\d/\w+/\d\d\d\d:\d\d:\d\d:\d\d.*)\] .*`)
		if r1.MatchString(line) {
			match := r1.FindStringSubmatch(line)
			d1, err := time.Parse(
				"02/Jan/2006:15:04:05 -0700", match[1])
			if err == nil {
				newFormat := d1.Format(time.Stamp)
				fmt.Print(strings.Replace(line, match[1], newFormat, 1))
			} else {
				notAMatch++
			}
			continue
		}

		r2 := regexp.MustCompile(
			`.*\[(\w+\-\d\d-\d\d:\d\d:\d\d:\d\d.*)\] .*`)
		if r2.MatchString(line) {
			match := r2.FindStringSubmatch(line)
			d1, err := time.Parse("Jan-02-06:15:04:05 -0700", match[1])
			if err == nil {
				newFormat := d1.Format(time.Stamp)
				fmt.Print(strings.Replace(line, match[1], newFormat, 1))
			} else {
				notAMatch++
			}
			continue
		}

	}
	fmt.Println(notAMatch, "lines did not match!")
}
