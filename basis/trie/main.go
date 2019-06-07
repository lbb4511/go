package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var SchoolList []*SchoolInfo

var trieTree *Trie = NewTrie()

func main() {

	filename := "./school.dat"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("open %s failed, err:%v\n", filename, err)
		return
	}
	defer file.Close()
	var id int
	reader := bufio.NewReader(file)
	for {
		line, errRet := reader.ReadString('\n')
		if errRet == io.EOF {
			break
		}

		if errRet != nil {
			err = errRet
			fmt.Printf("read %s failed, err:%v\n", filename, err)
			return
		}

		strSplit := strings.Split(line, "\t")
		if len(strSplit) != 4 {
			fmt.Printf("invalid schoool info, line:%s\n", line)
			continue
		}

		var schoolInfo SchoolInfo
		id++
		schoolInfo.SchoolId = id
		schoolInfo.Province = strings.TrimSpace(strSplit[0])
		schoolInfo.City = strings.TrimSpace(strSplit[1])
		schoolInfo.SchoolName = strings.TrimSpace(strSplit[2])

		schoolType, errRet := strconv.Atoi(strings.TrimSpace(strSplit[3]))
		if errRet != nil {
			fmt.Printf("invalid schoool info, line:%s, err:%v\n", line, err)
			continue
		}

		schoolInfo.SchoolType = schoolType
		SchoolList = append(SchoolList, &schoolInfo)
		trieTree.Add(fmt.Sprintf("%s%d", schoolInfo.SchoolName, schoolInfo.SchoolId), &schoolInfo)
		fmt.Printf("school:%+v\n", schoolInfo)
	}
	/*
		for i := 0; i < 1000000; i++ {
			var schoolInfo SchoolInfo
			schoolInfo.City = "长沙"
			schoolInfo.Province = "湖南"
			schoolInfo.SchoolId = id
			id++
			schoolInfo.SchoolName = fmt.Sprintf("学校名字%d", i)
			SchoolList= append(SchoolList, &schoolInfo)
			trieTree.Add(fmt.Sprintf("%s%d", schoolInfo.SchoolName, schoolInfo.SchoolId), &schoolInfo)
		}
	*/

	http.HandleFunc("/", Index)
	http.HandleFunc("/search", Search)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("listen failed, err:%v\n", err)
		return
	}
}
