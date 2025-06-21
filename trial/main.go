package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/http/httptrace"
	"time"
)

func main() {
	TraceHttpRequest("https://cn.bing.com")
	TraceHttpRequest("https://github.com")
	TraceHttpRequest("https://huggingface.co/")
	TraceHttpRequest("https://www.baidu.com/")
	TraceHttpRequest("https://www.ustc.edu.cn/")
}

// TraceHttpRequest 发起 HTTP 请求并输出各阶段耗时
func TraceHttpRequest(url string) error {
	var (
		dnsStart, dnsDone   time.Time
		connStart, connDone time.Time
		tlsStart, tlsDone   time.Time
		//firstByteTime       time.Time
	)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	trace := &httptrace.ClientTrace{
		DNSStart: func(info httptrace.DNSStartInfo) {
			dnsStart = time.Now()
		},
		DNSDone: func(info httptrace.DNSDoneInfo) {
			dnsDone = time.Now()
			fmt.Printf("📡 DNS 解析耗时: %v\n", dnsDone.Sub(dnsStart))
		},
		ConnectStart: func(network, addr string) {
			connStart = time.Now()
		},
		ConnectDone: func(network, addr string, err error) {
			connDone = time.Now()
			fmt.Printf("🔌 TCP 连接耗时: %v\n", connDone.Sub(connStart))
			if err != nil {
				fmt.Printf("❌ 连接错误: %v\n", err)
			}
		},
		TLSHandshakeStart: func() {
			tlsStart = time.Now()
		},
		TLSHandshakeDone: func(state tls.ConnectionState, err error) {
			tlsDone = time.Now()
			fmt.Printf("🔒 TLS 握手耗时: %v\n", tlsDone.Sub(tlsStart))
			if err != nil {
				fmt.Printf("❌ TLS 错误: %v\n", err)
			}
		},
		GotFirstResponseByte: func() {
			//firstByteTime = time.Now()
			fmt.Println("📥 首字节已接收")
		},
	}

	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))

	fmt.Printf("开始尝试访问 %v\n", req.URL)

	start := time.Now()
	resp, err := http.DefaultTransport.RoundTrip(req)
	total := time.Since(start)

	if err != nil {
		return fmt.Errorf("请求失败: %v", err)
	}
	defer resp.Body.Close()

	fmt.Printf("✅ 总耗时: %v\n", total)
	fmt.Printf("🌐 HTTP 状态码: %v\n", resp.StatusCode)
	return nil
}
