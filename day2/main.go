package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
	}
	var sum int64 = 0
	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanLines)
	page := 1
	for scan.Scan() {
		line := scan.Text()
		split := strings.Split(line, ";")
		var red int64 = 0
		var green int64 = 0
		var blue int64 = 0
		for _, ses := range split {
			regreen := regexp.MustCompile(`(\d+)\s*green`)
			if regreen.MatchString(ses) {
				all := regreen.FindAllStringSubmatch(ses, -1)
				for _, char := range all {
					i, err := strconv.ParseInt(char[1], 0, 64)
					if err == nil && i > green {
						green = i
					}
				}
			}
			reblue := regexp.MustCompile(`(\d+)\s*blue`)
			if reblue.MatchString(ses) {
				all := reblue.FindAllStringSubmatch(ses, -1)
				for _, char := range all {
					i, err := strconv.ParseInt(char[1], 0, 64)
					if err == nil && i > blue {
						blue = i
					}
				}
			}
			rered := regexp.MustCompile(`(\d+)\s*red`)
			if rered.MatchString(ses) {
				all := rered.FindAllStringSubmatch(ses, -1)
				for _, char := range all {
					i, err := strconv.ParseInt(char[1], 0, 64)
					if err == nil && i > red {
						red = i
					}
				}
			}
		}
		sum += red * green * blue
		page++
	}
	fmt.Println("SUM = ", sum)
	file.Close()
}
