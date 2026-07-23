package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! 2")
}

func main() {
	// 注册路由处理器
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/health", helloHandler)

	// 启动服务器
	fmt.Println("服务器启动在 http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
