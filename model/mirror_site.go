// Package model 定义模型类
package model

// MirrorSite struct
//
// Name: 镜像站名
// Description: 简单描述
// HomeUrl: 主页网址，必须携带 http:// 或 https:// 前缀
type MirrorSite struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	HomeUrl     string `json:"homeUrl"`
}

func (site *MirrorSite) GetName() string {
	return site.Name
}

func (site *MirrorSite) GetHomeUrl() string {
	return site.HomeUrl
}
