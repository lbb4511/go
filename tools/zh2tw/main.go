package main

// Copyright 2013 <lbb4511{AT}126.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ingore

//
// 简体字转繁体字, 1:1转换, 支持逆向转换.
//
// Example:
//	zh2tw dir
//	zh2tw dir "\.go$"
//	zh2tw dir "\.md$" zh2tw
//	zh2tw dir "\.md$" tw2zh
//
// Help:
//	zh2tw -h
//

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"unicode/utf8"
)

const MaxFileSize = 8 << 20 // 8MB

const usage = `
Usage: zh2tw dir [nameFilter]
       zh2tw -h

Example:
  zh2tw dir
  zh2tw dir "\.go$"
  zh2tw dir "\.md$" zh2tw
  zh2tw dir "\.md$" tw2zh

  Report bugs to <lbb4511{AT}126.com>.
`

func main() {
	if len(os.Args) < 2 || os.Args[1] == "-h" {
		fmt.Fprintln(os.Stderr, usage[1:len(usage)-1])
		os.Exit(0)
	}
	dir, nameFilter, method := os.Args[1], ".*", "zh2tw"
	if len(os.Args) > 2 {
		nameFilter = os.Args[2]
	}
	if len(os.Args) > 3 {
		method = os.Args[3]
	}

	if method != "zh2tw" && method != "tw2zh" {
		fmt.Fprintln(os.Stderr, usage[1:len(usage)-1])
		os.Exit(0)
	}

	total := 0
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal("filepath.Walk: ", err)
			return err
		}
		if info.IsDir() {
			return nil
		}
		relpath, err := filepath.Rel(dir, path)
		if err != nil {
			log.Fatal("filepath.Rel: ", err)
			return err
		}
		mathed, err := regexp.MatchString(nameFilter, relpath)
		if err != nil {
			log.Fatal("regexp.MatchString: ", err)
		}
		if mathed {
			if changed := convertFile(path, method); changed {
				fmt.Printf("%s\n", relpath)
				total++
			}
		}
		return nil
	})
	fmt.Printf("%s total %d\n", method, total)
}

func convertFile(path, method string) (changed bool) {
	abspath, err := filepath.Abs(path)
	if err != nil {
		log.Fatal("convertFile: filepath.Abs:", err)
	}

	fi, err := os.Lstat(abspath)
	if err != nil {
		log.Fatal("convertFile: os.Lstat:", err)
	}
	if fi.Size() > MaxFileSize {
		return false
	}

	oldData, err := ioutil.ReadFile(abspath)
	if err != nil {
		log.Fatal("convertFile: ioutil.ReadFile:", err)
	}
	if !utf8.Valid(oldData) {
		return false
	}

	newData := oldData
	switch {
	case method == "zh2tw":
		newData = []byte(zh2tw(string(oldData)))
	case method == "tw2zh":
		newData = []byte(tw2zh(string(oldData)))
	}

	if string(newData) == string(oldData) {
		return false
	}

	err = ioutil.WriteFile(abspath, newData, 0666)
	if err != nil {
		log.Fatal("convertFile: ioutil.WriteFile:", err)
	}
	return true
}

func zh2tw(s string) string {
	old := []rune(s)
	new := make([]rune, 0)
	for _, c := range old {
		if x, ok := zh2twMap[c]; ok {
			new = append(new, x)
		} else {
			new = append(new, c)
		}
	}
	return string(new)
}

func tw2zh(s string) string {
	old := []rune(s)
	new := make([]rune, 0)
	for _, c := range old {
		if x, ok := tw2zhMap[c]; ok {
			new = append(new, x)
		} else {
			new = append(new, c)
		}
	}
	return string(new)
}
