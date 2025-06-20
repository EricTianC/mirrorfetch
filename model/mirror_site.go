// Package model 定义模型类
package model

import (
	"io"
	"log/slog"
	"net/http"
)

// MirrorSite
//
// Name: 镜像站名
// Description: 简单描述
// HomeUrl: 主页网址，必须携带 http:// 或 https:// 前缀
type MirrorSite struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	HomeUrl     string `json:"homeUrl"`
}

// TouchHomePage 检查主页面可访问性
func (site *MirrorSite) TouchHomePage() bool {
	resp, err := http.Get(site.HomeUrl)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false
	}
	slog.Debug("Home page: %s", string(body))
	return true
}
