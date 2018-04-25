package main

import (
	"log"
	"net/http"
)

// myHandler
type myHandler struct{}

// ServeHTTP
// 参数:
//      响应 接口
//      请求 类型
func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// io.WriteString(w, "URL:"+r.URL.String())
	w.Write([]byte("Web v2, URL:" + r.URL.String()))

}

// handler
//      响应 接口
//      请求 类型
func sayHello(w http.ResponseWriter, r *http.Request) {
	// io.WriteString(w, "Hello Web -V 2")
	w.Write([]byte("Hello Web -V 2"))
}

// 输入:http://localhost:8080
// 输出:URL:/
// 输入:http://localhost:8080/aaaa
// 输出:URL:/aaaa
// 输入:http://localhost:8080/hello
// 输出:Hello Web -V 2
func main() {
	// ServeMux
	mux := http.NewServeMux()
	// 路由
	mux.Handle("/", &myHandler{})
	mux.HandleFunc("/hello", sayHello)

	log.Println("Starting server... v2")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
