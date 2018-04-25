package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type myHandle struct{}

//返回的jsonBean
type BaseJsonBean struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var mux map[string]func(http.ResponseWriter, *http.Request)

func (*myHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("请求url:", r.URL.String())
	log.Println("请求方法:", r.Method)
	//解析，默认不解析的，否者r.Form将拿不到数据
	r.ParseForm()
	log.Println("请求报文:", r)
	log.Println("请求的参数:", r.Form)
	if h, ok := mux[r.URL.String()]; ok {
		h(w, r)
	} else {
		// 文件过滤器
		path := r.URL.Path
		//判断是否有.
		if strings.Contains(path, ".") {
			request_type := path[strings.LastIndex(path, "."):]
			switch request_type {
			case ".css":
				w.Header().Set("content-type", "text/css; charset=utf-8")
			case ".js":
				w.Header().Set("content-type", "text/javascript; charset=utf-8")
			default:
			}
		}
		wd, err := os.Getwd()
		if err != nil {
			log.Println("获取系统路径失败：", err)
		}
		fin, err := os.Open(wd + path)
		if err != nil {
			log.Println("读取文件失败：", err)
			//关闭文件句柄
			fin.Close()
			//返回json头
			w.Header().Set("content-type", "text/json; charset=utf-8")
			bytes, _ := json.Marshal(&BaseJsonBean{Code: 404, Message: "", Data: ""})
			w.Write([]byte(string(bytes)))
			log.Println("返回数据：", string(bytes))
			return
		}
		fd, _ := ioutil.ReadAll(fin)
		w.Write([]byte(fd))
	}
}

func main() {
	server := http.Server{
		Addr:        ":8080",
		Handler:     &myHandle{},
		ReadTimeout: 5 * time.Second,
	}
	mux = make(map[string]func(http.ResponseWriter, *http.Request))
	//这里配置路由，可以添加自己的方法去处理对应的路由
	mux["/"] = func(w http.ResponseWriter, r *http.Request) {

		if t, err := template.ParseFiles("index.html"); err != nil {
			log.Println("未找到index.html文件，我为您展示默认的首页 :) 快去创建自己的首页吧")
			w.Header().Set("content-type", "text/html; charset=utf-8")
			w.Write([]byte(`<!DOCTYPE html>
				<html>
				
				<head>
					<meta charset="utf-8">
					<title>It Work</title>
				</head>
				
				<body>
					<h1>:)</h1>
					<h3>It Work</h3>
				</body>
				
				</html>`))
		} else {
			t.Execute(w, nil)
		}
	}
	log.Println("已为您启动了服务，您可以打开浏览器访问 http://127.0.0.1:8080 ，我会输入访问日志。")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
