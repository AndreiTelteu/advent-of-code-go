
package main

import (
//"strings"
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
	scan := bufio.NewScanner(file)
	var lines []string
	scan.Split(bufio.ScanLines)
	for scan.Scan() {
		lines = append(lines, scan.Text())
	}
	var total int64 = 0
	for index, line := range lines {
		reg := regexp.MustCompile(`[0-9]+`)
		if !reg.MatchString(line) { continue }
		matches := reg.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			numIndex := findStr(line, match[0])
			if numIndex == -1 { continue }
			//fmt.Println(index, line, match[0], numIndex)
			if hasAdjacent(lines, index, numIndex, len(match[0])) {
				num, err := strconv.ParseInt(match[0], 0, 64)
				if err == nil { total += num }
			}
		}
	}
	fmt.Println("TOTAL: ", total)
	file.Close()
}

func hasAdjacent(lines []string, index int, start int, length int) bool {
	if start > 0 && isSymbol(lines[index][start-1:start]) {
		return true
	}
	if start+length < len(lines[index]) && isSymbol(lines[index][start+length:start+length+1]) {
		return true
	}
	if index > 0 {
		if start > 0 && isSymbol(lines[index-1][start-1:start]) {
			return true
		}
		if start+length < len(lines[index-1]) && isSymbol(lines[index-1][start+length:start+length+1]) {
			return true
		}
		if isSymbol(lines[index-1][start:start+length]) {
			return true
		}
	}
	if index < len(lines)-1 {
		if start > 0 && isSymbol(lines[index+1][start-1:start]) {
			return true
		}
		if start+length < len(lines[index+1]) && isSymbol(lines[index+1][start+length:start+length+1]) {
			return true
		}
		if isSymbol(lines[index+1][start:start+length]) {
			return true
		}
	}
	fmt.Println("not good", lines[index], index, start, length)
	return false
}

func isSymbol(s string) bool {
	return regexp.MustCompile(`[^a-zA-z0-9\.]`).MatchString(s)
}

func findStr(haystack string, needle string) int {
	for i := range haystack {
		if haystack[i : i+len(needle)] == needle {
			return i
		}
	}
	return -1
}
