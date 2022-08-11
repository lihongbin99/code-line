package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	d  = "C:\\Users\\Lee\\IdeaProjects\\srp-go"
	sd = make(map[string]int)
	sf = make(map[string]int)

	line      = 0
	fileCount = 0
)

func init() {
	sd[".idea"] = 1
	sd["bin"] = 1
	sd["security"] = 1

	sf["go.mod"] = 1
	sf["srp-go.iml"] = 1
}

func main() {
	scanDir(d)
	fmt.Println(line, "line")
	fmt.Println(fileCount, "fileCount")
}

func scanDir(dirPath string) {
	dir, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.Fatalln(err)
	}

	for _, f := range dir {
		name := dirPath + string(os.PathSeparator) + f.Name()
		if f.IsDir() {
			if _, ok := sd[f.Name()]; !ok {
				scanDir(name)
			}
		} else {
			if _, ok := sf[f.Name()]; !ok {
				scanFile(name)
			}
		}
	}
}

func scanFile(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() { _ = file.Close() }()

	br := bufio.NewReader(file)
	for {
		sb, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		s := strings.TrimSpace(string(sb))
		if len(s) > 0 {
			line++
		}
	}
	fileCount++
}
