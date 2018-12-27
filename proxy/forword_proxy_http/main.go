package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
)

//Proxy 正向代理对象
type Proxy struct{}

//ServeHTTP 监听
func (p *Proxy) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	fmt.Printf("Received request %s %s %s \n", req.Method, req.Host, req.RemoteAddr)
	transport := http.DefaultTransport

	outReq := new(http.Request)

	//复制原来的请求对象
	*outReq = *req

	if clientIP, _, err := net.SplitHostPort(req.RemoteAddr); err == nil {
		//每个代理服务器会在 X-Forwarded-For 头部填上前一个节点的 ip 地址
		if prior, ok := outReq.Header["X-Forwarded-For"]; ok {
			clientIP = strings.Join(prior, ",") + "," + clientIP
		}

		//添加上 X-Forward-For 头部
		outReq.Header.Set("X-Forwarded-For", clientIP)
	}

	//把新请求发送到服务器端
	res, err := transport.RoundTrip(outReq)

	if err != nil {
		rw.WriteHeader(http.StatusBadGateway)
		return
	}

	//设置Header
	for key, val := range res.Header {
		for _, v := range val {
			rw.Header().Add(key, v)
		}
	}

	//设置状态码
	rw.WriteHeader(res.StatusCode)
	//复制响应主体
	io.Copy(rw, res.Body)
	res.Body.Close()
}

func main() {
	fmt.Println("serve on :8080")
	http.Handle("/", &Proxy{})
	http.ListenAndServe("0.0.0.0:8080", nil)
}
