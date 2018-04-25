package main

import (
	"log"
	"net/http"
)

// handler
// 参数:
//      响应 接口
//      请求 类型
func sayHello(w http.ResponseWriter, r *http.Request) {
	// io.WriteString(w, "Hello Web -V 1")
	w.Write([]byte("Hello Web -V 1"))
}

func main() {
	// 设置路由 --方法sayHello可以当作参数传递...好牛
	http.HandleFunc("/", sayHello)

	log.Println("Starting server... v1")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
