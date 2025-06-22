// Package model 储存公用的数据类型及方法
package model

import (
	"crypto/tls"
	"log/slog"
	"net/http"
	"net/http/httptrace"
	"time"
)

type HTTPTracesResponse struct {
	Name          string        `json:"name"`
	Reachable     bool          `json:"reachable"`
	DNSDuration   time.Duration `json:"dns_duration"`
	TCPDuration   time.Duration `json:"tcp_duration"`
	TLSDuration   time.Duration `json:"tls_duration"`
	TotalDuration time.Duration `json:"total_duration"`
	StatusCode    int           `json:"status_code"`
	ErrorMessage  string        `json:"error_message"`
}

type NamedRemote interface {
	GetName() string
	GetHomeURL() string
}

func TouchHome(source NamedRemote) HTTPTracesResponse {
	var (
		dnsStart, dnsDone   time.Time
		connStart, connDone time.Time
		tlsStart, tlsDone   time.Time
	)

	response := HTTPTracesResponse{Name: source.GetName()}

	req, err := http.NewRequest("GET", source.GetHomeURL(), nil)
	if err != nil {
		response.ErrorMessage = err.Error()
		return response
	}

	trace := &httptrace.ClientTrace{
		DNSStart: func(info httptrace.DNSStartInfo) {
			dnsStart = time.Now()
		},
		DNSDone: func(info httptrace.DNSDoneInfo) {
			dnsDone = time.Now()
		},
		ConnectStart: func(network, addr string) {
			connStart = time.Now()
		},
		ConnectDone: func(network, addr string, err error) {
			connDone = time.Now()
			if err != nil {
				slog.Debug("❌ 连接错误:", "error", err)
			}
		},
		TLSHandshakeStart: func() {
			tlsStart = time.Now()
		},
		TLSHandshakeDone: func(state tls.ConnectionState, err error) {
			tlsDone = time.Now()
			if err != nil {
				slog.Debug("❌ TLS 错误:", "error", err)
			}
		},
	}

	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))

	start := time.Now()
	resp, err := http.DefaultTransport.RoundTrip(req)
	total := time.Since(start)

	if err != nil {
		response.ErrorMessage = err.Error()
		return response
	}
	defer resp.Body.Close()

	response = HTTPTracesResponse{
		Name:          source.GetName(),
		Reachable:     true,
		DNSDuration:   dnsDone.Sub(dnsStart),
		TCPDuration:   connDone.Sub(connStart),
		TLSDuration:   tlsDone.Sub(tlsStart),
		TotalDuration: total,
		StatusCode:    resp.StatusCode,
	}
	return response
}

func ToNamedRemoteList[T NamedRemote](sources []T) []NamedRemote {
	ret := make([]NamedRemote, len(sources))
	for i, source := range sources {
		ret[i] = source
	}
	return ret
}
