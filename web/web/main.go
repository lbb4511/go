package main

import (
	//输入输出
	"log"      // 写log
	"net/http" //路由
	"time"     // 超时
)

// 自间mux
var mux map[string]func(http.ResponseWriter, *http.Request)

// myHandler
type myHandler struct{}

// ServeHTTP
// 参数:
//      响应 接口
//      请求 类型
func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := mux[r.URL.String()]; ok {
		h(w, r)
		return
	}
	w.Write([]byte("Web -V 3, URL:" + r.URL.String()))

}

// handler
// 参数:
//      响应 接口
//      请求 类型
func sayHello(w http.ResponseWriter, r *http.Request) {
	// io.WriteString(w, "Hello Web -V 3")
	w.Write([]byte("Hello Web -V 3"))
}

// handler
// 参数:
//      响应 接口
//      请求 类型
func sayBye(w http.ResponseWriter, r *http.Request) {
	// io.WriteString(w, "Bye Web -V 3")
	w.Write([]byte("Bye Web -V 3"))
}

// 输入:http://localhost:8080
// 输出:URL:/
// 输入:http://localhost:8080/aaaa
// 输出:URL:/aaaa
// 输入:http://localhost:8080/hello
// 输出:Hello Web -V 3
// 输入:http://localhost:8080/bye
// 输出:Bye Web -V 3
func main() {
	// server := &http.Server{
	// 	Addr:         ":8080",         // 地址
	//  ReadTimeout:  5 * time.Second, // 5秒超时
	//  WriteTimeout: 5 * time.Second, // 5秒超时
	// }

	// quit := make(chan os.Signal)
	// signal.Notify(quit, os.Interrupt)

	// mux := http.NewServeMux()
	// mux.Handle("/", &myHandler{})
	// mux.HandleFunc("/bye", sayBye)

	// go func() {
	// 	<-quit

	// 	if err := server.Close(); err != nil {
	// 		log.Fatal("Close server:", err)
	// 	}
	// }()

	// server.Handler = mux

	server := http.Server{
		Addr:         ":8080",         // 地址
		Handler:      &myHandler{},    // myHandler
		ReadTimeout:  5 * time.Second, // 5秒超时
		WriteTimeout: 5 * time.Second, // 5秒超时
	}
	mux = make(map[string]func(http.ResponseWriter, *http.Request))
	mux["/hello"] = sayHello
	mux["/bye"] = sayBye

	log.Print("Starting server... v3")
	err := server.ListenAndServe()
	if err != nil {
		if err == http.ErrServerClosed {
			log.Print("Server closed under request")
		} else {
			log.Fatal("Server closed unexpected")
		}
	}
}
