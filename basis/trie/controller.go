package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"time"
)

type SchoolInfo struct {
	SchoolId   int    `json:"school_id"`
	Province   string `json:"province"`
	City       string `json:"city"`
	SchoolType int    `json:"school_type"`
	SchoolName string `json:"school_name"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	if err := template.Must(template.New("school").Parse(`<html>
		<script src="http://code.jquery.com/jquery-latest.js"></script>
		<script>
			$(document).ready(function () {
				var flag = true;
				$('#input_id').on('compositionstart', function () {
					flag = false;
				})
		
				$('#input_id').on('compositionend', function () {
					flag = true;
				})
		
				$('#input_id').on('input', function () {
					var _this = this;
					setTimeout(function () {
						if (!flag) {
							return
						}
		
						keyword = $(_this).val()
						if (keyword.length == 0) {
							return
						}
		
						$.getJSON("http://localhost:8080/search?keyword=" + keyword, function (result) {
							if (result.code != 0) {
								return
							}
							$("#suggest_id option").remove();
							$.each(result.data, function (key, val) {
								$('#suggest_id').append("<option value='" + val.school_name + "'>");
							});
						});
					}, 0)
				})
			});
		</script>
		
		<body>
			<div>
				<span> 请输入学校名字:</span>
				<input id="input_id" type="text" name="keyword" list="suggest_id"></input>
				<datalist id="suggest_id">
				</datalist>
		</body>
		
		</html>`)).Execute(w, nil); err != nil {
		fmt.Fprintf(w, "Execute: %v", err)
		return
	}
}

func Search(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	keyword := r.FormValue("keyword")

	start := time.Now().UnixNano()
	//schools := logic.SearchV2(keyword, 16)
	schools := SearchV1(keyword, 16)
	end := time.Now().UnixNano()
	fmt.Printf("暴力:keyword:%s result:%d cost:%d us\n", keyword, len(schools), (end-start)/1000)

	m := make(map[string]interface{}, 16)
	m["code"] = 0
	m["message"] = "sucess"
	m["data"] = schools

	result, err := json.Marshal(m)
	if err != nil {
		return
	}

	w.Write(result)
}

func SearchV1(keyword string, limit int) (schools []*SchoolInfo) {
	for _, s := range SchoolList {
		if strings.HasPrefix(s.SchoolName, keyword) == true {
			schools = append(schools, s)
			if len(schools) > limit {
				break
			}
		}
	}

	return
}

func SearchV2(keyword string, limit int) (schools []*SchoolInfo) {

	nodes := trieTree.PrefixSearch(keyword, limit)
	//fmt.Printf("len:%d\n", len(nodes))
	for _, v := range nodes {
		school, ok := v.Data.(*SchoolInfo)
		if !ok {
			fmt.Printf("invalid school data:%v", v)
			continue
		}

		schools = append(schools, school)
	}
	return
}
