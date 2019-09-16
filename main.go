package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

// Takes line separated links from input and writes end URLs after all redirects into output file.

var inputPath = "input.txt"
var outputPath = "output.txt"

func main() {
	links, err := readLines(inputPath)
	if err != nil {
		fmt.Println(err)
	}
	var output []string
	for i, link := range links {
		fmt.Println(i, link)
		resp, err := http.Get(link)
		if err != nil {
			fmt.Printf("%v http.Get => %v", i, err.Error())
		}
		finalURL := resp.Request.URL.String()
		fmt.Println(i, finalURL)
		output = append(output, link+"\t"+finalURL+"\n")
	}
	writeStringArrayToFile(outputPath, output, 0775)
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func writeStringArrayToFile(filename string, strArray []string, perm os.FileMode) {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, perm)
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()
	for _, v := range strArray {
		if _, err = f.WriteString(v); err != nil {
			log.Panic(err)
		}
	}
}
