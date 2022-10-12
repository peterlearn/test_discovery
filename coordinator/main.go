package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

func handleHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle coordinate")
	_, _ = w.Write([]byte("Hello coordinate!\n"))
}

func main() {
	flag.Parse()
	http.HandleFunc("/hello", handleHello)
	go http.ListenAndServe(":8888", nil)
	register()
	select {}
}

const (
	ServiceName = "coordinator"
	Namespace   = "test"
)

func register() {
	if os.Getenv("DISCOVERY") != "" {
		_, err := RegisterDiscovery(fmt.Sprintf("%s.%s", Namespace, ServiceName))
		if err != nil {
			panic(err)
		}
	}
}
