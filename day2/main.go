
package main

import (
"strings"
"strconv"
"regexp"
"bufio"
"fmt"
"os"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
	}
	var sum int = 0
	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanLines)
	page := 1
	for scan.Scan() {
		line := scan.Text()
		split := strings.Split(line, ";")
		for _, ses := range split {
			var red int64 = 0
			var green int64 = 0
			var blue int64 = 0
			regreen := regexp.MustCompile(`(\d+)\s*green`)
			if regreen.MatchString(ses) {
				all := regreen.FindAllStringSubmatch(ses, -1)
				for _, char := range all {
					i, err := strconv.ParseInt(char[1], 0, 64)
					if err == nil { green += i }
				}
			}
			reblue := regexp.MustCompile(`(\d+)\s*blue`)
			if reblue.MatchString(ses) {
				all := reblue.FindAllStringSubmatch(ses, -1)
				for _, char := range all {
					i, err := strconv.ParseInt(char[1], 0, 64)
					if err == nil { blue += i }
				}
			}
			rered := regexp.MustCompile(`(\d+)\s*red`)
			if rered.MatchString(ses) {
				all := rered.FindAllStringSubmatch(ses, -1)
				for _, char := range all {
					i, err := strconv.ParseInt(char[1], 0, 64)
					if err == nil { red += i }
				}
			}
			if red == 12 && green == 13 && blue == 14 {
				fmt.Println("FOUNDD", page)
				sum += page
			}
			fmt.Println("GAME ", page, " / red = ", red, " / green = ", green, " / blue = ", blue, ses)
		}
		page++
	}
	fmt.Println("SUM = ", sum)
	file.Close()
}