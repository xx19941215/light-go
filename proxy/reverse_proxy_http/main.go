package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

//NewMultipleHostsReverseProxy 反向代理
func NewMultipleHostsReverseProxy(targets []*url.URL) *httputil.ReverseProxy {
	director := func(req *http.Request) {
		target := targets[rand.Int()%len(targets)]
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.URL.Path = target.Path
	}

	return &httputil.ReverseProxy{Director: director}
}

func main() {
	proxy := NewMultipleHostsReverseProxy([]*url.URL{
		{
			Scheme: "http",
			Host:   "localhost:9091",
		},
		{
			Scheme: "http",
			Host:   "localhost:9092",
		},
	})

	go serverMaster()

	log.Fatal(http.ListenAndServe(":9090", proxy))
}

func serverMaster() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello this master"))
		now := time.Now().Format("2006/01/02 15:04:05")
		fmt.Println(now + ": this is master")
	})

	log.Fatal(http.ListenAndServe(":9091", nil))
}
