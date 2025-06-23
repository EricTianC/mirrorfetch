// Package model
// 对单个资源服务进行的抽象
// 便于对有多种类型镜像的单个语言（开源软件）进行检测
package model

type Service struct {
	ID          string `json:"id,omitempty"` // 标识该 Service 的唯一 Id，遵守 "Lean4/Toolchain/Compressed" 格式  TODO: 细化该格式
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type ServiceProvider struct {
	ID          string   `json:"id,omitempty"`   // 标识对应的 Service 的唯一 Id，遵守 "Lean4/Toolchain/Compressed" 格式  TODO: 细化该格式
	Name        string   `json:"name,omitempty"` // 每个镜像站可能有不同的具体名字
	Description string   `json:"description,omitempty"`
	TestTarget  TestSpec `json:"test_target"` // 用于检测资源可访问性的 URL，可能不为 http/https 协议：
	// 现暂有: http/https 协议。
}

type TestResult interface {
	Ok() bool
	GetName() string
	ToInfo() string
}

type TestSpec struct {
	Sort string `json:"sort"`
	URL  string `json:"url"`
}
