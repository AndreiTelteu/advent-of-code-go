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
	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanLines)
	var total int64 = 0
	for scan.Scan() {
		line := scan.Text()
		line = strings.Replace(line, "one", "one1one", -1)
		line = strings.Replace(line, "two", "two2two", -1)
		line = strings.Replace(line, "three", "three3three", -1)
		line = strings.Replace(line, "four", "four4four", -1)
		line = strings.Replace(line, "five", "five5five", -1)
		line = strings.Replace(line, "six", "six6six", -1)
		line = strings.Replace(line, "seven", "seven7seven", -1)
		line = strings.Replace(line, "eight", "eight8eight", -1)
		line = strings.Replace(line, "nine", "nine9nine", -1)
		var first string
		var last string
		re := regexp.MustCompile(`\d`)
		if re.MatchString(line) {
			all := re.FindAllString(line, -1)
			for index, char := range all {
				if index == 0 {
					first = char
				}
				last = char
			}
		}
		sum, err := strconv.ParseInt(fmt.Sprintf("%s%s", first, last), 0, 64)
		if err != nil {
			fmt.Println(err)
		}
		if err == nil {
			total += sum
		}
		fmt.Println(sum)
	}
	fmt.Println(total)
	file.Close()
}
