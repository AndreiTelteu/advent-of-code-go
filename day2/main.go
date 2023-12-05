
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
		var red int64 = 0
		var green int64 = 0
		var blue int64 = 0
		poss := true
		for _, ses := range split {
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
			if red > 12 || green > 13 || blue > 14 {
				poss = false
			}
			red = 0; green = 0; blue = 0;
		}
		if poss == true {
			fmt.Println("GAME FOUND =", page)
			sum += page
		}
		fmt.Println("GAME ", page, " / red = ", red, " / green = ", green, " / blue = ", blue)
		page++
	}
	fmt.Println("SUM = ", sum)
	file.Close()
}