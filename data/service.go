package data

import "github.com/erictianc/mirrorfetch/model"

var MirrorServices = []*model.Service{
	{
		ID:          "Lean4/Elan",
		Name:        "Elan Executables",
		Description: "提供 Lean4 包管理工具的二进制分发文件",
	},
	{
		ID:          "Lean4/Toolchain",
		Name:        "Lean4 Toolchain files",
		Description: "Lean4 工具链文件",
	},
	{
		ID:          "Lean4/Servior",
		Name:        "Lean4 Lake 包中心",
		Description: "提供 Github 上所有 Star 数大于 2 的项目的检索",
	},
	{
		ID:          "Lean4/GithubRepo/Mathlib4",
		Name:        "Lean4 Github 仓库检查点",
		Description: "",
	},
}
